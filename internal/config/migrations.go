package config

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)


func RunMigrations(cfg *Config) error {
	db, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName))
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func _migrateDown(cfg *Config) error {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName))
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func _migrateToVersion(version uint,cfg *Config) error {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName))
	if err != nil {
		return err
	}
	if err := m.Migrate(version); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}
