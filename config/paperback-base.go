package config

import (
	"encoding/json"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
)

var defaultBackup = `{
  "library": [],
  "sourceMangas": [],
  "chapterMarkers": [],
  "backupSchemaVersion": 3,
  "date": 693525190.64852703,
  "tabs": [],
  "version": "v0.7-r45",
  "sourceRepositories": [
    {
      "name": "Extensions NepNep (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/nepnep",
      "type": 0
    },
    {
      "name": "GameFuzzy's Extensions",
      "url": "https://gamefuzzy.github.io/extensions-gamefuzzy/main",
      "type": 0
    },
    {
      "name": "BuddyComplex Extensions (0.6)",
      "url": "https://thenetsky.github.io/extensions-buddycomplex/0.6",
      "type": 0
    },
    {
      "name": "Extensions Primary (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-sources/primary",
      "type": 0
    },
    {
      "name": "Extensions MangaBox (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/mangabox",
      "type": 0
    },
    {
      "name": "Extensions Madara (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-generic/madara",
      "type": 0
    },
    {
      "name": "Extensions NSFW (Netsky Fork)",
      "url": "https://thenetsky.github.io/extensions-sources/nsfw",
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
    }
  ]
}`

func DefaultPaperbackBackup() *paperback.Backup {
	backup := &paperback.Backup{}
	if err := json.Unmarshal([]byte(defaultBackup), backup); err != nil {
		panic(err)
	}
	return backup
}
