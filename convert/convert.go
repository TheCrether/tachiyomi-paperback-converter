package convert

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert/paperbackConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/convert/tachiyomiConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
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
