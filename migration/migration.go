package migration

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/eriktate/wrkhub/postgres"
	_ "github.com/lib/pq"
)

func migrate(filename string, opts postgres.StoreOpts, retries int) error {
	migration, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load migration %s: %w", filename, err)
	}

	db, err := sql.Open("postgres", opts.String())
	if err != nil {
		if retries > 0 {
			time.Sleep(5 * time.Second)
			return migrate(filename, opts, retries-1)
		}
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if _, err := db.Exec(string(migration)); err != nil {
		if retries > 0 {
			time.Sleep(5 * time.Second)
			return migrate(filename, opts, retries-1)
		}
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}

func MigrateUp(opts postgres.StoreOpts, retries int) error {
	return migrate("migration/up.sql", opts, retries)
}

func MigrateDown(opts postgres.StoreOpts, retries int) error {
	return migrate("migration/down.sql", opts, retries)
}
