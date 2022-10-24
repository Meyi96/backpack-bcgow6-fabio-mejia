package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.2/cmd/server/handler"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.2/internal/domain"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.2/internal/users"
	"github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase3.2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type responseUser struct {
	Code  string      `json:"code"`
	Data  domain.User `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

var (
	initialData []domain.User = []domain.User{
		{
			Id:           18,
			Name:         "Andres",
			LastName:     "Mejia",
			Email:        "example@mail.com",
			Age:          23,
			Height:       191,
			Active:       true,
			CreationDate: time.Date(2022, 3, 12, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:           19,
			Name:         "Lorena",
			LastName:     "Zamora",
			Email:        "examplo@mail.com",
			Age:          34,
			Height:       163,
			Active:       true,
			CreationDate: time.Date(2022, 5, 23, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:           20,
			Name:         "Camilo",
			LastName:     "Meneses",
			Email:        "dummy@mail.com",
			Age:          42,
			Height:       178,
			Active:       true,
			CreationDate: time.Date(2021, 7, 5, 0, 0, 0, 0, time.UTC),
		},
	}
)

func initialDataCopy() []domain.User {
	data := make([]domain.User, len(initialData))
	copy(data, initialData)
	return data
}
func createServer(mockStore *store.MockStore) *gin.Engine {
	// Print routes in console
	gin.SetMode(gin.ReleaseMode)
	//Modify token
	os.Setenv("TOKEN", "12345")
	repo := users.NewRepository(mockStore)
	service := users.NewService(repo)
	user := handler.NewUser(service)

	router := gin.Default()
	productGroup := router.Group("/users")
	productGroup.Use(handler.TokenAuthMiddleware)
	productGroup.PUT("/:id", handler.IdValidationMiddleware, user.Update)
	productGroup.DELETE("/:id", handler.IdValidationMiddleware, user.Delete)

	return router
}
func createRequestTest(method string, url string, body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", os.Getenv("TOKEN"))
	return req, httptest.NewRecorder()
}

//Testing

// Put user
func TestPutOk(t *testing.T) {
	expId := 19
	expName := "David"
	expLastName := "Parra"
	expEmail := "dummy@mail.com"
	expAge := 32
	expHeight := 151
	expActive := true
	body := fmt.Sprintf(`{"Id":%d,"Name":"%s","LastName":"%s","Email":"%s","Age":%d,"Height":%d,"Active":%t}`, expId, expName, expLastName, expEmail, expAge, expHeight, expActive)
	objRes := responseUser{}
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodPut, "/users/19", []byte(body))
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, expName, objRes.Data.Name)
	assert.Equal(t, expLastName, objRes.Data.LastName)
	assert.Equal(t, expEmail, objRes.Data.Email)
	assert.Equal(t, expAge, objRes.Data.Age)
	assert.Equal(t, expHeight, objRes.Data.Height)
	assert.Equal(t, expActive, objRes.Data.Active)
	assert.Equal(t, initialData[1].CreationDate, objRes.Data.CreationDate)
}
func TestPutIDNotFound(t *testing.T) {
	//Age is not sent
	body := fmt.Sprintf(`{"Name":"dummy","LastName":"dummy","Email":"dummy","Age":1,"Height":1,"Active":false}`)
	objRes := responseUser{}
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodPut, "/users/31", []byte(body))
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, mockStore.ReadWasCalled)
	assert.False(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.NotEmpty(t, objRes.Error)
	assert.Empty(t, objRes.Data)
}
func TestPutUnprocessableEntity(t *testing.T) {
	//Age is not sent
	body := fmt.Sprintf(`{"Name":"dummy","LastName":"dummy","Email":"dummy","Height":12,"Active":false}`)
	objRes := responseUser{}
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodPut, "/users/19", []byte(body))
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.False(t, mockStore.ReadWasCalled)
	assert.False(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.NotEmpty(t, objRes.Error)
	assert.Empty(t, objRes.Data)
}
func TestDeleteFirst(t *testing.T) {
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodDelete, "/users/18", []byte{})
	// Act.
	router.ServeHTTP(rr, req)
	// Assert.
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
	assert.Equal(t, initialData[1:], mockStore.DummyData)
}
func TestDeleteMiddle(t *testing.T) {
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodDelete, "/users/19", []byte{})
	// Act.
	router.ServeHTTP(rr, req)
	// Assert.
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
	assert.Equal(t, append(initialData[:1], initialData[2:]...), mockStore.DummyData)
}
func TestDeleteLast(t *testing.T) {
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodDelete, "/users/20", []byte{})
	// Act.
	router.ServeHTTP(rr, req)
	// Assert.
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusNoContent, rr.Code)
	assert.Empty(t, rr.Body.Bytes())
	assert.Equal(t, initialData[:2], mockStore.DummyData)
}
func TestDeleteIDNotFound(t *testing.T) {
	objRes := responseUser{}
	mockStore := store.MockStore{DummyData: initialDataCopy()}
	router := createServer(&mockStore)
	req, rr := createRequestTest(http.MethodDelete, "/users/21", []byte{})
	// Act.
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &objRes)
	// Assert.
	assert.Nil(t, err)
	assert.True(t, mockStore.ReadWasCalled)
	assert.False(t, mockStore.WriteWasCalled)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.NotEmpty(t, objRes.Error)
	assert.Empty(t, objRes.Data)
	assert.Equal(t, initialData, mockStore.DummyData)
}
