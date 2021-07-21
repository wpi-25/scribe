package db

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func getConfig() string {
	if db_host := os.Getenv("DATABASE_URL"); db_host != "" {
		return db_host
	}
	db_pass := os.Getenv("POSTGRES_PASSWORD")
	return fmt.Sprintf("host=localhost user=scribe password=%s dbname=scribe port=5432 sslmode=disable", db_pass)
}

func Connect() error {
	db, _ := sqlx.Open("postgres", getConfig())

	err := db.Ping()
	if err != nil {
		return err
	}

	DB = db
	err = runMigrations()
	if err != nil {
		return err
	}
	return nil
}

func runMigrations() error {
	log.Println("Running migrations")
	driver, err := postgres.WithInstance(DB.DB, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance("file://./migrations", "postgres", driver)
	if err != nil {
		return err
	}
	m.Steps(2)
	return nil
}
