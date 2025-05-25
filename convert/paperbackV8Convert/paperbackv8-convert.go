package paperbackv8Convert

import (
	"github.com/TheCrether/tachiyomi-paperback-converter/config"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperbackv8"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

type distinctObjectData struct {
	libraryTabs *[]paperbackv8.LibraryMangaLibraryTab
}

type distinctObjectType struct {
	Id string
}

type distinctConverter[K any, V distinctObjectType] func(K) V

func distinct[K any](id string, value *K, distinctMap *map[string]*K) {
	// v := (*value)jj
	(*distinctMap)[id] = value
}

func distinctObjects(data *paperbackv8.PaperbackV8Data) *distinctObjectData {
	libraryTabMap := map[string]*paperbackv8.LibraryMangaLibraryTab{}

	for _, v := range data.LibraryMangas {
		for _, tab := range v.LibraryTabs {
			distinct(tab.Id, &tab, &libraryTabMap)
		}
	}

	libraryTabs := make([]paperbackv8.LibraryMangaLibraryTab, 0, len(libraryTabMap))
	for _, tab := range libraryTabMap {
		libraryTabs = append(libraryTabs, *tab)
	}

	distinctData := distinctObjectData{
		libraryTabs: &libraryTabs,
	}

	return &distinctData
}

func convertLibraryTabsToTachiyomiCategories(tabs *[]paperbackv8.LibraryMangaLibraryTab) []*tachiyomi.BackupCategory {
	categories := make([]*tachiyomi.BackupCategory, 0)
	for _, tab := range *tabs {
		categories = append(categories, &tachiyomi.BackupCategory{
			Name:  tab.Name,
			Flags: 64,
		})
	}
	return categories
}

func ConvertPaperbackV8ToTachiyomi(data *paperbackv8.PaperbackV8Data) (*tachiyomi.Backup, error) {
	backup := config.DefaultTachiyomiBackup()
	distinctData := distinctObjects(data)
	backup.BackupCategories = convertLibraryTabsToTachiyomiCategories(distinctData.libraryTabs)

	return backup, nil
}
