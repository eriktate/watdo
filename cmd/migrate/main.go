package main

import (
	"log"
	"os"

	"github.com/eriktate/wrkhub/migration"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "up":
		if err := migration.MigrateUp(); err != nil {
			log.Fatal(err)
		}
		log.Println("Migrated UP successfully!")
	case "down":
		if err := migration.MigrateDown(); err != nil {
			log.Fatal(err)
		}
		log.Println("Migrated DOWN successfully!")
	default:
		log.Fatalf("unrecognized migration command: %s", cmd)
	}
}
