package tickets

import (
	"desafio-go-web-fabio-mejia/internal/domain"
)

type Service interface {
	GetTicketsByCountry(destination string) ([]domain.Ticket, error)
	GetTotalTicketsByCountry(destination string) (int, error)
	AverageDestination(destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetTicketsByCountry(destination string) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	return tickets, nil
}
func (s *service) GetTotalTicketsByCountry(destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}
func (s *service) AverageDestination(destination string) (int, error) {
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return 0, err
	}
	return int(len(tickets) / 24), nil
}
