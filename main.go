package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/TheCrether/tachiyomi-paperback-converter/models/paperback"
	"github.com/TheCrether/tachiyomi-paperback-converter/models/tachiyomi"
	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage:  %s ADDRESS_BOOK_FILE\n", os.Args[0])
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

	log.Println("PAPERBACK\n")

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
