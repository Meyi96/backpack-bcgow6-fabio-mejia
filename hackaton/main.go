package main

import (
	"fmt"

	"github.com/bootcamp-go/hackaton-go-bases/internal/file"
	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

func main() {
	file := file.File{Path: "./tickets.csv"}
	tickets, _ := file.Read()
	booking := service.NewBookings(tickets)
	//Create ticket
	ticket, err := booking.Create(service.Ticket{
		Name:        "Laura",
		Email:       "laua@algo.com",
		Destination: "Medellin",
		Date:        "2:18",
		Price:       129,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("CREATE\t Se agrego con exito el ticket: %+v\n", ticket)

	//Read ticket
	id := 5
	ticket, err = booking.Read(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("READ\t El ticket con id %d es: %+v\n", id, ticket)

	//update ticket
	ticket, err = booking.Update(5, service.Ticket{
		Id:          5,
		Name:        "Diana",
		Email:       "diana@algo.com",
		Destination: "Cali",
		Date:        "17:11",
		Price:       231,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("UPDATE\t Se actualizo con exito el ticket: %+v\n", ticket)

	//Delete ticket
	id = 995
	idDeleted, err := booking.Delete(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("DELETE\t Se elimino con exito el ticket con id: %d\n", idDeleted)
	file.Write(booking.ReadAll())
}
