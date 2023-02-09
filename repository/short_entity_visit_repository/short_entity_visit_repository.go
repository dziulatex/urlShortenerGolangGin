package short_entity_visit_repository

import (
	db "goUrlShortener/database"
	"goUrlShortener/models"
)

func GetSingle(id string) (models.ShortenEntity, error) {
	connection := db.GetDatabaseConnection()
	shortedUrl := models.ShortenEntity{}
	err := connection.Where("shorted_url=?", id).First(&shortedUrl).Error
	return shortedUrl, err
}
func Save(entity *models.ShortenEntityVisitEntity) {
	connection := db.GetDatabaseConnection()
	connection.Save(&entity)
}
