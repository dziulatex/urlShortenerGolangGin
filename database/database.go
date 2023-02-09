package database

import (
	"goUrlShortener/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func GetDatabaseConnection() gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.ShortenEntity{})
	if err != nil {
		panic("failed to migrate")
	}
	return *db
}
