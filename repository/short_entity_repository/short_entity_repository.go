package short_entity_repository

import (
	db "goUrlShortener/database"
	"goUrlShortener/models"
	"time"
)

func GetSingle(id string) (models.ShortenEntity, error) {
	connection := db.GetDatabaseConnection()
	shortedUrl := models.ShortenEntity{}
	err := connection.Where("shorted_url=?", id).First(&shortedUrl).Error
	return shortedUrl, err
}
func Save(entity *models.ShortenEntity) {
	connection := db.GetDatabaseConnection()
	connection.Save(&entity)
}
func Delete(entity *models.ShortenEntity) {
	connection := db.GetDatabaseConnection()
	connection.Delete(&entity)
}
func GetAllExpiredShortenEntities() []models.ShortenEntity {
	connection := db.GetDatabaseConnection()
	var expiredShortenUrls []models.ShortenEntity
	currentDate := time.Now().Format(time.DateOnly)
	_ = connection.Where("expire_date<?", currentDate).Find(&expiredShortenUrls)
	return expiredShortenUrls
}
