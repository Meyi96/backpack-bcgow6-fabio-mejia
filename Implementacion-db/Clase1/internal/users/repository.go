package users

import (
	"database/sql"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase1/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.User, error)
	Get(id int) (domain.User, error)
	Store(name string, lastName string, email string, age int, height int) (int, error)
	Update(id int, name string, lastName string, email string, age int, height int, active bool) error
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

var (
	GetQuery    string = "select id, name, last_name, email, age, height, active, creation_date from user where id = ?"
	StoreQuery  string = "insert into user (name, last_name, email, age, height) values (?, ?, ?, ?, ?)"
	UpdateQuery string = "update user set name = ?, last_name = ?, email = ?, age = ?, height = ?, active = ? where id = ?"
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Get(id int) (domain.User, error) {
	row := r.db.QueryRow(GetQuery, id)
	s := domain.User{}
	err := row.Scan(&s.Id, &s.Name, &s.LastName, &s.Email, &s.Age, &s.Height, &s.Active, &s.CreationDate)
	if err != nil {
		return domain.User{}, err
	}
	return s, nil
}
func (r *repository) GetAll() ([]domain.User, error) {

	return nil, nil
}

func (r *repository) Store(name string, lastName string, email string, age int, height int) (int, error) {
	stmt, err := r.db.Prepare(StoreQuery)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(&name, &lastName, &email, &age, &height)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) Update(id int, name string, lastName string, email string, age int, height int, active bool) error {
	stmt, err := r.db.Prepare(UpdateQuery)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(&name, &lastName, &email, &age, &height, &active, &id)
	if err != nil {
		return err
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *repository) Delete(id int) error {
	return nil
}
