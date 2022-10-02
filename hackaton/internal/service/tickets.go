package service

import (
	"errors"
	"fmt"
)

type Bookings interface {
	// Create create a new Ticket
	Create(t Ticket) (Ticket, error)
	// Read read a Ticket by id
	Read(id int) (Ticket, error)
	// Update update values of a Ticket
	Update(id int, t Ticket) (Ticket, error)
	// Delete delete a Ticket by id
	Delete(id int) (int, error)

	ReadAll() []Ticket
}

type GeneratorID struct {
	Current int
	Limit   int
}

type bookings struct {
	Tickets []Ticket
	GenerId GeneratorID
}

type Ticket struct {
	Id          int    `csv:"id"`
	Name        string `csv:"nombre"`
	Email       string `csv:"email"`
	Destination string `csv:"destino"`
	Date        string `csv:"fecha"`
	Price       int    `csv:"precio"`
}

// NewBookings creates a new bookings service
func NewBookings(Tickets []Ticket) Bookings {
	generId := GeneratorID{Current: 1000, Limit: 2000}
	return &bookings{Tickets: Tickets, GenerId: generId}
}

func (b *bookings) Create(t Ticket) (Ticket, error) {
	if err := validateTicket(t); err != nil {
		return Ticket{}, fmt.Errorf("no se puerde crear el ticket: %w\n", err)
	}
	id, err := b.GenerId.generatorIds()
	if err != nil {
		return Ticket{}, err
	}
	t.Id = id
	b.Tickets = append(b.Tickets, t)
	return t, nil
}

func (b *bookings) Read(id int) (Ticket, error) {
	index, err := b.findById(id)
	if err != nil {
		return Ticket{}, err
	}
	return b.Tickets[index], nil
}

func (b *bookings) Update(id int, t Ticket) (Ticket, error) {
	if err := validateTicket(t); err != nil {
		return Ticket{}, fmt.Errorf("no se puerde crear el ticket: %w\n", err)
	}
	index, err := b.findById(id)
	if err != nil {
		return Ticket{}, err
	}
	b.Tickets[index] = t
	return t, nil
}

func (b *bookings) Delete(id int) (int, error) {
	index, err := b.findById(id)
	if err != nil {
		return 0, err
	}
	if index+1 == len(b.Tickets) {
		b.Tickets = b.Tickets[:index]
	} else {
		b.Tickets = append(b.Tickets[:index], b.Tickets[index+1:]...)
	}
	return id, nil
}

func (b *bookings) ReadAll() []Ticket {
	return b.Tickets
}

func (b bookings) findById(id int) (index int, err error) {
	for i, ticket := range b.Tickets {
		if ticket.Id == id {
			index = i
			break
		}
		if i+1 == len(b.Tickets) {
			err = fmt.Errorf("No existe ticket con el id: %d\n", id)
		}
	}
	return
}

func validateTicket(ticket Ticket) (err error) {
	if ticket.Date == "" {
		return errors.New("el campo data es requerido")
	}
	if ticket.Destination == "" {
		return errors.New("el campo destino es requerido")
	}
	if ticket.Email == "" {
		return errors.New("el campo email es requerido")
	}
	if ticket.Name == "" {
		return errors.New("el campo nombres es requerido")
	}
	if ticket.Price == 0 {
		return errors.New("el campo precio es requerido")
	}
	return
}

func (gId *GeneratorID) generatorIds() (id int, err error) {
	if gId.Current+1 > gId.Limit {
		return 0, errors.New("Limite de Ids alcanzado")
	}
	id = gId.Current + 1
	return
}
