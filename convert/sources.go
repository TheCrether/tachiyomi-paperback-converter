package convert

import (
	"errors"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
)

func reverseSourceMap[K comparable, V comparable](m map[K]V) map[V]K {
	n := make(map[V]K, len(m))
	for k, v := range m {
		// check if value is default value of type V (like 0 for int, which means it is invalid)
		if v == *new(V) {
			continue
		}
		n[v] = k
	}
	return n
}

var (
	PaperbackToTachiyomi = map[string]int{
		"BatoTo":                    7890050626002177109,
		"MangaDex":                  2499283573021220255,
		"Mangakakalot":              2528986671771677900,
		"MangaBuddy":                5020395055978987501,
		"KissManga":                 3470433521851976761,
		"Manganato":                 1024627298672457456,
		"Manhwatop":                 0, // TODO: find extension for Manhwatop
		"Toonily":                   5190569675461947007,
		"Mangasee":                  9,
		"Webtoons":                  2522335540328470744,
		"Webtoons.com Translations": 0, // TODO check if normal webtoons paperback extension is enough
		"TeenManhua":                4667040294697888218,
		// "6350607071566689772": 6350607071566689772,
		// "Bato.to": 4531444389842992129, // TODO: check second bato.to source in tachiyomi
	}

	TachiyomiToPaperback = reverseSourceMap(PaperbackToTachiyomi)

	TachiyomiSources = map[int64]func(*tachiyomi.BackupManga, string){
		7890050626002177109: func(tManga *tachiyomi.BackupManga, mangaId string) { // batoto
			tManga.Url = "/series/" + mangaId
		},
		2499283573021220255: func(tManga *tachiyomi.BackupManga, mangaId string) { // mangadex
			tManga.Url = "/manga/" + mangaId
		},
		2528986671771677900: func(tManga *tachiyomi.BackupManga, mangaId string) { // mangakakalot
			// change source id depending on mangaId -> manganato -> manga-<id>, mangakakalot -> manga are kinda interchangeable
			// TODO: handler which requests mangakakalot to get the correct mangaId (if it has to be manganato or else)
			tManga.Url = mangaId
		},
		5020395055978987501: func(tManga *tachiyomi.BackupManga, mangaId string) { // mangabuddy
			tManga.Url = "/" + mangaId
		},
		3470433521851976761: func(tManga *tachiyomi.BackupManga, mangaId string) { // kissmanga
			tManga.Url = "/manga/" + mangaId + "/"
		},
		1024627298672457456: func(tManga *tachiyomi.BackupManga, mangaId string) { // manganato
			// TODO: handler which requests mangakakalot to get the correct mangaId (if it has to be manganato or else)
			tManga.Url = "/manga/" + mangaId
		},
		5190569675461947007: func(tManga *tachiyomi.BackupManga, mangaId string) { // toonily
			tManga.Url = "/webtoon/" + mangaId + "/"
		},
		9: func(tManga *tachiyomi.BackupManga, mangaId string) { // mangasee
			tManga.Url = "/manga/" + strings.ToLower(mangaId)
		},
		2522335540328470744: func(tManga *tachiyomi.BackupManga, mangaId string) { // webtoons
			split := strings.Split(mangaId, "/")
			tManga.Genres = []string{split[0]}
			tManga.Url = "/episodeList?titleNo=" + strings.Split(mangaId, "=")[1]
		},
		4667040294697888218: func(tManga *tachiyomi.BackupManga, mangaId string) { // teenmanhua
			tManga.Url = "/manga/" + mangaId
		},
	}
)

func getSourceForPaperbackmanga(pBackup *paperback.Backup, manga *paperback.Manga) (*paperback.SourceManga, error) {
	for _, sourceManga := range pBackup.SourceMangas {
		if sourceManga.Manga.Id == manga.Id {
			return &sourceManga, nil
		}
	}
	return nil, errors.New("source not found")
}

func ConvertPaperbackSourceDataToTachiyomi(paperbackBackup *paperback.Backup, paperbackManga *paperback.Manga, tachiyomiManga *tachiyomi.BackupManga) error {
	paperbackSource, err := getSourceForPaperbackmanga(paperbackBackup, paperbackManga)
	if err != nil {
		return err
	}
	tachiyomiSourceId, err := ConvertSourceMangaToTachiyomi(paperbackSource)
	if err != nil {
		return err
	}
	tachiyomiManga.Source = tachiyomiSourceId
	TachiyomiSources[tachiyomiSourceId](tachiyomiManga, paperbackSource.MangaId)
	return nil
}

func ConvertSourceMangaToTachiyomi(paperbackSource *paperback.SourceManga) (int64, error) {
	source, ok := PaperbackToTachiyomi[paperbackSource.SourceId]
	if !ok {
		return 0, errors.New("source not found")
	}
	return int64(source), nil
}

func ConvertTachiyomiSourceIdToPaperbackId(sourceId int) (string, error) {
	source, ok := TachiyomiToPaperback[sourceId]
	if !ok {
		return "", errors.New("source not found")
	}
	return source, nil
}
