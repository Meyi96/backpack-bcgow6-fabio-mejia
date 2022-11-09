package users

import (
	"context"
	"database/sql"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/internal/domain"
)

type Repository interface {
	GetAll(c context.Context) ([]domain.User, error)
	Get(c context.Context, id int) (domain.User, error)
	Store(c context.Context, user domain.User) (int, error)
	Update(c context.Context, id int, user domain.User) error
	Delete(c context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

var (
	GetAllQuery string = "select id, name, last_name, email, age, height, active, creation_date from user"
	GetQuery    string = "select id, name, last_name, email, age, height, active, creation_date from user where id = ?"
	StoreQuery  string = "insert into user (name, last_name, email, age, height) values (?, ?, ?, ?, ?)"
	UpdateQuery string = "update user set name = ?, last_name = ?, email = ?, age = ?, height = ?, active = ? where id = ?"
	DeleteQuery string = "delete from user where id = ?"
)

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Get(c context.Context, id int) (domain.User, error) {
	row := r.db.QueryRowContext(c, GetQuery, id)
	s := domain.User{}
	err := row.Scan(&s.Id, &s.Name, &s.LastName, &s.Email, &s.Age, &s.Height, &s.Active, &s.CreationDate)
	if err != nil {
		return domain.User{}, err
	}
	return s, nil
}
func (r *repository) GetAll(c context.Context) ([]domain.User, error) {
	rows, err := r.db.QueryContext(c, GetAllQuery)
	if err != nil {
		return nil, err
	}
	var users []domain.User
	for rows.Next() {
		s := domain.User{}
		err := rows.Scan(&s.Id, &s.Name, &s.LastName, &s.Email, &s.Age, &s.Height, &s.Active, &s.CreationDate)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, s)
	}
	return users, nil
}

func (r *repository) Store(c context.Context, user domain.User) (int, error) {
	stmt, err := r.db.PrepareContext(c, StoreQuery)
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(c, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) Update(c context.Context, id int, user domain.User) error {
	stmt, err := r.db.PrepareContext(c, UpdateQuery)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(c, &user.Name, &user.LastName, &user.Email, &user.Age, &user.Height, &user.Active, &id)
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

func (r *repository) Delete(c context.Context, id int) error {
	stmt, err := r.db.PrepareContext(c, DeleteQuery)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(c, id)
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
