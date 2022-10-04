package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int       `json:"id"`
	Name         string    `json:"name" binding:"required"`
	LastName     string    `json:"lastName" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Age          int       `json:"age" binding:"required"`
	Height       int       `json:"height" binding:"required"`
	Active       bool      `json:"active"`
	CreationDate time.Time `json:"creationDate"`
}

var users []User

func main() {
	var err error
	users, err = loadData()
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	productGroup := router.Group("/users")
	productGroup.GET("/", getAll)
	productGroup.POST("/", createUser)
	router.Run()
}

func createUser(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "12345" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "does not have permissions to perform the requested request"})
		return
	}
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	addUser(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func getAll(c *gin.Context) {
	var temp []User
	for _, user := range users {
		if (c.Query("Name") == "" || c.Query("Name") == user.Name) &&
			(c.Query("LastName") == "" || c.Query("LastName") == user.LastName) &&
			(c.Query("Email") == "" || c.Query("Email") == user.Email) &&
			(c.Query("Age") == "" || c.Query("Age") == strconv.Itoa(user.Age)) &&
			(c.Query("Height") == "" || c.Query("Height") == strconv.Itoa(user.Height)) &&
			(c.Query("Active") == "" || c.Query("Active") == strconv.FormatBool(user.Active)) {
			temp = append(temp, user)
		}
	}
	c.JSON(200, gin.H{"users": temp})
}

func loadData() (data []User, err error) {
	file, err := os.ReadFile("./Go web/users.json")
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &data)
	return
}

func addUser(u *User) {
	u.Id = users[len(users)-1].Id + 1
	u.CreationDate = time.Now()
	u.Active = true
	users = append(users, *u)
}
