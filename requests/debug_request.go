package requests

import (
	"github.com/gin-gonic/gin"
	"goUrlShortener/repository/short_entity_repository"
	"net/http"
	"os"
)

func DebugRequest() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != os.Getenv("INTERNAL_TOKEN") {
			c.IndentedJSON(http.StatusForbidden, "forbidden")
			return
		}
		expiredShortenEntities := short_entity_repository.GetAllExpiredShortenEntities()
		for _, expiredShortenEntity := range expiredShortenEntities {
			short_entity_repository.Delete(&expiredShortenEntity)
		}
	}
}
