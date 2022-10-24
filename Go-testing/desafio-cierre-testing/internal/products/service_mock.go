package products

type MockService struct {
	GetAllWasCalled bool
	ErrConsult      error
	DummyData       []Product
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {
	m.GetAllWasCalled = true
	if m.ErrConsult != nil {
		return []Product{}, m.ErrConsult
	}
	result := []Product{}
	for _, product := range m.DummyData {
		if product.SellerID == sellerID {
			result = append(result, product)
		}
	}
	return result, nil
}
