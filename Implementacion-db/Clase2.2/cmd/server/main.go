package main

import (
	"log"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/cmd/server/handler"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/internal/users"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/pkg/db"
	"github.com/gin-gonic/gin"
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
	db := db.ConnecDatabase()
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	user := handler.NewUser(service)

	router := gin.Default()

	productGroup := router.Group("/users")
	productGroup.Use(handler.TokenAuthMiddleware)
	productGroup.GET("/", user.GetAll)
	productGroup.GET("/:id", user.Get)
	productGroup.POST("/", user.Store)
	productGroup.PUT("/:id", handler.IdValidationMiddleware, user.Update)
	productGroup.DELETE("/:id", handler.IdValidationMiddleware, user.Delete)
	err := router.Run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
