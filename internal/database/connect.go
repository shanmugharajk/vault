package database

import (
	"fmt"
	"os"

	"github.com/shanmugharajk/vault/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Automigrate bool
	Recreate    bool
}

const DSN = "vault.db"

func Connect(config *Config) {
	if config.Recreate {
		os.Remove(DSN)
	}

	var err error
	Db, err = gorm.Open(sqlite.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// This is needed to enable foreign key constraints.
	Db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

	if config.Automigrate {
		Db.AutoMigrate(&models.Secret{})
	}

	fmt.Println("Database Migrated")
}
