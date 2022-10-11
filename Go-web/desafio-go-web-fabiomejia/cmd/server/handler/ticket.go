package handler

import (
	"net/http"

	"desafio-go-web-fabio-mejia/internal/tickets"
	"desafio-go-web-fabio-mejia/pkg/web"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	service tickets.Service
}

func NewTicket(s tickets.Service) *Ticket {
	return &Ticket{service: s}
}

// GetTicketsByCountry godoc
// @Summary Get Ticket by country
// @Tags Tickets
// @Description get list a tickets from a country
// @Produce  json
// @Param token header string true "token"
// @Param dest path string true "destination"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 500 {object} web.Response
// @Router /tickets/getByCountry/{dest} [get]
func (t *Ticket) GetTicketsByCountry(c *gin.Context) {
	destination := c.Param("dest")
	tickets, err := t.service.GetTicketsByCountry(destination)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, tickets, ""))
}

// GetTotalTicketsByCountry godoc
// @Summary Get total tickets by country
// @Tags Tickets
// @Description get total tickets from a country
// @Produce  json
// @Param token header string true "token"
// @Param dest path string true "destination"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 500 {object} web.Response
// @Router /tickets/getTotalByCountry/{dest} [get]
func (t *Ticket) GetTotalTicketsByCountry(c *gin.Context) {
	destination := c.Param("dest")
	total, err := t.service.GetTotalTicketsByCountry(destination)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, gin.H{"total": total}, ""))
}

// AverageDestination godoc
// @Summary Get average tickets per hour in a country
// @Tags Tickets
// @Description get average tickets per hour in a country
// @Produce  json
// @Param token header string true "token"
// @Param dest path string true "destination"
// @Success 200 {object} web.Response
// @failure 401 {object} web.Response
// @failure 400 {object} web.Response
// @failure 500 {object} web.Response
// @Router /tickets/getAverage/{dest} [get]
func (t *Ticket) AverageDestination(c *gin.Context) {
	destination := c.Param("dest")
	avg, err := t.service.AverageDestination(destination)
	if err != nil {
		c.JSON(web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		return
	}
	c.JSON(web.NewResponse(http.StatusOK, gin.H{"average": avg}, ""))
}
