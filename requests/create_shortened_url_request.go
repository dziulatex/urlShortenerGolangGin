package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/cache"
	models "goUrlShortener/models"
	"goUrlShortener/repository/short_entity_repository"
	"goUrlShortener/utils"
	"net/http"
	"strings"
	"time"
)

type shortenPostBodyRequestStruct struct {
	UrlToShorten string `binding:"required,url"`
	AccessKey    string
	ExpireDate   time.Time `time_format:"2006-01-02" binding:"isDateGtThanTodayOrNull"`
}

// CreateShortenUrl
// @Description creates shorten url
// @Accept application/json
// @Param urlToShorten body string true "https://yourbasic.org/golang/structs-explained/"
// @Param expireDate body string false "2023-02-09T00:00:00Z"
// @Param accessKey body string false "randomKeyx"
// @Produce json
// @Success 200 {object} models.ShortenEntityResponse
// @Router /shorten [post]
func CreateShortenUrl() func(c *gin.Context) {
	return func(c *gin.Context) {
		var shortenPostBodyRequest shortenPostBodyRequestStruct
		if err := c.ShouldBindJSON(&shortenPostBodyRequest); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "VALIDATEERR-1",
					"message": err.Error()})
			return
		}
		shortedUrlLen, accessKeyValue, expireDateValue := createShortenUrlParseInputParams(shortenPostBodyRequest)
		newShortenEntity := models.ShortenEntity{UrlToShorten: shortenPostBodyRequest.UrlToShorten, AccessKey: accessKeyValue, ShortedUrl: utils.RandomUrlShortenAndGenerateNewUntilNotUsed(shortedUrlLen), ExpireDate: expireDateValue}
		c.IndentedJSON(http.StatusCreated, newShortenEntity)
		short_entity_repository.Save(&newShortenEntity)
		cache.Del(newShortenEntity.ShortedUrl)
	}
}

func createShortenUrlParseInputParams(shortenPostBodyRequest shortenPostBodyRequestStruct) (int, *string, *models.JSONTime) {
	var shortedUrlLen = utils.GetEnvVarAsInt("SHORTED_URL_LEN")
	shortenPostBodyRequest.AccessKey = strings.TrimSpace(shortenPostBodyRequest.AccessKey)
	var accessKeyValue = &shortenPostBodyRequest.AccessKey
	if len(shortenPostBodyRequest.AccessKey) == 0 {
		accessKeyValue = nil
	}
	var expireDate = &shortenPostBodyRequest.ExpireDate
	expireDateValue := &models.JSONTime{Time: *expireDate}
	if shortenPostBodyRequest.ExpireDate.IsZero() {
		expireDateValue = nil
	}
	return shortedUrlLen, accessKeyValue, expireDateValue
}
