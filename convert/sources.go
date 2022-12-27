package convert

func reverseMap[K comparable, V comparable](m map[K]V) map[V]K {
	n := make(map[V]K, len(m))
	for k, v := range m {
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

	TachiyomiToPaperback = reverseMap(PaperbackToTachiyomi)
)

func ConvertSourceToTachiyomi(sourceId string) int {
	return PaperbackToTachiyomi[sourceId]
}

func ConvertSourceToPaperback(sourceId int) string {
	return TachiyomiToPaperback[sourceId]
}
