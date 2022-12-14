package main

import (
	"log"
	"os"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.1/cmd/server/handler"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.1/docs"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.1/internal/users"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.1/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const JsonPath string = "./users.json"

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore(store.FileType, JsonPath)
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	user := handler.NewUser(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productGroup := router.Group("/users")
	productGroup.Use(handler.TokenAuthMiddleware)
	productGroup.GET("/", user.GetAll)
	productGroup.POST("/", user.Store)
	productGroup.PUT("/:id", handler.IdValidationMiddleware, user.Update)
	productGroup.PATCH("/:id", handler.IdValidationMiddleware, user.UpdateLastNameAndAge)
	productGroup.DELETE("/:id", handler.IdValidationMiddleware, user.Delete)
	err = router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
