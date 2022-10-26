package database

import (
	"fmt"
	"os"
	"path"

	"github.com/shanmugharajk/vault/internal/models"
	"github.com/shanmugharajk/vault/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const DB_NAME = "vault.store"

var dbPath string

func init() {
	dbPath = utils.GetConfig().DbPath
}

func Connect() {
	var err error

	dbFilePath := path.Join(dbPath, DB_NAME)
	Db, err = gorm.Open(sqlite.Open(dbFilePath), &gorm.Config{
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
	os.Remove(DB_NAME)
	Connect()
	Migrate()
}
