package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage:  %s <output-type: tachiyomi/paperback> <input> <output>\n", os.Args[0])
	}

	if strings.ToLower(os.Args[1]) == "tachiyomi" {
		in, err := os.ReadFile(os.Args[2])
		if err != nil {
			log.Fatalln("Error reading file:", err)
		}

		var backup *paperback.Backup
		err = json.Unmarshal(in, &backup)
		if err != nil {
			log.Fatalln("Error unmarshalling json:", err)
		}
		tachiyomiBackup, err := convert.ConvertPaperbackToTachiyomi(backup)
		if err != nil {
			log.Fatalln("Error converting backup:", err)
		}
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
	} else if strings.ToLower(os.Args[1]) == "paperback" {
		in, err := os.ReadFile(os.Args[2])
		if err != nil {
			log.Fatalln("Error reading file:", err)
		}
		var backup *tachiyomi.Backup
		err = proto.Unmarshal(in, backup)
		if err != nil {
			log.Fatalln("Error unmarshalling proto:", err)
		}
		paperbackBackup, err := convert.ConvertTachiyomiToPaperback(backup)
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
	} else {
		log.Fatalf("Usage:  %s <tachiyomi/paperback> <input> <output>\n", os.Args[0])
	}
}
