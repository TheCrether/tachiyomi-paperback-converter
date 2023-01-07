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

| Source               | Manga | Chapters Read | Notes              |
| -------------------- | ----- | ------------- | ------------------ |
| Asura Scans          | ✔️     | ✔️             |                    |
| BatoTo               | ✔️     | ✔️             |                    |
| Flame Scans          | ✔️     | ✔️             |                    |
| Imperfect Comics     | ✔️     | ✔️             |                    |
| KissManga (1st Kiss) | ✔️     | ✔️             |                    |
| MangaBuddy           | ✔️     | ✔️             |                    |
| MangaDex             | ✔️     | ✔️             |                    |
| Mangakakalot         | ✔️     | ✔️             | might be unstable¹ |
| Manganato            | ✔️     | ✔️             | might be unstable¹ |
| Mangasee             | ✔️     | ✔️             |                    |
| MCReader.net         | ✔️     | ✔️             |                    |
| Reaper Scans         | ✔️     | ✔️             |                    |
| TeenManhua           | ✔️     | ✔️             |                    |
| Toonily              | ✔️     | ✔️             |                    |
| Webtoons             | ✔️     | ❌             |                    |

¹ - MangaDex and Mangakakalot are known to be unstable, since the Paperback implementation is inconsistent. If there is an issue when converting with these two sources, please open an issue. (especially for the chapters read conversion)

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
