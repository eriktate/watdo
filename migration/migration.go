package migration

import (
	"fmt"
	"io/ioutil"

	"github.com/eriktate/wrkhub/postgres"
	_ "github.com/lib/pq"
)

func migrate(filename string, opts ...postgres.ConfigOpt) error {
	migration, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to load migration %s: %w", filename, err)
	}

	store, err := postgres.New(opts...)
	if err != nil {
		return fmt.Errorf("failed to initialize store for migration: %w", err)
	}

	if _, err := store.DB().Exec(string(migration)); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	return nil
}

func MigrateUp(opts ...postgres.ConfigOpt) error {
	return migrate("migration/up.sql", opts...)
}

func MigrateDown(opts ...postgres.ConfigOpt) error {
	return migrate("migration/down.sql", opts...)
}
