package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/gookit/cache"
	"github.com/gookit/cache/redis"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goUrlShortener/crons"
	docs "goUrlShortener/docs"
	"goUrlShortener/requests"
	"goUrlShortener/utils"
	"goUrlShortener/validators"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func main() {
	r := setupRouter()
	r.Run()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/docs"
	r.POST("/shorten", requests.CreateShortenUrl())
	r.GET("/de-short/:id", requests.GetSingleShortenUrl())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/debug", requests.DebugRequest())
	ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
	)
	return r
}

func init() {
	bindCustomValidators()
	registerCustomSerializers()
	loadEnv()
	registerCache()
	go crons.RunCrons()
}
func registerCustomSerializers() {
	schema.RegisterSerializer("sqlDateTimeToGoTime", utils.SQLDateTimeToGoTime{})
}
func registerCache() {
	cache.Register(os.Getenv("CACHE_DRIVER"), redis.Connect(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PWD"), 0))
	cache.DefaultUse(os.Getenv("CACHE_DRIVER"))
}
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func bindCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isDateGtThanToday", validators.IsDateGtThanToday)
		v.RegisterValidation("isDateGtThanTodayOrNull", validators.IsDateGtThanTodayOrNull)
	}
}
