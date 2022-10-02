package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/bootcamp-go/hackaton-go-bases/internal/service"
)

type File struct {
	Path string
}

func (f *File) Read() (tickets []service.Ticket, err error) {
	file, err := os.Open(f.Path)
	if err != nil {
		return nil, fmt.Errorf("Unable to read input file "+f.Path, err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	for _, register := range lines {
		ticket := stringToTicket(register)
		tickets = append(tickets, ticket)
	}
	return
}

func (f *File) Write(tickets []service.Ticket) (err error) {
	file, err := os.OpenFile(f.Path, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("Unable to read input file "+f.Path, err)
	}
	defer file.Close()

	writter := csv.NewWriter(file)
	_ = os.Truncate(f.Path, 0)
	for _, ticket := range tickets {
		register := ticketToString(ticket)
		if err := writter.Write(register); err != nil {
			panic(err)
		}
	}
	writter.Flush()
	return
}

func stringToTicket(register []string) (ticket service.Ticket) {
	id, err := strconv.Atoi(register[0])
	if err != nil {
		panic(err)
	}
	price, err := strconv.Atoi(register[5])
	if err != nil {
		panic(err)
	}
	ticket = service.Ticket{
		Id:          id,
		Name:        register[1],
		Email:       register[2],
		Destination: register[3],
		Date:        register[4],
		Price:       price,
	}
	return
}

func ticketToString(ticket service.Ticket) (register []string) {
	register = []string{
		strconv.Itoa(ticket.Id),
		ticket.Name,
		ticket.Email,
		ticket.Destination,
		ticket.Date,
		strconv.Itoa(ticket.Price),
	}
	return
}
