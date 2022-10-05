package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase3.1/cmd/server/handler"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase3.1/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Json
	data, err := loadData()
	if err != nil {
		fmt.Println(err)
		return
	}
	repo := users.NewRepository(data)
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

func loadData() (data []users.User, err error) {
	file, err := os.ReadFile("./Go-web/users.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &data)
	return
}
