package handler

import (
	"net/http"
	"strconv"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/internal/domain"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/internal/users"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Height   int    `json:"height" binding:"required"`
}

type RequestUpdate struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Height   int    `json:"height" binding:"required"`
	Active   *bool  `json:"active" binding:"required"`
}

type RequestUpdateLastNameAndAge struct {
	LastName string `json:"lastName" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type User struct {
	service users.Service
}

func NewUser(service users.Service) *User {
	return &User{service: service}
}

// ListUser godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 404 {object} web.Response
// @Router /users [get]
func (u *User) GetAll(c *gin.Context) {
	users, err := u.service.GetAll(c)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, users, ""))
}

func (u *User) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := u.service.Get(c, id)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, user, ""))
}

// StoreUser godoc
// @Summary Store user
// @Tags Users
// @Description store user
// @accept json
// @Produce  json
// @Param token header string true "token"
// @Param product body Request true "User to store"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 500 {object} web.Response
// @Router /users [post]
func (u *User) Store(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(web.NewResponse(http.StatusBadRequest, nil, err.Error()))
		return
	}
	user := domain.User{
		Name:     request.Name,
		LastName: request.LastName,
		Email:    request.Email,
		Age:      request.Age,
		Height:   request.Height,
	}
	result, err := u.service.Store(c, user)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, result, ""))
}

// UpdateUser godoc
// @Summary Update user
// @Tags Users
// @Description update user
// @accept json
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Param product body RequestUpdate true "User to update"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 404 {object} web.Response
// @failure 500 {object} web.Response
// @Router /users/{id} [put]
func (u *User) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var request RequestUpdate
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
		return
	}
	user := domain.User{
		Name:     request.Name,
		LastName: request.LastName,
		Email:    request.Email,
		Age:      request.Age,
		Height:   request.Height,
		Active:   *request.Active,
	}
	result, err := u.service.Update(c, id, user)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, result, ""))
}

// DeleteUser godoc
// @Summary Delete user
// @Tags Users
// @Description delete user
// @Produce  json
// @Param token header string true "token"
// @Param id path string true "User id"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 404 {object} web.Response
// @Router /users/{id} [delete]
func (u *User) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := u.service.Delete(c, id)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusNotFound, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusNoContent, nil, ""))
}
