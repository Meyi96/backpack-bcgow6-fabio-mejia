package domain

import "time"

type User struct {
	Id           int       `db:"id"`
	Name         string    `db:"name"`
	LastName     string    `db:"last_name"`
	Email        string    `db:"email"`
	Age          int       `db:"age"`
	Height       int       `db:"height"`
	Active       bool      `db:"active"`
	CreationDate time.Time `db:"creation_date"`
}
