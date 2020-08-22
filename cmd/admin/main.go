package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/qreasio/go-starter-kit/internal/config"
	"github.com/qreasio/go-starter-kit/pkg/log"

	_ "github.com/go-sql-driver/mysql"
	migrater "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New()
	if len(os.Args) <= 3 {
		logger.Error("Usage:", os.Args[1], "command", "argument")
		return errors.New("invalid command")
	}

	cfg, err := config.Load(os.Args[2], logger)
	if err != nil {
		logger.Errorf("failed to load application configuration: %s", os.Args[2])
		return err
	}

	switch os.Args[1] {
	case "migrate":
		err = Migrate(cfg, logger, os.Args[3])
	case "seed":
		err = Seed(cfg, logger, os.Args[3])
	default:
		err = errors.New("must specify a command")
	}

	if err != nil {
		return err
	}

	return nil
}

// Migrate to run database migration up or down
func Migrate(cfg *config.Config, logger log.Logger, command string) error {
	db, err := sql.Open("mysql", cfg.DB.Dsn)
	if err != nil {
		logger.Error(err)
		return err
	}
	path, err := os.Getwd()
	if err != nil {
		logger.Error(err)
		return err
	}

	migrationPath := fmt.Sprintf("file://%s/migration", path)
	logger.Infof("migrationPath : %s", migrationPath)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logger.Error(err)
		return err
	}
	m, err := migrater.NewWithDatabaseInstance(
		migrationPath,
		"mysql",
		driver,
	)
	if err != nil {
		logger.Error(err)
		return err
	}
	if command == "up" {
		logger.Info("Migrate up")
		if err := m.Up(); err != nil && err != migrater.ErrNoChange {
			logger.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if command == "down" {
		logger.Info("Migrate down")
		if err := m.Down(); err != nil && err != migrater.ErrNoChange {
			logger.Errorf("An error occurred while syncing the database.. %v", err)
			return err
		}
	}

	if err != nil {
		logger.Error(err)
		return err
	}

	logger.Info("Migrate complete")
	return nil
}

// Seed to populate database with seed data
func Seed(cfg *config.Config, logger log.Logger, sqlFilename string) error {
	db, err := sqlx.Open("mysql", cfg.DB.Dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	// load from SQL file
	bytes, err := ioutil.ReadFile(sqlFilename)
	if err != nil {
		return err
	}
	sql := string(bytes)
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(sql); err != nil {
		logger.Error(err)
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("Seed data complete")
	return nil
}
