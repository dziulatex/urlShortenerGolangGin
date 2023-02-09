package requests

import (
	"github.com/gin-gonic/gin"
	"goUrlShortener/models"
	"goUrlShortener/repository/short_entity_repository"
	"goUrlShortener/repository/short_entity_visit_repository"
	"log"
	"net/http"
	"time"
)

// GetSingleShortenUrl
// @Accept application/json
// @Description gets single shorten url by id
// @Summary Get single shortenUrl by id
// @Param accessKey path string false "accessKey"
// @Success 302
// @Router  /de-short/{id} [get]
func GetSingleShortenUrl() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		accessKey, _ := c.GetQuery("accessKey")
		shortedEntity, error := checkIfShortenEntityExists(c, id)
		if error != nil {
			log.Fatal(error.Error())
			return
		}
		if shortedEntity.AccessKey == nil || accessKey == *shortedEntity.AccessKey {
			short_entity_visit_repository.Save(&models.ShortenEntityVisitEntity{ShortedUrl: shortedEntity.ShortedUrl})
			shortedEntity.LastAccessedDate = &models.JSONTime{Time: time.Now()}
			short_entity_repository.Save(&shortedEntity)
			c.Redirect(http.StatusFound, shortedEntity.UrlToShorten)
		}
		if accessKey != *shortedEntity.AccessKey {
			c.IndentedJSON(http.StatusForbidden, "Wrong access key")
			return
		}

	}
}

func checkIfShortenEntityExists(c *gin.Context, id string) (models.ShortenEntity, error) {
	shortedEntity, err := short_entity_repository.GetSingle(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "entity_not_found")
		return models.ShortenEntity{}, err
	}
	return shortedEntity, err
}
