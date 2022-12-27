package config

import "github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"

func DefaultTachiyomiBackup() *tachiyomi.Backup {
	return &tachiyomi.Backup{
		BackupManga: []*tachiyomi.BackupManga{},
		BackupSources: []*tachiyomi.BackupSource{
			{
				Name:     "Bato.to",
				SourceId: 7890050626002177109,
			},
			{
				Name:     "MangaDex",
				SourceId: 2499283573021220255,
			},
			{
				Name:     "Mangakakalot",
				SourceId: 2528986671771677900,
			},
			{
				Name:     "MangaBuddy",
				SourceId: 5020395055978987501,
			},
			{
				Name:     "1st Kiss",
				SourceId: 3470433521851976761,
			},
			{
				Name:     "Manganato",
				SourceId: 1024627298672457456,
			},
			{
				Name:     "Manhwatop",
				SourceId: 5177484976652938680,
			},
			{
				Name:     "Toonily",
				SourceId: 5190569675461947007,
			},
			{
				Name:     "MangaSee",
				SourceId: 9,
			},
			{
				Name:     "Webtoons.com",
				SourceId: 2522335540328470744,
			},
			{
				Name:     "Webtoons.com Translations",
				SourceId: 2909932855774980787,
			},
			{
				Name:     "TeenManhua",
				SourceId: 4667040294697888218,
			},
			{
				Name:     "6350607071566689772",
				SourceId: 6350607071566689772,
			},
			{
				Name:     "Bato.to",
				SourceId: 4531444389842992129,
			},
		},
	}
}
