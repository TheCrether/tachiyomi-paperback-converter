package tachiyomiConvert

import (
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

func mangakakalotUrlHandler(tManga *tachiyomi.BackupManga, mangaId string) {
	if strings.HasPrefix(mangaId, "http") {
		// this is the normal case since the paperback extension is sensible enough to be consistent
		tManga.Url = mangaId
		return
	}
	tManga.Url = "/manga/" + mangaId
}

func manganatoReadHandler(marker *paperback.ChapterMarker, chapterId string) string {
	return chapterId
}

var (
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

	// the handler functions are used to convert the mangaId to the url of the manga depending on the source.
	// if a source is not found in the map, the source is not supported for conversion
	paperbackUrlHandler = map[string]func(*tachiyomi.BackupManga, string){
		"AsuraScans": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + mangaId
		},
		"BatoTo": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/series/" + mangaId
		},
		"FlameScans": func(tManga *tachiyomi.BackupManga, mangaId string) {
			split := strings.SplitN(mangaId, "-", 2)
			tManga.Url = "/series/" + split[1] + "/"
		},
		"KissManga": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + mangaId + "/"
		},
		"MangaBuddy": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/" + mangaId
		},
		"MangaDex": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + mangaId
		},
		"Mangakakalot": mangakakalotUrlHandler,
		"Manganato":    mangakakalotUrlHandler,
		"Mangasee": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + strings.ToLower(mangaId)
		},
		"McReader": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + mangaId + "/"
		},
		"ReaperScans": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/comics/" + mangaId
		},
		"TeenManhua": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/manga/" + mangaId
		},
		"Toonily": func(tManga *tachiyomi.BackupManga, mangaId string) {
			tManga.Url = "/webtoon/" + mangaId + "/"
		},
		"Webtoons": func(tManga *tachiyomi.BackupManga, mangaId string) {
			split := strings.Split(mangaId, "/")
			tManga.Genres = []string{split[0]}
			tManga.Url = "/episodeList?titleNo=" + strings.Split(mangaId, "=")[1]
		},
	}

	tachiyomiReadHandler = map[string]func(*paperback.ChapterMarker, string) string{
		"AsuraScans": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/" + chapterId + "/"
		},
		"BatoTo": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/chapter/" + chapterId
		},
		"FlameScans": func(marker *paperback.ChapterMarker, chapterId string) string {
			chapterId = strings.ReplaceAll(chapterId, "https://flamescans.org/"+marker.Chapter.MangaId, "")
			return "/" + chapterId + "/"
		},
		"KissManga": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "https://1stkissmanga.io/manga/" + chapterId + "/?style=list"
		},
		"MangaBuddy": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/" + marker.Chapter.MangaId + "/" + chapterId
		},
		"MangaDex": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/chapter/" + chapterId
		},
		"Mangakakalot": manganatoReadHandler,
		"Manganato":    manganatoReadHandler,
		"Mangasee": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/read-online/" + chapterId
		},
		"McReader": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/reader/en/" + chapterId + "/"
		},
		"ReaperScans": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "/comics/" + marker.Chapter.MangaId + "/chapters/" + chapterId
		},
		"TeenManhua": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "https://teenmanhua.com/manga/" + chapterId + "/?style=list"
		},
		"Toonily": func(marker *paperback.ChapterMarker, chapterId string) string {
			return "https://toonily.com/webtoon/" + chapterId + "/?style=list"
		},
		// TODO find good way to convert paperback structure to weird tachiyomi string
		// "Webtoons": func(marker *paperback.ChapterMarker, chapterId string) string {
		// 	split := strings.Split(chapterId, "/list")
		// 	return fmt.Sprintf("/en/%s/episode-%d", split[0], int(marker.Chapter.ChapNum))
		// },
	}
)
