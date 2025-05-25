package convert

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert/paperbackConvert"
	paperbackv8Convert "github.com/TheCrether/tachiyomi-paperback-converter/convert/paperbackV8Convert"
	"github.com/TheCrether/tachiyomi-paperback-converter/convert/tachiyomiConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperbackv8"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"google.golang.org/protobuf/proto"
)

type ConvertError struct {
	Msg string
	Err error
}

func (e *ConvertError) Error() string {
	return e.Msg + "\ndetailed: " + e.Err.Error()
}

func convError(msg string, err error) error {
	return &ConvertError{msg, err}
}

func ConvertTachiyomi(reader *bytes.Reader) ([]byte, error) {
	g, err := gzip.NewReader(reader)
	if err != nil {
		return []byte{}, convError("Error decompressing backup:", err)
	}
	var b bytes.Buffer
	if _, err := b.ReadFrom(g); err != nil {
		return []byte{}, convError("Error reading decompressing backup:", err)
	}
	decompressed := b.Bytes()

	backup := &tachiyomi.Backup{}
	err = proto.Unmarshal(decompressed, backup)
	if err != nil {
		return []byte{}, convError("Error unmarshalling proto:", err)
	}

	paperbackBackup, err := paperbackConvert.ConvertTachiyomiToPaperback(backup)
	if err != nil {
		log.Fatalln("Error converting backup:", err)
	}

	out, err := json.Marshal(paperbackBackup)
	if err != nil {
		return []byte{}, convError("Error marshalling json:", err)
	}
	return out, nil
}

func ConvertPaperback(in []byte) ([]byte, error) {
	var backup *paperback.Backup
	err := json.Unmarshal(in, &backup)
	if err != nil {
		return []byte{}, convError("Error unmarshalling json:", err)
	}
	// log.Printf("trying to convert %d Mangas\n", len(backup.Library))
	tachiyomiBackup, err := tachiyomiConvert.ConvertPaperbackToTachiyomi(backup)
	if err != nil {
		return []byte{}, convError("Error converting backup:", err)
	}
	// log.Printf("converted %d mangas\n", len(tachiyomiBackup.BackupManga))
	out, err := proto.Marshal(tachiyomiBackup)
	if err != nil {
		return []byte{}, convError("Error marshalling proto:", err)
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(out); err != nil {
		return []byte{}, convError("(1) Error compressing backup:", err)
	}
	if err := gz.Flush(); err != nil {
		return []byte{}, convError("(2) Error compressing backup:", err)
	}
	if err := gz.Close(); err != nil {
		return []byte{}, convError("(3) Error compressing backup:", err)
	}
	out = b.Bytes()
	return out, nil
}

func ConvertPaperbackV8(zipFile *zip.ReadCloser) ([]byte, error) {
	libraryMangas := map[string]paperbackv8.LibraryManga{}
	mangaInfos := map[string]paperbackv8.MangaInfo{}
	sourceMangas := map[string]paperbackv8.SourceManga{}
	chapters := map[string]paperbackv8.Chapter{}
	chapterProgressMarkers := map[string]paperbackv8.ChapterProgressMarker{}

	for _, f := range zipFile.File {
		name := f.FileInfo().Name()
		// fmt.Println("file: ", name)

		if name == "__LIBRARY_MANGA_V4" {
			convertJSONIntoMap(f, &libraryMangas)
		} else if name == "__MANGA_INFO_V4" {
			convertJSONIntoMap(f, &mangaInfos)
		} else if name == "__SOURCE_MANGA_V4" {
			convertJSONIntoMap(f, &sourceMangas)
		} else if name == "__CHAPTER_V4" {
			convertJSONIntoMap(f, &chapters)
		} else if strings.HasPrefix(name, "__CHAPTER_PROGRESS_MARKER_V4") {
			tempProgressMarkers := map[string]paperbackv8.ChapterProgressMarker{}
			convertJSONIntoMap(f, &tempProgressMarkers)
			for k, v := range tempProgressMarkers {
				chapterProgressMarkers[k] = v
			}
		}
	}

	data := paperbackv8.PaperbackV8Data{
		LibraryMangas:          libraryMangas,
		MangaInfos:             mangaInfos,
		SourceMangas:           sourceMangas,
		Chapters:               chapters,
		ChapterProgressMarkers: chapterProgressMarkers,
	}

	tachiyomiBackup, err := paperbackv8Convert.ConvertPaperbackV8ToTachiyomi(&data)
	if err != nil {
		return []byte{}, convError("Error converting backup:", err)
	}

	fmt.Println("libraryMangas: ", len(libraryMangas))
	fmt.Println("mangaInfos: ", len(mangaInfos))
	fmt.Println("sourceMangas: ", len(sourceMangas))
	fmt.Println("chapters: ", len(chapters))
	fmt.Println("chapterProgressMarkers: ", len(chapterProgressMarkers))

	out, err := proto.Marshal(tachiyomiBackup)
	if err != nil {
		return []byte{}, convError("Error marshalling proto:", err)
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(out); err != nil {
		return []byte{}, convError("(1) Error compressing backup:", err)
	}
	if err := gz.Flush(); err != nil {
		return []byte{}, convError("(2) Error compressing backup:", err)
	}
	if err := gz.Close(); err != nil {
		return []byte{}, convError("(3) Error compressing backup:", err)
	}
	out = b.Bytes()
	return out, nil
}

func convertJSONIntoMap[V any](file *zip.File, output *map[string]V) error {
	reader, err := file.Open()
	if err != nil {
		return err
	}

	jsonDec := json.NewDecoder(reader)
	for {
		err = jsonDec.Decode(output)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	reader.Close()

	return nil
}
