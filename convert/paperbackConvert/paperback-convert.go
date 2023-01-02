package paperbackConvert

import (
	"errors"
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

func convertTachiyomiGenres(genres []string) []paperback.Tag {
	// TODO make handler for each source
	genreList := make([]paperback.TagTag, len(genres))
	for i, genre := range genres {
		genreList[i] = paperback.TagTag{
			Id:    strings.ToLower(genre),
			Value: genre,
		}
	}
	paperbackGenres := []paperback.Tag{
		{
			Id:    "0",
			Label: "genres",
			Tags:  genreList,
		},
	}
	return paperbackGenres
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

// TODO ConvertTachiyomiToPaperback
func ConvertTachiyomiToPaperback(tBackup *tachiyomi.Backup) (*paperback.Backup, error) {
	backup := config.DefaultPaperbackBackup()

	mangaUUIDs := make([]string, 0, len(tBackup.BackupManga))
	sourceMangaUUIDs := make([]string, 0, len(tBackup.BackupManga))
	mangaUUID := ""
	sourceMangaUUID := ""

	backup.Tabs = getTabs(tBackup)

	for _, manga := range tBackup.BackupManga {
		source, ok := convert.ConvertTachiyomiSourceIdToPaperbackId(manga.Source)
		if ok != nil {
			continue
		}
		mangaUUIDs, mangaUUID = createUniqueUUID(mangaUUIDs)
		pManga := &paperback.Manga{
			Id:     mangaUUID,
			Rating: float64(0),
			Author: manga.Author,
			Artist: manga.Artist,
			Titles: []string{manga.Title},
			Tags:   convertTachiyomiGenres(manga.Genres),
			Desc:   manga.Description,
			Image:  manga.ThumbnailUrl,
			Hentai: false, // tachiyomi doesn't seem to have a flag for this
			AdditionalInfo: paperback.AdditionalInfo{
				LangFlag:  "en", // just set en as default since tachiyomi doesn't have a flag for this
				Users:     "0",
				Follows:   "0",
				AvgRating: "0.0",
				Views:     "0",
				LangName:  "English",
			},
			Status: paperbackStatus[manga.Status],
		}
		libraryElement := &paperback.LibraryElement{
			Manga:          *pManga,
			LastRead:       config.ConvertMilliDateToSwiftReferenceDate(getLastRead(manga)),
			LastUpdated:    0, // TODO maybe get from highest dateFetch from chapter?
			DateBookmarked: config.ConvertMilliDateToSwiftReferenceDate(manga.DateAdded),
			LibraryTabs:    []paperback.LibraryTab{}, // TODO look at tachiyomi way of saving categories
			Updates:        0,
		}
		sourceMangaUUIDs, sourceMangaUUID = createUniqueUUID(sourceMangaUUIDs)
		sourceManga := &paperback.SourceManga{
			Manga:        *pManga,
			OriginalInfo: *pManga,
			SourceId:     source,
			Id:           sourceMangaUUID,
			MangaId:      "TODO", // TODO add handlers to convert URL
		}

		backup.Library = append(backup.Library, *libraryElement)
		backup.SourceMangas = append(backup.SourceMangas, *sourceManga)
	}
	return backup, errors.New("not implemented")
}
