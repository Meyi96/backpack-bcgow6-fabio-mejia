package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate time.Time
}

func main() {
	router := gin.Default()

	router.GET("/saludar", saludar)
	router.GET("/users", getAll)

	router.Run()
}

func saludar(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hola meyi!"})
}

func getAll(c *gin.Context) {
	file, err := os.ReadFile("./Go web/users.json")
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}
	var users []User
	json.Unmarshal(file, &users)
	c.JSON(200, gin.H{"users": users})
}
