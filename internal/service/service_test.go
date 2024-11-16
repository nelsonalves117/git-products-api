package service

import (
	"errors"
	"testing"
	"time"

	"github.com/nelsonalves117/go-products-api/internal/canonical"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllProducts_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productsTest := []canonical.Product{
		{
			Id:        "xpto",
			Name:      "test",
			Category:  "testCategory",
			Price:     200,
			Stock:     10,
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAllProducts").Return(productsTest, nil)

	service := &service{
		repo: mockRepo,
	}

	products, err := service.GetAllProducts()

	assert.Nil(t, err)
	assert.Equal(t, "xpto", products[0].Id)
	assert.Equal(t, "test", products[0].Name)
	assert.Equal(t, "testCategory", products[0].Category)
	assert.Equal(t, float32(200), products[0].Price)
	assert.Equal(t, 10, products[0].Stock)
	assert.True(t, products[0].CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}
func TestGetAllProducts_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetAllProducts").Return([]canonical.Product{}, errors.New("error occurred when trying to get all products"))

	service := &service{
		repo: mockRepo,
	}

	products, err := service.GetAllProducts()

	assert.NotNil(t, err)
	assert.Empty(t, products)
	assert.Equal(t, "error occurred when trying to get all products", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetProductsByCategory_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productsTest := []canonical.Product{
		{
			Id:        "xpto",
			Name:      "test",
			Category:  "testCategory",
			Price:     200,
			Stock:     10,
			CreatedAt: time.Now(),
		},
	}

	mockRepo.On("GetProductsByCategory", "testCategory").Return(productsTest, nil)

	service := &service{
		repo: mockRepo,
	}

	products, err := service.GetProductsByCategory("testCategory")

	assert.Nil(t, err)
	assert.Equal(t, "xpto", products[0].Id)
	assert.Equal(t, "test", products[0].Name)
	assert.Equal(t, "testCategory", products[0].Category)
	assert.Equal(t, float32(200), products[0].Price)
	assert.Equal(t, 10, products[0].Stock)
	assert.True(t, products[0].CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestGetProductsByCategory_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetProductsByCategory", "testCategory").Return([]canonical.Product{}, errors.New("error occurred when trying to get a product"))

	service := &service{
		repo: mockRepo,
	}

	products, err := service.GetProductsByCategory("testCategory")

	assert.NotNil(t, err)
	assert.Empty(t, products)
	assert.Equal(t, "error occurred when trying to get a product", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestGetProductById_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{

		Id:        "xpto",
		Name:      "test",
		Category:  "testCategory",
		Price:     200,
		Stock:     10,
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetProductById", "xpto").Return(productTest, nil)

	service := &service{
		repo: mockRepo,
	}

	product, err := service.GetProductById("xpto")

	assert.Nil(t, err)
	assert.Equal(t, "xpto", product.Id)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, "testCategory", product.Category)
	assert.Equal(t, float32(200), product.Price)
	assert.Equal(t, 10, product.Stock)
	assert.True(t, product.CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestGetProductById_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	mockRepo.On("GetProductById", "xpto").Return(canonical.Product{}, errors.New("error occurred when trying to get a product"))

	service := &service{
		repo: mockRepo,
	}

	product, err := service.GetProductById("xpto")

	assert.NotNil(t, err)
	assert.Empty(t, product)
	assert.Equal(t, "error occurred when trying to get a product", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestCreateProduct_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Name:     "test",
		Category: "testCategory",
		Price:    200,
		Stock:    10,
	}

	updatedProduct := canonical.Product{
		Id:       "xpto",
		Name:     "test",
		Category: "testCategory",
		Price:    200,
		Stock:    10,
	}

	mockRepo.On("CreateProduct", mock.MatchedBy(func(product canonical.Product) bool {
		return product.Name == "test" && product.Category == "testCategory" && product.Price == 200 && product.Stock == 10
	})).Return(updatedProduct, nil)

	service := &service{
		repo: mockRepo,
	}

	product, err := service.CreateProduct(productTest)

	assert.Nil(t, err)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, "testCategory", product.Category)
	assert.Equal(t, float32(200), product.Price)
	assert.Equal(t, 10, product.Stock)

	mockRepo.AssertExpectations(t)
}

func TestCreateProduct_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Name:     "test",
		Category: "testCategory",
		Price:    200,
		Stock:    10,
	}

	mockRepo.On("CreateProduct", mock.MatchedBy(func(product canonical.Product) bool {
		return product.Name == "test" && product.Category == "testCategory" && product.Price == 200 && product.Stock == 10
	})).Return(canonical.Product{}, errors.New("error occurred when trying to create a product"))

	service := &service{
		repo: mockRepo,
	}

	product, err := service.CreateProduct(productTest)

	assert.NotNil(t, err)
	assert.Empty(t, product)
	assert.Equal(t, "error occurred when trying to create a product", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Name:     "test",
		Category: "testCategory",
		Price:    200,
		Stock:    10,
	}

	updatedProduct := canonical.Product{
		Id:        "xpto",
		Name:      "test",
		Category:  "testCategory",
		Price:     200,
		Stock:     10,
		CreatedAt: time.Now(),
	}

	mockRepo.On("UpdateProduct", "xpto", mock.MatchedBy(func(product canonical.Product) bool {
		return product.Name == "test" && product.Category == "testCategory" && product.Price == 200 && product.Stock == 10
	})).Return(updatedProduct, nil)

	service := &service{
		repo: mockRepo,
	}

	product, err := service.UpdateProduct("xpto", productTest)

	assert.Nil(t, err)
	assert.Equal(t, "xpto", product.Id)
	assert.Equal(t, "test", product.Name)
	assert.Equal(t, "testCategory", product.Category)
	assert.Equal(t, float32(200), product.Price)
	assert.Equal(t, 10, product.Stock)
	assert.True(t, product.CreatedAt.After(time.Date(2021, 1, 1, 0, 0, 0, 0, time.Local)))

	mockRepo.AssertExpectations(t)
}

func TestUpdateProduct_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Name:     "test",
		Category: "testCategory",
		Price:    200,
		Stock:    10,
	}

	mockRepo.On("UpdateProduct", "xpto", mock.MatchedBy(func(product canonical.Product) bool {
		return product.Name == "test" && product.Category == "testCategory" && product.Price == 200 && product.Stock == 10
	})).Return(canonical.Product{}, errors.New("error occurred when trying to update a product"))

	service := &service{
		repo: mockRepo,
	}

	product, err := service.UpdateProduct("xpto", productTest)

	assert.NotNil(t, err)
	assert.Empty(t, product)
	assert.Equal(t, "error occurred when trying to update a product", err.Error())

	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_Success(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Id:        "xpto",
		Name:      "test",
		Category:  "testCategory",
		Price:     200,
		Stock:     10,
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetProductById", "xpto").Return(productTest, nil)

	mockRepo.On("DeleteProduct", "xpto").Return(nil)

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteProduct("xpto")

	assert.Nil(t, err)

	mockRepo.AssertExpectations(t)
}

func TestDeleteProduct_Error(t *testing.T) {
	mockRepo := new(MockRepository)

	productTest := canonical.Product{
		Id:        "xpto",
		Name:      "test",
		Category:  "testCategory",
		Price:     200,
		Stock:     10,
		CreatedAt: time.Now(),
	}

	mockRepo.On("GetProductById", "xpto").Return(productTest, nil)

	mockRepo.On("DeleteProduct", "xpto").Return(errors.New("error occurred when trying to delete a product"))

	service := &service{
		repo: mockRepo,
	}

	err := service.DeleteProduct("xpto")

	assert.NotNil(t, err)
	assert.Equal(t, "error occurred when trying to delete a product", err.Error())

	mockRepo.AssertExpectations(t)
}
