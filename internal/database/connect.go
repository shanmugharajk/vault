package database

import (
	"fmt"
	"os"

	"github.com/shanmugharajk/vault/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DSN = "vault.store"

func Connect() {
	var err error
	Db, err = gorm.Open(sqlite.Open(DSN), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// This is needed to enable foreign key constraints.
	Db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	Db.AutoMigrate(&models.Secret{})
	fmt.Println("Database Migrated")
}

func Recreate() {
	os.Remove(DSN)
	Connect()
	Migrate()
}
