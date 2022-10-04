package handler

import (
	"net/http"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-web/Clase2.2/internal/users"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Height   int    `json:"height" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(service users.Service) *User {
	return &User{service: service}
}

func (u *User) GetAll(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "12345" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "does not have permissions to perform the requested request"})
		return
	}
	users, err := u.service.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *User) Store(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "12345" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "does not have permissions to perform the requested request"})
		return
	}
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := u.service.Store(request.Name, request.LastName, request.Email, request.Age, request.Height)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}
