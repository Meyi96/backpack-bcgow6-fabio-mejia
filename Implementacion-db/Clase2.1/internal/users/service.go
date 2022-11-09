package users

import (
	"context"
	"errors"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.1/internal/domain"
)

type Service interface {
	GetAll(c context.Context) ([]domain.User, error)
	Get(c context.Context, id int) (domain.User, error)
	Store(c context.Context, user domain.User) (domain.User, error)
	Update(c context.Context, id int, user domain.User) (domain.User, error)
	Delete(c context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(c context.Context) ([]domain.User, error) {
	return s.repository.GetAll(c)
}

func (s *service) Get(c context.Context, id int) (domain.User, error) {
	return s.repository.Get(c, id)
}

func (s *service) Store(c context.Context, user domain.User) (domain.User, error) {
	id, err := s.repository.Store(c, user)
	if err != nil {
		err = errors.New("no se pudo crear el usuario")
		return domain.User{}, err
	}
	user.Id = id
	return user, nil
}

func (s *service) Update(c context.Context, id int, user domain.User) (domain.User, error) {
	err := s.repository.Update(c, id, user)
	if err != nil {
		return domain.User{}, err
	}
	user, err = s.repository.Get(c, id)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (s *service) Delete(c context.Context, id int) error {
	return s.repository.Delete(c, id)
}
