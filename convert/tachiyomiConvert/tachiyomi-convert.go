package tachiyomiConvert

import (
	"errors"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/convert/commonConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

// use default viewer flag:
// https://github.com/tachiyomiorg/tachiyomi/blob/db3c98fe729ef6b00beba8d605bc002a7b8d1712/app/src/main/java/eu/kanade/tachiyomi/ui/reader/setting/ReadingModeType.kt#L14
var TACHIYOMI_VIEWER_FLAG int32 = 0

func getChapterMarkers(pBackup *paperback.Backup, pSourceManga *paperback.SourceManga) []paperback.ChapterMarker {
	var chapterMarkers []paperback.ChapterMarker
	for _, chapterMarker := range pBackup.ChapterMarkers {
		if chapterMarker.Chapter.MangaId == pSourceManga.MangaId {
			chapterMarkers = append(chapterMarkers, chapterMarker)
		}
	}
	return chapterMarkers
}

func getSourceMangaForPaperbackmanga(pBackup *paperback.Backup, manga *paperback.Manga) (*paperback.SourceManga, error) {
	for _, sourceManga := range pBackup.SourceMangas {
		if sourceManga.Manga.Id == manga.Id {
			return &sourceManga, nil
		}
	}
	return nil, errors.New("source not found")
}

func convertPaperbackSourceDataToTachiyomi(paperbackBackup *paperback.Backup, paperbackManga *paperback.Manga, tachiyomiManga *tachiyomi.BackupManga) error {
	pSourceManga, err := getSourceMangaForPaperbackmanga(paperbackBackup, paperbackManga)
	if err != nil {
		return err
	}
	tachiyomiSourceId, err := commonConvert.ConvertSourceMangaToTachiyomi(pSourceManga)
	if err != nil {
		return err
	}
	tachiyomiManga.Source = tachiyomiSourceId
	if handler, ok := paperbackUrlHandler[pSourceManga.SourceId]; ok {
		handler(tachiyomiManga, pSourceManga.MangaId)
	} else {
		return errors.New("could not convert source")
	}
	if handler, ok := paperbackReadHandler[pSourceManga.SourceId]; ok {
		chapterMarkers := getChapterMarkers(paperbackBackup, pSourceManga)
		for _, chapterMarker := range chapterMarkers {
			tChapter := &tachiyomi.BackupChapter{
				Url:           handler(&chapterMarker, chapterMarker.Chapter.Id),
				Name:          chapterMarker.Chapter.Name,
				Read:          chapterMarker.Completed,
				LastPageRead:  int64(chapterMarker.LastPage),
				DateFetch:     config.ConvertSwiftReferenceDateToMilliDate(paperbackBackup.Date),
				DateUpload:    config.ConvertSwiftReferenceDateToMilliDate(chapterMarker.Chapter.Time),
				ChapterNumber: float32(config.PreciseChapterNumberPaperback(chapterMarker.Chapter.ChapNum)),
				Scanlator:     chapterMarker.Chapter.Group,
			}
			tachiyomiManga.Chapters = append(tachiyomiManga.Chapters, tChapter)
		}
	}
	return nil
}

func convertPaperbackTagsToTachiyomiGenres(tags []paperback.Tag) []string {
	genres := make([]string, 2) // assume minimum length of 2 genres per manga
	for _, tag := range tags {
		if strings.ToLower(tag.Label) == "genres" {
			for _, genre := range tag.Tags {
				genres = append(genres, strings.ToUpper(genre.Value[:1])+strings.ToLower(genre.Value[1:]))
			}
		}
	}
	return genres
}

func convertPaperbackTabsToTachiyomiCategories(tabs []paperback.Tab) []*tachiyomi.BackupCategory {
	categories := make([]*tachiyomi.BackupCategory, 0)
	for _, tab := range tabs {
		categories = append(categories, &tachiyomi.BackupCategory{
			Name:  tab.Name,
			Flags: 64,
		})
	}
	return categories
}

func ConvertPaperbackToTachiyomi(paperbackBackup *paperback.Backup) (*tachiyomi.Backup, error) {
	backup := config.DefaultTachiyomiBackup()
	backup.BackupCategories = convertPaperbackTabsToTachiyomiCategories(paperbackBackup.Tabs)
	for _, libraryElement := range paperbackBackup.Library {
		tachiyomiManga := &tachiyomi.BackupManga{
			Title:        libraryElement.Manga.Titles[0],
			Artist:       libraryElement.Manga.Artist,
			Author:       libraryElement.Manga.Author,
			Description:  libraryElement.Manga.Desc,
			Genres:       convertPaperbackTagsToTachiyomiGenres(libraryElement.Manga.Tags),
			Status:       int32(tachiyomiStatus[libraryElement.Manga.Status]),
			ThumbnailUrl: libraryElement.Manga.Image,
			DateAdded:    config.ConvertSwiftReferenceDateToMilliDate(libraryElement.DateBookmarked),
			Viewer:       TACHIYOMI_VIEWER_FLAG,
			ViewerFlags:  TACHIYOMI_VIEWER_FLAG,
			History:      make([]*tachiyomi.BackupHistory, 0),
			Tracking:     make([]*tachiyomi.BackupTracking, 0),
			Chapters:     make([]*tachiyomi.BackupChapter, 0),
		}
		err := convertPaperbackSourceDataToTachiyomi(paperbackBackup, &libraryElement.Manga, tachiyomiManga)
		if err != nil {
			// TODO add error handling/logging of which manga could not be converted
			continue // skip manga if source data is not found
		}
		backup.BackupManga = append(backup.BackupManga, tachiyomiManga)
	}
	return backup, nil
}
