package main

import (
	"log"
	"os"

	"desafio-go-web-fabio-mejia/cmd/server/handler"
	"desafio-go-web-fabio-mejia/docs"
	"desafio-go-web-fabio-mejia/internal/tickets"
	"desafio-go-web-fabio-mejia/pkg"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Tickets.
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
	data, err := pkg.LoadTicketsFromFile("tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}

	repo := tickets.NewRepository(data)
	service := tickets.NewService(repo)
	ticket := handler.NewTicket(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ticketGroup := router.Group("/tickets")
	ticketGroup.Use(handler.TokenAuthMiddleware)
	ticketGroup.GET("/getByCountry/:dest", handler.DestParamValidationMiddleware, ticket.GetTicketsByCountry)
	ticketGroup.GET("/getTotalByCountry/:dest", handler.DestParamValidationMiddleware, ticket.GetTotalTicketsByCountry)
	ticketGroup.GET("/getAverage/:dest", handler.DestParamValidationMiddleware, ticket.AverageDestination)

	if err := router.Run(); err != nil {
		panic(err)
	}

}
