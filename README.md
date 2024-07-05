# tachiyomi-paperback-converter

## This project is not being actively maintained at the moment and does not support the latest Paperback versions & Tachiyomi alternatives

This is a project which can convert backups from Paperback/Tachiyomi to make it easier to switch platforms.

## Supported Manga Readers

The two main manga readers that are supported are:

- Paperback v0.7 (0.8 currently WIP)
- Tachiyomi

Support for other manga readers may be added in the future.

## Paperback -> Tachiyomi

### Collections

Groups of Manga are called Collections in Paperback. These are converted to Categories in Tachiyomi.

### Sources 0.7

| Source               | Manga | Chapters Read | Notes     |
| -------------------- | ----- | ------------- | --------- |
| Asura Scans          | ✔️     | ✔️             |           |
| BatoTo               | ✔️     | ✔️             |           |
| KissManga (1st Kiss) | ✔️     | ✔️             |           |
| MangaBuddy           | ✔️     | ✔️             |           |
| MangaDex             | ✔️     | ✔️             |           |
| Mangakakalot         | ✔️¹    | ✔️¹            | unstable¹ |
| Manganato            | ✔️¹    | ✔️¹            | unstable¹ |
| Mangasee             | ✔️     | ✔️             |           |
| Mgeko (McReader)     | ✔️     | ✔️             |           |
| Reaper Scans         | ✔️     | ✔️             |           |
| Toonily              | ✔️     | ✔️             |           |
| Webtoons             | ✔️     | ❌             |           |

¹ - Mangakakalot and Manganato are known to be unstable, since the Paperback implementation is inconsistent. If there is an issue when converting with these two sources, please open an issue. (especially for the chapters read conversion)

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
