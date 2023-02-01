package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASS") + " dbname=" + os.Getenv("DB_SCHEMA") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
