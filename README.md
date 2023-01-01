# tachiyomi-paperback-converter

This is a project which can convert backups from Paperback/Tachiyomi to make it easier to switch platforms.

## Supported Manga Readers

The two main manga readers that are supported are:

- Paperback
- Tachiyomi

Support for other manga readers may be added in the future.

## Paperback -> Tachiyomi

### Collections

Groups of Manga are called Collections in Paperback. These are converted to Categories in Tachiyomi.

### Sources

| Source               | Manga | Chapters Read | Notes |
| -------------------- | ----- | ------------- | ----- |
| BatoTo               | ✔️     | ❌             |       |
| MangaDex             | ✔️     | ❌             |       |
| Mangakakalot         | ✔️     | ❌             |       |
| MangaBuddy           | ✔️     | ❌             |       |
| KissManga (1st Kiss) | ✔️     | ❌             |       |
| Manganato            | ✔️     | ❌             |       |
| Toonily              | ✔️     | ❌             |       |
| Mangasee             | ✔️     | ❌             |       |
| Webtoons             | ✔️     | ❌             |       |
| TeenManhua           | ✔️     | ❌             |       |
| Flame Scans          | ❌     | ❌             |       |
| Asura Scans          | ❌     | ❌             |       |
| Reaper Scans         | ❌     | ❌             |       |
| Imperfect Comics     | ❌     | ❌             |       |
| MCReader.net         | ❌     | ❌             |       |

## Tachiyomi -> Paperback

### Categories

Categories in Tachiyomi are converted to Collections in Paperback.

### Sources

WIP - no support for any sources yet

## Development

TODO

## Technologies used

- protobuf / protoc
- json
