package models

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var database *gorm.DB
	var err error

	dbType := os.Getenv("DB_TYPE")

	switch dbType {
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
			os.Getenv("POSTGRES_ADDRESS"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("TIMEZONE"))
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		database, err = gorm.Open(sqlite.Open(os.Getenv("SQLITE_DATABASE")), &gorm.Config{})
	default:
		slog.Error("Unsupported DB_TYPE. Use 'postgres' or 'sqlite'.", err)
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&ProjectDesc{})
	database.AutoMigrate(&ProjectQuota{})
	database.AutoMigrate(&ProjectQuotaUsage{})
	database.AutoMigrate(&ServerDesc{})
	database.AutoMigrate(&ServerSpec{})
	database.AutoMigrate(&ServerUsage{})
	database.AutoMigrate(&ServerOwnership{})
	database.AutoMigrate(&FlavorDesc{})
	database.AutoMigrate(&FlavorSpec{})
	database.AutoMigrate(&UserDesc{})
	database.AutoMigrate(&UserProject{})

	DB = database
}
