package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitDb() (*sqlx.DB, error) {
	var db *sqlx.DB
	db, err := sqlx.Open("postgres", os.Getenv("DATABASE_URI"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)

	duration, _ := time.ParseDuration("15m")
	db.SetConnMaxIdleTime(duration)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connected!")

	// apply the migrations
	runMigrations(db.DB)

	return db, nil
}

func runMigrations(db *sql.DB) {
	// create the migration driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	// initialize the migrations
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // directory for the migrations
		"gomicroservicedb",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to initialize the migrations: %v", err)
	}

	// apply the migrations
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply the migrations: %v", err)
	}

	log.Println("Migrations are updated!")
}
