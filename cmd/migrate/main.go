package main

import (
	"log"
	"os"

	"github.com/eriktate/watdo/env"
	"github.com/eriktate/watdo/migration"
	"github.com/eriktate/watdo/postgres"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "up":
		if err := migration.MigrateUp(postgres.NewStoreOpts(), env.GetInt("RETRIES", 6)); err != nil {
			log.Fatal(err)
		}
		log.Println("Migrated UP successfully!")
	case "down":
		if err := migration.MigrateDown(postgres.NewStoreOpts(), env.GetInt("RETRIES", 6)); err != nil {
			log.Fatal(err)
		}
		log.Println("Migrated DOWN successfully!")
	default:
		log.Fatalf("unrecognized migration command: %s", cmd)
	}

}
