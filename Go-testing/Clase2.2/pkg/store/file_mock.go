package store

import "github.com/Meyi96/backpack-bcgow6-fabio-mejia/Go-testing/Clase2.2/internal/domain"

type MockStore struct {
	ReadWasCalled  bool
	WriteWasCalled bool
	ErrRead        error
	ErrWrite       error
	DummyData      []domain.User
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true
	if m.ErrRead != nil {
		return m.ErrRead
	}
	*data.(*[]domain.User) = m.DummyData
	return nil
}
func (m *MockStore) Write(data interface{}) error {
	m.WriteWasCalled = true
	if m.ErrWrite != nil {
		return m.ErrWrite
	}
	m.DummyData = data.([]domain.User)
	return nil
}
