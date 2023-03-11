package paperbackConvert

import (
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
)

func mangakakalotGenreHandler(genre string) string {
	// TODO add more genre conversions for mangakakalot/-nato
	convertMap := map[string]string{
		"Comedy":        "6",
		"Romance":       "27",
		"School Life":   "28",
		"Seinen":        "30",
		"Slice of Life": "35",
	}
	if converted, ok := convertMap[genre]; ok {
		return converted
	}
	return genre
}

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

	// TODO modify handlers to fit extensions
	tachiyomiUrlHandler = map[int64]func(*paperback.SourceManga, string){
		6247824327199706550: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/manga/", "", -1)
		},
		7890050626002177109: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/series/", "", -1)
		},
		6350607071566689772: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId[:len(mangaId)-1], "/series/", "", -1)
		},
		3470433521851976761: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId[:len(mangaId)-1], "/manga/", "", -1)
		},
		5020395055978987501: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId[1:]
		},
		2499283573021220255: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/manga/", "", -1)
		},
		2528986671771677900: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		1024627298672457456: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = mangaId
		},
		9: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/manga/", "", -1)
		},
		734865402529567092: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/manga/", "", -1)
		},
		5177220001642863679: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/comics/", "", -1)
		},
		4667040294697888218: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId, "/manga/", "", -1)
		},
		5190569675461947007: func(pManga *paperback.SourceManga, mangaId string) {
			pManga.MangaId = strings.Replace(mangaId[:len(mangaId)-1], "/webtoon/", "", -1)
		},
		2522335540328470744: func(pManga *paperback.SourceManga, mangaId string) {
			genre := pManga.Manga.Tags[0].Tags[0].Value
			titleNo := strings.Replace(mangaId, "/episodeList", "", -1)
			pManga.MangaId = genre + "/" + mangaId + "/list" + titleNo
		},
	}

	// -1 converts the string to lowercase and the sources with a different approach have their own handler
	genreIdConverter = map[int64]func(string) string{
		-1: func(genre string) string {
			return strings.ToLower(genre)
		},
		7890050626002177109: func(genre string) string {
			return genre
		},
		6350607071566689772: func(genre string) string {
			return "https:/flamescans.org/ygd/genres/" + strings.ToLower(genre) + "/"
		},
		2499283573021220255: func(genre string) string {
			// TODO add more genre conversions for mangadex
			convertMap := map[string]string{
				"Romance":       "423e2eae-a7a2-4a8b-ac03-a8351462d71d",
				"Comedy":        "4d32cc48-9f00-4cca-9b5a-a839f0764984",
				"School Life":   "caaa44eb-cd40-4177-b930-79d3ef2afe87",
				"Slice of Life": "e5301a23-ebd9-49dd-a0cb-2add944c7fe9",
			}
			if converted, ok := convertMap[genre]; ok {
				return converted
			}
			return strings.ToLower(genre)
		},
		2528986671771677900: mangakakalotGenreHandler,
		1024627298672457456: mangakakalotGenreHandler,
		734865402529567092: func(genre string) string {
			return genre
		},
		// Reaper Scans apparently doesn't have genres
		5190569675461947007: func(genre string) string {
			return strings.Replace(strings.ToLower(genre), " ", "-", -1)
		},
		2522335540328470744: func(genre string) string {
			return strings.ToUpper(genre[:1]) + strings.ToLower(genre[1:])
		},
	}
)
