package convert

import (
	"errors"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
)

func reverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	n := make(map[V]K, len(m))
	for k, v := range m {
		// check if value is default value of type V (like 0 for int, which we will view is invalid)
		if v == *new(V) {
			continue
		}
		n[v] = k
	}
	return n
}

var (
	PaperbackToTachiyomi = map[string]int64{
		"AsuraScans":   6247824327199706550,
		"BatoTo":       7890050626002177109,
		"FlameScans":   6350607071566689772,
		"KissManga":    3470433521851976761,
		"MangaBuddy":   5020395055978987501,
		"MangaDex":     2499283573021220255,
		"Mangakakalot": 2528986671771677900,
		"Manganato":    1024627298672457456,
		"Mangasee":     9,
		"McReader":     734865402529567092,
		"ReaperScans":  5177220001642863679,
		"TeenManhua":   4667040294697888218,
		"Toonily":      5190569675461947007,
		"Webtoons":     2522335540328470744,
	}

	TachiyomiToPaperback = reverseMap(PaperbackToTachiyomi)

	// information from: https://github.com/tachiyomiorg/tachiyomi-extensions/blob/repo/index.json
	TachiyomiToLangCode = map[int64]string{
		6247824327199706550: "gb",
		7890050626002177109: "gb",
		6350607071566689772: "gb",
		3470433521851976761: "gb",
		5020395055978987501: "gb",
		2499283573021220255: "gb",
		2528986671771677900: "gb",
		1024627298672457456: "gb",
		9:                   "gb",
		734865402529567092:  "gb",
		5177220001642863679: "gb",
		4667040294697888218: "gb",
		5190569675461947007: "gb",
		2522335540328470744: "gb",
	}

	LangCodeToFullLang = map[string]string{
		"en": "English",
	}
)

func ConvertSourceMangaToTachiyomi(paperbackSource *paperback.SourceManga) (int64, error) {
	source, ok := PaperbackToTachiyomi[paperbackSource.SourceId]
	if !ok {
		return 0, errors.New("source not found")
	}
	return int64(source), nil
}

func ConvertTachiyomiSourceIdToPaperbackId(sourceId int64) (string, error) {
	source, ok := TachiyomiToPaperback[sourceId]
	if !ok {
		return "", errors.New("source not found")
	}
	return source, nil
}
