package service

import (
	"github.com/nelsonalves117/go-products-api/internal/canonical"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAllProducts() ([]canonical.Product, error) {
	args := m.Called()
	return args.Get(0).([]canonical.Product), args.Error(1)
}

func (m *MockRepository) GetProductsByCategory(category string) ([]canonical.Product, error) {
	args := m.Called(category)
	return args.Get(0).([]canonical.Product), args.Error(1)
}

func (m *MockRepository) GetProductById(id string) (canonical.Product, error) {
	args := m.Called(id)
	return args.Get(0).(canonical.Product), args.Error(1)
}

func (m *MockRepository) CreateProduct(product canonical.Product) (canonical.Product, error) {
	args := m.Called(product)
	return args.Get(0).(canonical.Product), args.Error(1)
}

func (m *MockRepository) UpdateProduct(id string, product canonical.Product) (canonical.Product, error) {
	args := m.Called(id, product)
	return args.Get(0).(canonical.Product), args.Error(1)
}

func (m *MockRepository) DeleteProduct(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
