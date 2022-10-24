package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var initialData []Product = []Product{
	{
		ID:          "5",
		SellerID:    "2",
		Description: "Iphone 8",
		Price:       100.0,
	},
	{
		ID:          "6",
		SellerID:    "2",
		Description: "Iphone X",
		Price:       250.0,
	},
	{
		ID:          "7",
		SellerID:    "2",
		Description: "Iphone 13 Pro Max",
		Price:       1200.0,
	},
	{
		ID:          "8",
		SellerID:    "13",
		Description: "TV samsung 45 pulgadas",
		Price:       500.0,
	},
}

func TestGetAllBySeller(t *testing.T) {
	sellerId := "2"
	expProducts := []Product{
		{
			ID:          "5",
			SellerID:    "2",
			Description: "Iphone 8",
			Price:       100.0,
		},
		{
			ID:          "6",
			SellerID:    "2",
			Description: "Iphone X",
			Price:       250.0,
		},
		{
			ID:          "7",
			SellerID:    "2",
			Description: "Iphone 13 Pro Max",
			Price:       1200.0,
		},
	}
	repo := MockRepository{DummyData: initialData}
	service := NewService(&repo)

	result, err := service.GetAllBySeller(sellerId)

	assert.Nil(t, err)
	assert.True(t, repo.GetAllWasCalled)
	assert.Equal(t, expProducts, result)
}
func TestGetAllBySellerIDNotExist(t *testing.T) {
	sellerId := "3"
	expProducts := []Product{}
	repo := MockRepository{DummyData: initialData}
	service := NewService(&repo)

	result, err := service.GetAllBySeller(sellerId)

	assert.Nil(t, err)
	assert.True(t, repo.GetAllWasCalled)
	assert.Equal(t, expProducts, result)
}
func TestGetAllBySellerErrWithRepo(t *testing.T) {
	sellerId := "13"
	ErrRepo := errors.New("error en el repositorio")
	repo := MockRepository{DummyData: initialData, ErrConsult: ErrRepo}
	service := NewService(&repo)

	result, err := service.GetAllBySeller(sellerId)

	assert.ErrorContains(t, err, ErrRepo.Error())
	assert.True(t, repo.GetAllWasCalled)
	assert.Nil(t, result)
}
