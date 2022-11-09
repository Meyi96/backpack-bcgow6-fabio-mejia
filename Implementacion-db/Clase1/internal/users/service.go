package users

import (
	"errors"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase1/internal/domain"
)

type Service interface {
	GetAll() ([]domain.User, error)
	Get(id int) (domain.User, error)
	Store(name string, lastName string, email string, age int, height int) (domain.User, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) (domain.User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]domain.User, error) {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (domain.User, error) {
	return s.repository.Get(id)
}

func (s *service) Store(name string, lastName string, email string, age int, height int) (domain.User, error) {
	id, err := s.repository.Store(name, lastName, email, age, height)
	if err != nil {
		err = errors.New("no se pudo crear el usuario")
		return domain.User{}, err
	}
	user, err := s.repository.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, err
}

func (s *service) Update(id int, name string, lastName string, email string, age int, height int, active bool) (domain.User, error) {
	err := s.repository.Update(id, name, lastName, email, age, height, active)
	if err != nil {
		return domain.User{}, err
	}
	user, err := s.repository.Get(id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
