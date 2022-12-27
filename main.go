package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheCrether/tachiyomi-paperback-converter/convert"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"google.golang.org/protobuf/proto"
)

func main() {
	log.Println(convert.PaperbackToTachiyomi)
	log.Println(convert.TachiyomiToPaperback)
	os.Exit(0)
	if len(os.Args) != 3 {
		log.Fatalf("Usage:  %s <tachiyomi> <paperback>\n", os.Args[0])
	}
	fmt.Println("Hello, World!")
	in, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	backup := &tachiyomi.Backup{}
	if err := proto.Unmarshal(in, backup); err != nil {
		log.Fatalln("Failed to parse backup:", err)
	}

	log.Println(backup.BackupManga[0])
	jsonOut, err := json.MarshalIndent(backup, "", "  ")
	if err != nil {
		log.Fatalln("Failed to marshal backup:", err)
	}
	os.WriteFile("test-data/tachiyomi.json", jsonOut, 0644)

	log.Println("\nPAPERBACK")

	in, err = os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	pbBackup := &paperback.Backup{}
	if err := json.Unmarshal(in, pbBackup); err != nil {
		log.Fatalln("Failed to parse backup:", err)
	}
	log.Println(pbBackup.Library[0])
}
