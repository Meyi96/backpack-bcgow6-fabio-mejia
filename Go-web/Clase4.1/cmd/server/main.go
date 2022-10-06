package main

import (
	"log"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase4.1/cmd/server/handler"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase4.1/internal/users"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase4.1/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const JsonPath string = "./users.json"

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
	productGroup := router.Group("/users")
	productGroup.GET("/", user.GetAll)
	productGroup.POST("/", user.Store)
	productGroup.PUT("/:id", user.Update)
	productGroup.PATCH("/:id", user.UpdateLastNameAndAge)
	productGroup.DELETE("/:id", user.Delete)
	router.Run()
}
