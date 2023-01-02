package config

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
)

var defaultPaperbackSources = `{
	"sourceRepositories": [
    {
      "name": "Extensions Madara (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/madara",
      "type": 0
    },
    {
      "name": "GameFuzzy's Extensions",
      "url": "https://gamefuzzy.github.io/extensions-gamefuzzy/main",
      "type": 0
    },
    {
      "name": "Extensions NSFW (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-sources/nsfw",
      "type": 0
    },
    {
      "name": "Extensions NepNep (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/nepnep",
      "type": 0
    },
    {
      "name": "BuddyComplex Extensions (0.6)",
      "url": "https://thenetsky.github.io/extensions-buddycomplex/0.6",
      "type": 0
    },
    {
      "name": "NmN's Extensions",
      "url": "https://pandeynmn.github.io/nmns-extensions/main",
      "type": 0
    },
    {
      "name": "Extensions Primary (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-sources/primary",
      "type": 0
    },
    {
      "name": "Netsky's Extensions (0.6)",
      "url": "https://thenetsky.github.io/netskys-extensions/0.6",
      "type": 0
    },
    {
      "name": "Extensions MangaBox (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/mangabox",
      "type": 0
    },
    {
      "name": "MangaStream Extensions (0.6)",
      "url": "https://thenetsky.github.io/extensions-mangastream/0.6",
      "type": 0
    }
  ],
  "activeSources": [
    {
      "author": "Daniel Kovalevich",
      "desc": "Extension that pulls manga from MangaSee, includes Advanced Search and Updated manga fetching",
      "website": "https://github.com/DanielKovalevich",
      "id": "Mangasee",
      "tags": [
        { "type": "success", "text": "Notifications" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://mangasee123.com",
      "repo": "https://thenetsky.github.io/extensions-generic/nepnep",
      "version": "2.2.0",
      "icon": "icon.png",
      "name": "MangaSee"
    },
    {
      "author": "NmN",
      "desc": "New Reaperscans source.",
      "website": "http://github.com/pandeynmm",
      "id": "ReaperScans",
      "tags": [
        { "type": "info", "text": "English" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://reaperscans.com",
      "repo": "https://pandeynmn.github.io/nmns-extensions/main",
      "version": "3.0.13",
      "icon": "icon.png",
      "name": "ReaperScans"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from https://1stkissmanga.io",
      "website": "http://github.com/TheNetsky",
      "id": "KissManga",
      "tags": [
        { "type": "success", "text": "Notifications" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://1stkissmanga.io",
      "repo": "https://thenetsky.github.io/extensions-generic/madara",
      "version": "2.2.5",
      "icon": "icon.png",
      "name": "1stKissManga"
    },
    {
      "author": "nar1n",
      "desc": "Extension that pulls manga from MangaDex",
      "website": "https://github.com/nar1n",
      "id": "MangaDex",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://mangadex.org",
      "repo": "https://thenetsky.github.io/extensions-sources/primary",
      "version": "2.1.10",
      "icon": "icon.png",
      "name": "MangaDex"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from https://toonily.com",
      "website": "http://github.com/TheNetsky",
      "id": "Toonily",
      "tags": [
        { "type": "success", "text": "Notifications" },
        { "type": "warning", "text": "18+" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "ADULT",
      "websiteBaseURL": "https://toonily.com",
      "repo": "https://thenetsky.github.io/extensions-generic/madara",
      "version": "2.2.7",
      "icon": "icon.png",
      "name": "Toonily"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from https://teenmanhua.com",
      "website": "http://github.com/TheNetsky",
      "id": "TeenManhua",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://teenmanhua.com",
      "repo": "https://thenetsky.github.io/extensions-generic/madara",
      "version": "2.2.6",
      "icon": "icon.png",
      "name": "TeenManhua"
    },
    {
      "author": "nar1n",
      "desc": "Extension that pulls manga from manganato.com",
      "website": "https://github.com/nar1n",
      "id": "Manganato",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://manganato.com",
      "repo": "https://thenetsky.github.io/extensions-generic/mangabox",
      "version": "3.0.6",
      "icon": "icon.png",
      "name": "Manganato"
    },
    {
      "author": "btylerh7",
      "desc": "Extension that pulls comics from Webtoons.",
      "website": "http://github.com/btylerh7",
      "id": "Webtoons",
      "tags": [{ "type": "info", "text": "Multi-Language" }],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://www.webtoons.com",
      "repo": "https://thenetsky.github.io/extensions-sources/primary",
      "version": "2.1.4",
      "icon": "logo.png",
      "name": "Webtoons"
    },
    {
      "author": "GameFuzzy & NmN",
      "desc": "Extension that pulls western comics from bato.to",
      "website": "http://github.com/gamefuzzy",
      "id": "BatoTo",
      "tags": [
        { "type": "success", "text": "Notifications" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "ADULT",
      "websiteBaseURL": "https://bato.to",
      "repo": "https://gamefuzzy.github.io/extensions-gamefuzzy/main",
      "version": "2.0.3",
      "icon": "icon.png",
      "name": "Bato.To"
    },
    {
      "author": "nar1n",
      "desc": "Extension that pulls manga from mangakakalot.com",
      "website": "https://github.com/nar1n",
      "id": "Mangakakalot",
      "tags": [{ "type": "success", "text": "Partial-Notifications" }],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://mangakakalot.com",
      "repo": "https://thenetsky.github.io/extensions-generic/mangabox",
      "version": "3.0.6",
      "icon": "mangakakalot.com.png",
      "name": "Mangakakalot"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from MangaBuddy",
      "website": "http://github.com/TheNetsky",
      "id": "MangaBuddy",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://mangabuddy.com",
      "repo": "https://thenetsky.github.io/extensions-buddycomplex/0.6",
      "version": "1.1.2",
      "icon": "icon.png",
      "name": "MangaBuddy"
    },
    {
      "author": "NotMarek",
      "desc": "Extension which pulls 18+ content from nHentai.",
      "website": "https://github.com/notmarek",
      "id": "NHentai",
      "tags": [
        { "type": "warning", "text": "18+" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "ADULT",
      "websiteBaseURL": "https://nhentai.net",
      "repo": "https://thenetsky.github.io/extensions-sources/nsfw",
      "version": "3.2.4",
      "icon": "icon.png",
      "name": "nhentai"
    },
    {
      "author": "NmN",
      "desc": "Extension that pulls manga from Flame Scans.",
      "website": "http://github.com/pandeynmm",
      "id": "FlameScans",
      "tags": [
        { "type": "info", "text": "English" },
        { "type": "danger", "text": "Cloudflare" }
      ],
      "contentRating": "EVERYONE",
      "websiteBaseURL": "https://flamescans.org",
      "repo": "https://pandeynmn.github.io/nmns-extensions/main",
      "version": "2.0.5",
      "icon": "icon.ico",
      "name": "FlameScans"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from AsuraScans",
      "website": "http://github.com/TheNetsky",
      "id": "AsuraScans",
      "tags": [
        { "type": "success", "text": "Notifications" },
        { "type": "danger", "text": "CloudFlare" },
        { "type": "danger", "text": "Buggy" }
      ],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://www.asurascans.com",
      "repo": "https://thenetsky.github.io/extensions-mangastream/0.6",
      "version": "2.1.13",
      "icon": "icon.png",
      "name": "AsuraScans"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from ImperfectComic",
      "website": "http://github.com/TheNetsky",
      "id": "ImperfectComic",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://imperfectcomic.org",
      "repo": "https://thenetsky.github.io/extensions-mangastream/0.6",
      "version": "2.1.7",
      "icon": "icon.png",
      "name": "ImperfectComic"
    },
    {
      "author": "Netsky",
      "desc": "Extension that pulls manga from mcreader.net (Manga-Raw.club)",
      "website": "https://github.com/TheNetsky",
      "id": "McReader",
      "tags": [{ "type": "success", "text": "Notifications" }],
      "contentRating": "MATURE",
      "websiteBaseURL": "https://www.mreader.co",
      "repo": "https://thenetsky.github.io/netskys-extensions/0.6",
      "version": "1.0.4",
      "icon": "icon.png",
      "name": "McReader"
    }
  ]
}`

// https://stackoverflow.com/a/33330705/13156660
var SWIFT_REFERENCE_DATE = time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)

func ConvertSwiftReferenceDateToMilliDate(timeSinceReferenceDate float64) int64 {
	return SWIFT_REFERENCE_DATE.Add(time.Duration(int(timeSinceReferenceDate)) * time.Second).UnixMilli()
}

func ConvertTimeToSwiftReferenceDate(time time.Time) float64 {
	return time.Sub(SWIFT_REFERENCE_DATE).Seconds()
}

func ConvertMilliDateToSwiftReferenceDate(milliDate int64) float64 {
	return ConvertTimeToSwiftReferenceDate(time.UnixMilli(milliDate))
}

func PreciseChapterNumberPaperback(chapterNumber float64) float64 {
	str := fmt.Sprintf("%f", chapterNumber)
	if !strings.Contains(str, ".") {
		return chapterNumber
	}
	dotIndex := strings.Index(str, ".")
	firstZero := strings.Index(str, "0")
	if firstZero == -1 {
		return chapterNumber
	}
	firstZero -= dotIndex
	return math.Round(chapterNumber*math.Pow10(firstZero)) / math.Pow10(firstZero)
}

func DefaultPaperbackBackup() *paperback.Backup {
	backup := &paperback.Backup{
		Library:             []paperback.LibraryElement{},
		SourceMangas:        []paperback.SourceManga{},
		ChapterMarkers:      []paperback.ChapterMarker{},
		BackupSchemaVersion: 3,
		Date:                ConvertMilliDateToSwiftReferenceDate(time.Now().UnixMilli()),
		Tabs:                []paperback.Tab{},
		Version:             "v0.7-r45",
	}
	if err := json.Unmarshal([]byte(defaultPaperbackSources), backup); err != nil {
		panic(err)
	}
	return backup
}
