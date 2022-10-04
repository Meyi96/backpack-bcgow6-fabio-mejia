package users

import "time"

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
