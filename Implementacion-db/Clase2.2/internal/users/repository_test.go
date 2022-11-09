package users

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/internal/domain"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Implementacion-db/Clase2.2/pkg/util"
	"github.com/stretchr/testify/assert"
)

var (
	usersDummy = []domain.User{
		{
			Id:           1,
			Name:         "Fabio",
			LastName:     "Mejia",
			Email:        "fabio@mail.com",
			Age:          23,
			Height:       191,
			Active:       true,
			CreationDate: time.Now(),
		},
		{
			Id:           2,
			Name:         "Claudia",
			LastName:     "Aguirre",
			Email:        "agui@mail.com",
			Age:          42,
			Height:       154,
			Active:       true,
			CreationDate: time.Now(),
		},
	}
	errorDB = errors.New("some error in data base")
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("GetAll ok", func(t *testing.T) {
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
	})
	t.Run("GetAll error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(GetAllQuery)).WillReturnError(errorDB)

		repo := NewRepository(db)
		result, err := repo.GetAll(context.TODO())

		assert.ErrorIs(t, err, errorDB)
		assert.Nil(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Get ok", func(t *testing.T) {
		id := 1
		rows := sqlmock.NewRows([]string{"id", "name", "last_name", "email", "age", "height", "active", "creation_date"})
		rows.AddRow(usersDummy[0].Id, usersDummy[0].Name, usersDummy[0].LastName, usersDummy[0].Email, usersDummy[0].Age, usersDummy[0].Height, usersDummy[0].Active, usersDummy[0].CreationDate)
		mock.ExpectQuery(regexp.QuoteMeta(GetQuery)).WithArgs(id).WillReturnRows(rows)

		repo := NewRepository(db)
		result, err := repo.Get(context.TODO(), id)

		assert.NoError(t, err)
		assert.Equal(t, usersDummy[0], result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Get error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(GetQuery)).WithArgs(0).WillReturnError(errorDB)

		repo := NewRepository(db)
		result, err := repo.Get(context.TODO(), 0)

		assert.ErrorIs(t, err, errorDB)
		assert.Empty(t, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestStore(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	t.Run("Store ok", func(t *testing.T) {
		lastId := 5
		mock.ExpectPrepare(regexp.QuoteMeta(StoreQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(int64(lastId), 0))

		repo := NewRepository(db)
		id, err := repo.Store(context.TODO(), usersDummy[0])

		assert.NoError(t, err)
		assert.Equal(t, lastId, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Store error", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(StoreQuery)).ExpectExec().WillReturnError(errorDB)

		repo := NewRepository(db)
		id, err := repo.Store(context.TODO(), usersDummy[0])

		assert.ErrorIs(t, err, errorDB)
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	t.Run("Update ok", func(t *testing.T) {
		affectedRows := 1
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, int64(affectedRows)))

		repo := NewRepository(db)
		err := repo.Update(context.TODO(), 1, usersDummy[1])

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Update error", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery)).ExpectExec().WillReturnError(errorDB)

		repo := NewRepository(db)
		err := repo.Update(context.TODO(), 0, usersDummy[0])

		assert.ErrorIs(t, err, errorDB)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Update not found", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(UpdateQuery)).ExpectExec().WillReturnResult(sqlmock.NewResult(0, 0))

		repo := NewRepository(db)
		err := repo.Update(context.TODO(), 0, usersDummy[1])

		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Delete Ok", func(t *testing.T) {
		deleteID := 5
		mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery)).ExpectExec().WithArgs(deleteID).WillReturnResult(sqlmock.NewResult(0, 1))
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		repo := NewRepository(db)
		err = repo.Delete(c, deleteID)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Delete error", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(DeleteQuery)).ExpectExec().WithArgs(0).WillReturnError(errorDB)
		c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		repo := NewRepository(db)
		err = repo.Delete(c, 0)

		assert.ErrorIs(t, err, errorDB)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

// Testing with TXDB

func TestGetAllTxbd(t *testing.T) {
	t.Run("Get all ok", func(t *testing.T) {
		db := util.InitTxDB()

		repo := NewRepository(db)
		result, err := repo.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestGetTxbd(t *testing.T) {
	t.Run("Get by id ok", func(t *testing.T) {
		id := 2
		expectName := "lola"
		expectLastName := "otro"
		db := util.InitTxDB()

		repo := NewRepository(db)
		result, err := repo.Get(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
		assert.Equal(t, id, result.Id)
		assert.Equal(t, expectName, result.Name)
		assert.Equal(t, expectLastName, result.LastName)
	})
	t.Run("Get by id dosen't exist", func(t *testing.T) {
		id := 1
		db := util.InitTxDB()

		repo := NewRepository(db)
		result, err := repo.Get(context.TODO(), id)

		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.Empty(t, result)
	})
}

func TestStoreTxbd(t *testing.T) {
	t.Run("Store ok", func(t *testing.T) {
		db := util.InitTxDB()
		expectUser := usersDummy[1]
		repo := NewRepository(db)
		id, err := repo.Store(context.TODO(), expectUser)

		assert.NoError(t, err)

		user, err := repo.Get(context.TODO(), id)
		assert.NoError(t, err)
		assert.Equal(t, expectUser.Name, user.Name)
		assert.Equal(t, expectUser.LastName, user.LastName)
	})
}

func TestUpdateTxbd(t *testing.T) {
	t.Run("Update ok", func(t *testing.T) {
		db := util.InitTxDB()
		expectUser := usersDummy[0]
		updateId := 6
		repo := NewRepository(db)
		err := repo.Update(context.TODO(), updateId, expectUser)

		assert.NoError(t, err)

		user, err := repo.Get(context.TODO(), updateId)
		assert.NoError(t, err)
		assert.Equal(t, expectUser.Name, user.Name)
		assert.Equal(t, expectUser.LastName, user.LastName)
	})
	t.Run("Update id dosen't exist", func(t *testing.T) {
		db := util.InitTxDB()
		updateId := 1
		repo := NewRepository(db)
		err := repo.Update(context.TODO(), updateId, domain.User{})

		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}
func TestDeleteTxbd(t *testing.T) {
	t.Run("Update ok", func(t *testing.T) {
		db := util.InitTxDB()
		deleteId := 6
		repo := NewRepository(db)
		err := repo.Delete(context.TODO(), deleteId)

		assert.NoError(t, err)

		user, err := repo.Get(context.TODO(), deleteId)
		assert.ErrorIs(t, err, sql.ErrNoRows)
		assert.Empty(t, user)
	})
	t.Run("Update id dosen't exist", func(t *testing.T) {
		db := util.InitTxDB()
		deleteId := 1
		repo := NewRepository(db)
		err := repo.Delete(context.TODO(), deleteId)

		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}
