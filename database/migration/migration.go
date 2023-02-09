package main

import (
	db "goUrlShortener/database"
	"goUrlShortener/models"
)

func main() {
	connection := db.GetDatabaseConnection()
	connection.AutoMigrate(&models.ShortenEntity{}, &models.ShortenEntityVisitEntity{})
}
