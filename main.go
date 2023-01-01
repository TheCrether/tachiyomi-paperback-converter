package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert/paperbackConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/convert/tachiyomiConvert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage:  %s <output-type: tachiyomi/paperback> <input> <output>\n", os.Args[0])
	}

	if strings.ToLower(os.Args[1]) == "tachiyomi" {
		ConvertPaperback()
	} else if strings.ToLower(os.Args[1]) == "paperback" {
		ConvertTachiyomi()
	} else if strings.ToLower(os.Args[1]) == "tachiyomi-json" {
		ConvertTachiyomiToJSON()
	} else {
		log.Fatalf("Usage: %s <tachiyomi/paperback> <input> <output>\n", os.Args[0])
	}
}

func ConvertTachiyomiToJSON() {
	in, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	g, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		log.Fatalln("Error decompressing backup:", err)
	}
	var b bytes.Buffer
	if _, err := b.ReadFrom(g); err != nil {
		log.Fatalln("Error decompressing backup:", err)
	}
	in = b.Bytes()

	backup := &tachiyomi.Backup{}
	err = proto.Unmarshal(in, backup)
	if err != nil {
		log.Fatalln("Error unmarshalling proto:", err)
	}

	out, err := json.Marshal(backup)
	if err != nil {
		log.Fatalln("Error marshalling json:", err)
	}
	err = os.WriteFile(os.Args[3], out, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}

func ConvertPaperback() {
	in, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	var backup *paperback.Backup
	err = json.Unmarshal(in, &backup)
	if err != nil {
		log.Fatalln("Error unmarshalling json:", err)
	}
	log.Printf("trying to convert %d Mangas\n", len(backup.Library))
	tachiyomiBackup, err := tachiyomiConvert.ConvertPaperbackToTachiyomi(backup)
	if err != nil {
		log.Fatalln("Error converting backup:", err)
	}
	log.Printf("converted %d mangas\n", len(tachiyomiBackup.BackupManga))
	out, err := proto.Marshal(tachiyomiBackup)
	if err != nil {
		log.Fatalln("Error marshalling proto:", err)
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(out); err != nil {
		log.Fatalln("Error compressing backup:", err)
	}
	if err := gz.Flush(); err != nil {
		log.Fatalln("Error compressing backup:", err)
	}
	if err := gz.Close(); err != nil {
		log.Fatalln("Error compressing backup:", err)
	}
	out = b.Bytes()
	if !strings.HasSuffix(os.Args[3], ".gz") {
		os.Args[3] += ".gz"
	}
	err = os.WriteFile(os.Args[3], out, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}

func ConvertTachiyomi() {
	in, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	g, err := gzip.NewReader(bytes.NewReader(in))
	if err != nil {
		log.Fatalln("Error decompressing backup:", err)
	}
	var b bytes.Buffer
	if _, err := b.ReadFrom(g); err != nil {
		log.Fatalln("Error decompressing backup:", err)
	}
	in = b.Bytes()

	backup := &tachiyomi.Backup{}
	err = proto.Unmarshal(in, backup)
	if err != nil {
		log.Fatalln("Error unmarshalling proto:", err)
	}
	paperbackBackup, err := paperbackConvert.ConvertTachiyomiToPaperback(backup)
	if err != nil {
		log.Fatalln("Error converting backup:", err)
	}
	out, err := json.Marshal(paperbackBackup)
	if err != nil {
		log.Fatalln("Error marshalling json:", err)
	}
	err = os.WriteFile(os.Args[3], out, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}
