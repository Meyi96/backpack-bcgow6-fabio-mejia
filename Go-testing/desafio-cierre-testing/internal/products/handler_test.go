package products

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockService *MockService) *gin.Engine {
	// Print routes in console
	gin.SetMode(gin.ReleaseMode)

	products := NewHandler(mockService)

	router := gin.Default()
	prodRoute := router.Group("/products")
	{
		prodRoute.GET("", products.GetProducts)
	}
	return router
}
func createRequestTest(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}

func TestGetProductsOk(t *testing.T) {
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
	mockService := MockService{DummyData: initialData}
	router := createServer(&mockService)
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=2", []byte{})
	objRes := []Product{}
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, mockService.GetAllWasCalled)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expProducts, objRes)
}
func TestGetProductsWithoutQueryParam(t *testing.T) {
	expErrorMsg := "seller_id query param is required"
	mockService := MockService{DummyData: initialData}
	router := createServer(&mockService)
	req, rr := createRequestTest(http.MethodGet, "/products", []byte{})
	objRes := map[string]string{}
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.False(t, mockService.GetAllWasCalled)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expErrorMsg, objRes["error"])
}
func TestGetProductsErrService(t *testing.T) {
	ErrService := errors.New("error con el servidor")
	mockService := MockService{DummyData: initialData, ErrConsult: ErrService}
	router := createServer(&mockService)
	req, rr := createRequestTest(http.MethodGet, "/products?seller_id=1", []byte{})
	objRes := map[string]string{}
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, mockService.GetAllWasCalled)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	assert.Equal(t, ErrService.Error(), objRes["error"])
}
