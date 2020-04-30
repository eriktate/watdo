package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/eriktate/wrkhub/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type ConfigOpt func(s *Store) *Store

// A Store implements the various *Store interfaces with a postgres backend.
type Store struct {
	db *sqlx.DB

	host     string
	dbname   string
	user     string
	password string
	sslmode  string
	port     uint
	retries  int
}

func WithHost(host string) ConfigOpt {
	return func(s *Store) *Store {
		s.host = host
		return s
	}
}

func WithDB(dbname string) ConfigOpt {
	return func(s *Store) *Store {
		s.dbname = dbname
		return s
	}
}

func WithUser(user string) ConfigOpt {
	return func(s *Store) *Store {
		s.user = user
		return s
	}
}

func WithPassword(password string) ConfigOpt {
	return func(s *Store) *Store {
		s.password = password
		return s
	}
}

func WithPort(port uint) ConfigOpt {
	return func(s *Store) *Store {
		s.port = port
		return s
	}
}

func WithSSLMode(sslmode string) ConfigOpt {
	return func(s *Store) *Store {
		s.sslmode = sslmode
		return s
	}
}

func WithRetries(retries int) ConfigOpt {
	return func(s *Store) *Store {
		s.retries = retries
		return s
	}
}

// DB returns the underlying database pointer used by the Store.
func (s *Store) DB() *sqlx.DB {
	return s.db
}

func (s *Store) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", s.host, s.port, s.dbname, s.user, s.password, s.sslmode)
}

// New returns a new Store with a postgres connection.
func New(opts ...ConfigOpt) (*Store, error) {
	store := &Store{
		host:     env.GetString("WRKHUB_DB_HOST", "localhost"),
		dbname:   env.GetString("WRKHUB_DB_NAME", "wrkhub"),
		user:     env.GetString("WRKHUB_DB_USER", "wrkhub"),
		password: env.GetString("WRKHUB_DB_PASSWORD", "password"),
		port:     env.GetUint("WRKHUB_DB_PORT", 5432),
		sslmode:  env.GetString("WRKHUB_DB_SSL_MODE", "disable"),
		retries:  env.GetInt("RETRIES", 5),
	}

	for _, opt := range opts {
		store = opt(store)
	}

	db, err := attemptConnect(store.ConnectionString(), store.retries)
	if err != nil {
		return nil, err
	}

	store.db = db
	return store, nil
}

func attemptConnect(dsn string, retries int) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		if retries > 0 {
			time.Sleep(5 * time.Second)
			return attemptConnect(dsn, retries-1)
		}

		return nil, fmt.Errorf("failed to connect to postgres after %d attempts: %w", retries, err)
	}

	return db, nil
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
