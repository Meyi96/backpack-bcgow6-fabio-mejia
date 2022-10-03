package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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

func main() {
	jsonData, err := json.Marshal(getListUsers())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonData))

}

func getListUsers() []User {
	u1 := User{
		Name:         "Andres",
		LastName:     "Mejia",
		Email:        "example@mail.com",
		Age:          23,
		Height:       191,
		Active:       true,
		CreationDate: time.Now(),
	}
	u2 := User{
		Name:         "Carla",
		LastName:     "Lopez",
		Email:        "carla@mail.com",
		Age:          19,
		Height:       162,
		Active:       true,
		CreationDate: time.Now(),
	}
	u3 := User{
		Name:         "Pedro",
		LastName:     "Silva",
		Email:        "pedro@mail.com",
		Age:          27,
		Height:       182,
		Active:       true,
		CreationDate: time.Now(),
	}
	return append([]User{}, u1, u2, u3)
}
