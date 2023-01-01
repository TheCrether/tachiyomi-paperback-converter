package config

import (
	"encoding/json"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

var defaultSources = `[
    { "name": "MangaBuddy", "sourceId": 5020395055978987501 },
    { "name": "TeenManhua", "sourceId": 4667040294697888218 },
    { "name": "1st Kiss", "sourceId": 3470433521851976761 },
    { "name": "MangaDex", "sourceId": 2499283573021220255 },
    { "name": "Toonily", "sourceId": 5190569675461947007 },
    { "name": "Bato.to", "sourceId": 7890050626002177109 },
    { "name": "Flame Scans", "sourceId": 6350607071566689772 },
    { "name": "Asura Scans", "sourceId": 6247824327199706550 },
    { "name": "Imperfect Comics", "sourceId": 5904527501696587341 },
    { "name": "Reaper Scans", "sourceId": 5177220001642863679 },
    { "name": "mcreader.net", "sourceId": 734865402529567092 },
    { "name": "Mangakakalot", "sourceId": 2528986671771677900 },
    { "name": "Manganato", "sourceId": 1024627298672457456 },
    { "name": "MangaSee", "sourceId": 9 },
    { "name": "Webtoons.com", "sourceId": 2522335540328470744 }
  ]`

func DefaultTachiyomiBackup() *tachiyomi.Backup {
	sources := []*tachiyomi.BackupSource{}
	err := json.Unmarshal([]byte(defaultSources), &sources)
	if err != nil {
		panic(err)
	}
	return &tachiyomi.Backup{
		BackupManga:      []*tachiyomi.BackupManga{},
		BackupCategories: []*tachiyomi.BackupCategory{},
		BackupSources:    sources,
	}
}
