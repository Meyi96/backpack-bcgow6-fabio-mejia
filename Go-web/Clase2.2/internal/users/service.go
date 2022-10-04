package users

import (
	"errors"
	"time"
)

type Service interface {
	GetAll() ([]User, error)
	Store(name string, lastName string, email string, age int, height int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (s *service) Store(name string, lastName string, email string, age int, height int) (user User, err error) {
	id, err := s.repository.LastId()
	if err != nil {
		err = errors.New("no se pudo generar un Id")
		return
	}
	id++
	creationDate := time.Now()
	user, err = s.repository.Store(id, name, lastName, email, age, height, true, creationDate)
	if err != nil {
		err = errors.New("no se pudo crear el usuario")
		return
	}
	return user, err
}
