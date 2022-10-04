package main

import (
	"encoding/json"
	"os"
	"strconv"
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
	router.GET("/users", getAll)
	router.GET("/users/:id", findById)

	router.Run()
}

func getAll(c *gin.Context) {
	data, err := loadData()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	var users []User
	for _, user := range data {
		if (c.Query("Name") == "" || c.Query("Name") == user.Name) &&
			(c.Query("LastName") == "" || c.Query("LastName") == user.LastName) &&
			(c.Query("Email") == "" || c.Query("Email") == user.Email) &&
			(c.Query("Age") == "" || c.Query("Age") == strconv.Itoa(user.Age)) &&
			(c.Query("Height") == "" || c.Query("Height") == strconv.Itoa(user.Height)) &&
			(c.Query("Active") == "" || c.Query("Active") == strconv.FormatBool(user.Active)) {
			users = append(users, user)
		}
	}
	c.JSON(200, gin.H{"users": users})
}

func findById(c *gin.Context) {
	data, err := loadData()
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	for _, user := range data {
		if c.Param("id") == strconv.Itoa(user.Id) {
			c.JSON(200, gin.H{"user": user})
			return
		}
	}
	c.JSON(404, "Not found")
}

func loadData() (data []User, err error) {
	file, err := os.ReadFile("./Go web/users.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &data)
	return
}
