package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/mcutalo88/shurly/internal/config"
)

func New(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Db,
		cfg.Database.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Errorf("error connecting to the database: %s", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("unable to ping database: %s", err))
	}

	return db
}
