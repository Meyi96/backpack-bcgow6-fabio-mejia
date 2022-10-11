package handler

import (
	"net/http"
	"os"

	"desafio-go-web-fabio-mejia/pkg/web"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware(c *gin.Context) {
	token := c.GetHeader("token")
	if token != os.Getenv("TOKEN") {
		c.AbortWithStatusJSON(web.NewResponse(http.StatusUnauthorized, nil, "does not have permissions to perform the requested request"))
		return
	}
	c.Next()
}

func DestParamValidationMiddleware(c *gin.Context) {
	if destination := c.Param("dest"); destination == "" {
		c.AbortWithStatusJSON(web.NewResponse(http.StatusBadRequest, nil, "invalid destination"))
		return
	}
	c.Next()
}
