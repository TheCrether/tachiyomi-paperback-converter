package paperbackConvert

import (
	"sort"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/convert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"github.com/google/uuid"
)

func contains[K comparable](list []K, toCheck K) bool {
	for _, item := range list {
		if item == toCheck {
			return true
		}
	}
	return false
}

func createUUID() string {
	return strings.ToUpper(uuid.New().String())
}

func createUniqueUUID(existingUUIDs []string) ([]string, string) {
	for {
		uuid := createUUID()
		if !contains(existingUUIDs, uuid) {
			return append(existingUUIDs, uuid), uuid
		}
	}
}

func convertTachiyomiGenres(tManga *tachiyomi.BackupManga) []paperback.Tag {
	var genreIdHandler func(string) string
	var ok bool
	if genreIdHandler, ok = genreIdConverter[tManga.Source]; !ok {
		genreIdHandler = genreIdConverter[-1]
	}
	genres := tManga.Genres
	if len(genres) == 0 {
		return []paperback.Tag{}
	}
	genreList := make([]paperback.TagTag, len(genres))
	for i, genre := range genres {
		genreList[i] = paperback.TagTag{
			Id:    genreIdHandler(genre),
			Value: genre,
		}
	}
	genreTag := paperback.Tag{
		Id:    "0",
		Label: "genres",
		Tags:  genreList,
	}
	var genreTagHandler func(*paperback.Tag)
	if genreTagHandler, ok = genreTagConverter[tManga.Source]; ok {
		genreTagHandler(&genreTag)
	}
	tags := []paperback.Tag{
		genreTag,
	}
	var extras []paperback.Tag
	if extras, ok = tagExtras[tManga.Source]; ok {
		tags = append(tags, extras...)
	}
	return tags
}

func getLastRead(manga *tachiyomi.BackupManga) int64 {
	lastRead := int64(0)
	for _, historyItem := range manga.History {
		if historyItem.LastRead > lastRead {
			lastRead = historyItem.LastRead
		}
	}
	return lastRead
}

func getTabs(tBackup *tachiyomi.Backup) []paperback.Tab {
	tabs := make([]paperback.Tab, len(tBackup.BackupCategories))
	tabUUIDs := make([]string, 0, len(tBackup.BackupCategories))
	tabUUID := ""
	for i, category := range tBackup.BackupCategories {
		tabUUIDs, tabUUID = createUniqueUUID(tabUUIDs)
		tabs[i] = paperback.Tab{
			Id:        tabUUID,
			Name:      category.Name,
			SortOrder: int(category.Order), // TODO check behaviour of category.Order with multiple categories
		}
	}
	return tabs
}

func getTabsForManga(pBackup *paperback.Backup, manga *tachiyomi.BackupManga) []paperback.LibraryTab {
	tabs := make([]paperback.LibraryTab, len(manga.Categories))
	for i, category := range manga.Categories {
		tabs[i] = paperback.LibraryTab(pBackup.Tabs[category])
	}
	return tabs
}

func getLastDateFetch(tManga *tachiyomi.BackupManga) int64 {
	highestDateFetch := int64(0)
	for _, chapter := range tManga.Chapters {
		if chapter.DateFetch > highestDateFetch {
			highestDateFetch = chapter.DateFetch
		}
	}
	return highestDateFetch
}

func convertTachiyomiChapters(tManga *tachiyomi.BackupManga, sourceMangaId string) []paperback.ChapterMarker {
	filteredChapters := make([]*tachiyomi.BackupChapter, 0, len(tManga.Chapters))
	for _, chapter := range tManga.Chapters {
		if chapter.Read || chapter.LastPageRead > 0 {
			filteredChapters = append(filteredChapters, chapter)
		}
	}

	sort.Slice(filteredChapters, func(i, j int) bool {
		return filteredChapters[i].ChapterNumber < filteredChapters[j].ChapterNumber
	})
	chapters := make([]paperback.ChapterMarker, 0, len(tManga.Chapters))

	for i, chapter := range tManga.Chapters {
		sourceId, err := convert.ConvertTachiyomiSourceIdToPaperbackId(tManga.Source)
		if err != nil {
			continue
		}

		chapters[i] = paperback.ChapterMarker{
			Completed:  chapter.Read,
			Time:       config.ConvertMilliDateToSwiftReferenceDate(chapter.DateFetch),
			TotalPages: 0, // HINT no total pages field in tachiyomi backup
			LastPage:   int(chapter.GetLastPageRead()),
			Hidden:     false, // HINT no hidden field in tachiyomi backup
			Chapter: paperback.Chapter{
				Id:           tachiyomiIdConverter[tManga.Source](tManga.Url),
				ChapNum:      float64(chapter.ChapterNumber),
				MangaId:      sourceMangaId,
				Volume:       0, // HINT no volume field in tachiyomi backup
				LangCode:     convert.TachiyomiToLangCode[tManga.Source],
				Time:         config.ConvertMilliDateToSwiftReferenceDate(chapter.DateUpload),
				SortingIndex: float64(i),
				Group:        chapter.Scanlator,
				Name:         chapter.Name,
				SourceId:     sourceId,
			},
		}
	}
	return chapters
}

// TODO ConvertTachiyomiToPaperback
func ConvertTachiyomiToPaperback(tBackup *tachiyomi.Backup) (*paperback.Backup, error) {
	backup := config.DefaultPaperbackBackup()

	mangaUUIDs := make([]string, 0, len(tBackup.BackupManga))
	sourceMangaUUIDs := make([]string, 0, len(tBackup.BackupManga))
	mangaUUID := ""
	sourceMangaUUID := ""

	backup.Tabs = getTabs(tBackup)

	for _, manga := range tBackup.BackupManga {
		source, err := convert.ConvertTachiyomiSourceIdToPaperbackId(manga.Source)
		if err != nil {
			continue
		}

		mangaIdHandler, ok := tachiyomiUrlHandler[manga.Source]
		if !ok {
			// TODO error handling/logging which mangas could not be converted
			continue
		}

		mangaUUIDs, mangaUUID = createUniqueUUID(mangaUUIDs)
		pManga := &paperback.Manga{
			Id:     mangaUUID,
			Rating: float64(0),
			Covers: []string{},
			Author: manga.Author,
			Artist: manga.Artist,
			Titles: []string{manga.Title},
			Tags:   convertTachiyomiGenres(manga),
			Desc:   manga.Description,
			Image:  manga.ThumbnailUrl,
			Hentai: false, // tachiyomi doesn't seem to have a flag for this
			AdditionalInfo: paperback.AdditionalInfo{
				LangFlag:  convert.TachiyomiToLangCode[manga.Source],
				Users:     "0",
				Follows:   "0",
				AvgRating: "0.0",
				Views:     "0",
				LangName:  convert.LangCodeToFullLang[convert.TachiyomiToLangCode[manga.Source]],
			},
			Status: paperbackStatus[manga.Status],
		}
		libraryElement := &paperback.LibraryElement{
			Manga:          *pManga,
			LastRead:       config.ConvertMilliDateToSwiftReferenceDate(getLastRead(manga)),
			LastUpdated:    config.ConvertMilliDateToSwiftReferenceDate(getLastDateFetch(manga)),
			DateBookmarked: config.ConvertMilliDateToSwiftReferenceDate(manga.DateAdded),
			LibraryTabs:    getTabsForManga(backup, manga),
			Updates:        0,
		}
		sourceMangaUUIDs, sourceMangaUUID = createUniqueUUID(sourceMangaUUIDs)
		sourceManga := &paperback.SourceManga{
			Manga:        *pManga,
			OriginalInfo: *pManga,
			SourceId:     source,
			Id:           sourceMangaUUID,
		}
		mangaIdHandler(sourceManga, manga.Url)

		backup.Library = append(backup.Library, *libraryElement)
		backup.SourceMangas = append(backup.SourceMangas, *sourceManga)
		backup.ChapterMarkers = append(backup.ChapterMarkers, convertTachiyomiChapters(manga, sourceMangaUUID)...)
	}
	return backup, nil
}
