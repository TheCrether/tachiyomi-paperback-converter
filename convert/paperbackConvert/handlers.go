package paperbackConvert

import (
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

var (
	paperbackStatus = map[int32]string{
		0: "Unknown",
		1: "Ongoing",
		2: "Completed",
		3: "Licensed",
		4: "Publishing Finished",
		5: "Cancelled",
		6: "Hiatus",
	}

	// TODO make genre handler for each source
	genreConverter = map[int64]func(*tachiyomi.BackupManga) []paperback.Tag{
		-1: func(tManga *tachiyomi.BackupManga) []paperback.Tag {
			genres := tManga.Genres
			genreList := make([]paperback.TagTag, len(genres))
			for i, genre := range genres {
				genreList[i] = paperback.TagTag{
					Id:    strings.ToLower(genre),
					Value: genre,
				}
			}
			return []paperback.Tag{
				{
					Id:    "0",
					Label: "genres",
					Tags:  genreList,
				},
			}
		},
	}

	// TODO modify handlers to fit extensions
	tachiyomiUrlHandler = map[int64]func(paperback.SourceManga, string){
		-1: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		6247824327199706550: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		7890050626002177109: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		6350607071566689772: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		5904527501696587341: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		3470433521851976761: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		5020395055978987501: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		2499283573021220255: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		2528986671771677900: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		1024627298672457456: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		9: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		734865402529567092: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		5177220001642863679: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		4667040294697888218: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		5190569675461947007: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		2522335540328470744: func(pManga paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
	}
)
