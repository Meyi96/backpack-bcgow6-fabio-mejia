package users

import (
	"fmt"
	"time"
)

var users []User

type User struct {
	Id           int
	Name         string
	LastName     string
	Email        string
	Age          int
	Height       int
	Active       bool
	CreationDate time.Time
}

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, name string, lastName string, email string, age int, height int, active bool, creationDate time.Time) (User, error)
	LastId() (int, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error)
	UpdateLastNameAndAge(id int, lastName string, age int) (User, error)
	Delete(id int) error
}

type repository struct {
	users []User
}

func NewRepository(users []User) Repository {
	return &repository{users: users}
}

func (r *repository) GetAll() ([]User, error) {
	return r.users, nil
}

func (r *repository) Store(id int, name string, lastName string, email string, age int, height int, active bool, creationDate time.Time) (User, error) {
	newUser := User{
		Id:           id,
		Name:         name,
		LastName:     lastName,
		Email:        email,
		Age:          age,
		Height:       height,
		Active:       active,
		CreationDate: creationDate,
	}
	r.users = append(r.users, newUser)
	return newUser, nil
}

func (r *repository) LastId() (int, error) {
	if len(r.users) < 0 {
		return -1, nil
	}
	return r.users[len(r.users)-1].Id, nil
}

func (r *repository) Update(id int, name string, lastName string, email string, age int, height int, active bool) (User, error) {
	user := User{
		Id:       id,
		Name:     name,
		LastName: lastName,
		Email:    email,
		Age:      age,
		Height:   height,
		Active:   active,
	}
	updated := false
	for i := range r.users {
		if r.users[i].Id == id {
			user.CreationDate = r.users[i].CreationDate
			r.users[i] = user
			updated = true
			break
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}

func (r *repository) UpdateLastNameAndAge(id int, lastName string, age int) (User, error) {
	updated := false
	var index int
	for i := range r.users {
		if r.users[i].Id == id {
			r.users[i].LastName = lastName
			r.users[i].Age = age
			updated = true
			index = i
			break
		}
	}
	if !updated {
		return User{}, fmt.Errorf("user with id %d not found", id)
	}
	return r.users[index], nil
}

func (r *repository) Delete(id int) error {
	found := false
	var index int
	for i := range r.users {
		if r.users[i].Id == id {
			index = i
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("user with id %d not found", id)
	}
	r.users = append(r.users[:index], r.users[index+1:]...)
	return nil
}
