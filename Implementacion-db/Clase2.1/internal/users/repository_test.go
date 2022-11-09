package users

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.1/internal/domain"
	"github.com/stretchr/testify/assert"
)

var usersDummy = []domain.User{
	{
		Name:         "Fabio",
		LastName:     "Mejia",
		Email:        "fabio@mail.com",
		Age:          23,
		Height:       191,
		Active:       true,
		CreationDate: time.Now(),
	},
	{
		Name:         "Claudia",
		LastName:     "Aguirre",
		Email:        "agui@mail.com",
		Age:          42,
		Height:       154,
		Active:       true,
		CreationDate: time.Now(),
	},
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "last_name", "email", "age", "height", "active", "creation_date"})
	for _, u := range usersDummy {
		rows.AddRow(u.Id, u.Name, u.LastName, u.Email, u.Age, u.Height, u.Active, u.CreationDate)
	}
	mock.ExpectQuery(regexp.QuoteMeta(GetAllQuery)).WillReturnRows(rows)

	repo := NewRepository(db)
	result, err := repo.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, usersDummy, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	lastId := 5
	mock.ExpectPrepare(regexp.QuoteMeta(StoreQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(int64(lastId), 0))

	repo := NewRepository(db)
	id, err := repo.Store(context.TODO(), usersDummy[0])

	assert.NoError(t, err)
	assert.Equal(t, lastId, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	deleteID := 5
	mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery)).ExpectExec().WithArgs(deleteID).WillReturnResult(sqlmock.NewResult(0, 1))
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := NewRepository(db)
	err = repo.Delete(c, deleteID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
