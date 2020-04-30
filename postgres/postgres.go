package postgres

import (
	"context"
	"fmt"

	"github.com/eriktate/watdo/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// A Store implements the various *Store interfaces with a postgres backend.
type Store struct {
	db *sqlx.DB
}

type StoreOpts struct {
	host     string
	dbname   string
	user     string
	password string
	sslMode  string
	port     uint
}

func NewStoreOpts() StoreOpts {
	return StoreOpts{
		host:     env.GetString("WATDO_DB_HOST", "localhost"),
		dbname:   env.GetString("WATDO_DB_NAME", "watdo"),
		user:     env.GetString("WATDO_DB_USER", "watdo"),
		password: env.GetString("WATDO_DB_PASSWORD", "password"),
		port:     env.GetUint("WATDO_DB_PORT", 5432),
		sslMode:  env.GetString("WATDO_DB_SSL_MODE", "disable"),
	}
}

func (c StoreOpts) WithHost(host string) StoreOpts {
	c.host = host
	return c
}

func (c StoreOpts) WithDB(dbname string) StoreOpts {
	c.dbname = dbname
	return c
}

func (c StoreOpts) WithUser(user string) StoreOpts {
	c.user = user
	return c
}

func (c StoreOpts) WithPassword(password string) StoreOpts {
	c.password = password
	return c
}

func (c StoreOpts) WithPort(port uint) StoreOpts {
	c.port = port
	return c
}

func (c StoreOpts) WithSSLMode(sslMode string) StoreOpts {
	c.sslMode = sslMode
	return c
}

func (c StoreOpts) String() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", c.host, c.port, c.dbname, c.user, c.password, c.sslMode)
}

// New returns a new Store with a postgres connection.
func New(opts StoreOpts) (*Store, error) {
	db, err := sqlx.Connect("postgres", opts.String())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return &Store{db}, nil
}

func runNamedTx(ctx context.Context, db *sqlx.DB, query string, arg interface{}) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	if _, err := tx.NamedExecContext(ctx, query, arg); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
