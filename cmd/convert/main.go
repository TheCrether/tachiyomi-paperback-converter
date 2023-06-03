package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert"
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
	out, err := convert.ConvertPaperback(in)
	if err != nil {
		log.Fatalln("Error converting backup:", err)
	}
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
	out, err := convert.ConvertTachiyomi(bytes.NewReader(in))
	if err != nil {
		log.Fatalln("Error converting backup:", err)
	}
	err = os.WriteFile(os.Args[3], out, 0644)
	if err != nil {
		log.Fatalln("Error writing file:", err)
	}
}
