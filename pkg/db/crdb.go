package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/mcutalo88/shurly/pkg/config"
)

func New(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/shurly?sslmode=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.SSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error connecting to the database: ", err)
	}

	return db
}
