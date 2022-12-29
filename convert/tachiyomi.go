package convert

import (
	"strings"
	"time"

	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

// use default viewer flag:
// https://github.com/tachiyomiorg/tachiyomi/blob/db3c98fe729ef6b00beba8d605bc002a7b8d1712/app/src/main/java/eu/kanade/tachiyomi/ui/reader/setting/ReadingModeType.kt#L14
const TACHIYOMI_VIEWER_FLAG = 0

var (
	// https://stackoverflow.com/a/33330705/13156660
	SWIFT_REFERENCE_DATE = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

	// https://github.com/tachiyomiorg/tachiyomi/blob/db3c98fe729ef6b00beba8d605bc002a7b8d1712/source-api/src/main/java/eu/kanade/tachiyomi/source/model/SManga.kt#L75
	tachiyomiStatus = map[string]int{
		"Unknown":             0,
		"Ongoing":             1,
		"Completed":           2,
		"Licensed":            3,
		"Publishing Finished": 4,
		"Cancelled":           5,
		"Hiatus":              6,
	}
)

func convertPaperbackTagsToTachiyomiGenres(tags []paperback.Tag) []string {
	genres := make([]string, 2) // assume minimum length of 2 genres per manga
	for _, tag := range tags {
		if strings.ToLower(tag.Label) == "genres" {
			for _, genre := range tag.Tags {
				genres = append(genres, genre.Value)
			}
		}
	}
	return genres
}

func convertDateBookmarkedToTimeAdded(timeBookmarked float64) int64 {
	return SWIFT_REFERENCE_DATE.Add(time.Duration(int(timeBookmarked)) * time.Second).UnixMilli()
}

func ConvertPaperbackToTachiyomi(paperbackBackup *paperback.Backup) (*tachiyomi.Backup, error) {
	backup := config.DefaultTachiyomiBackup()
	for _, libraryElement := range paperbackBackup.Library {
		tachiyomiManga := &tachiyomi.BackupManga{
			Title:        libraryElement.Manga.Titles[0],
			Artist:       libraryElement.Manga.Artist,
			Author:       libraryElement.Manga.Author,
			Description:  libraryElement.Manga.Desc,
			Genres:       convertPaperbackTagsToTachiyomiGenres(libraryElement.Manga.Tags),
			Status:       int32(tachiyomiStatus[libraryElement.Manga.Status]),
			ThumbnailUrl: libraryElement.Manga.Image,
			DateAdded:    convertDateBookmarkedToTimeAdded(libraryElement.DateBookmarked),
			Viewer:       TACHIYOMI_VIEWER_FLAG,
			ViewerFlags:  TACHIYOMI_VIEWER_FLAG,
			History:      make([]*tachiyomi.BackupHistory, 0),
			Tracking:     make([]*tachiyomi.BackupTracking, 0),
			Chapters:     make([]*tachiyomi.BackupChapter, 0),
		}
		err := ConvertPaperbackSourceDataToTachiyomi(paperbackBackup, &libraryElement.Manga, tachiyomiManga)
		if err != nil {
			// TODO add error handling/logging of which manga could not be converted
			continue // skip manga if source data is not found
		}
		backup.BackupManga = append(backup.BackupManga, tachiyomiManga)
	}
	return backup, nil
}
