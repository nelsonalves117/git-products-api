package service

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nelsonalves117/go-products-api/internal/canonical"
	"github.com/nelsonalves117/go-products-api/internal/repositories"
	"github.com/sirupsen/logrus"
)

type Service interface {
	GetAllProducts() ([]canonical.Product, error)
	GetProductsByCategory(category string) ([]canonical.Product, error)
	GetProductById(id string) (canonical.Product, error)
	CreateProduct(product canonical.Product) (canonical.Product, error)
	UpdateProduct(id string, product canonical.Product) (canonical.Product, error)
	DeleteProduct(id string) error
}

type service struct {
	repo repositories.Repository
}

func New() Service {
	return &service{
		repo: repositories.New(),
	}
}

func (service *service) GetAllProducts() ([]canonical.Product, error) {
	product, err := service.repo.GetAllProducts()
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get all products")
		return []canonical.Product{}, err
	}

	return product, nil
}

func (service *service) GetProductsByCategory(category string) ([]canonical.Product, error) {
	product, err := service.repo.GetProductsByCategory(category)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get a product")
		return []canonical.Product{}, err
	}

	return product, nil
}

func (service *service) GetProductById(id string) (canonical.Product, error) {
	product, err := service.repo.GetProductById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get a product")
		return canonical.Product{}, err
	}

	return product, nil
}

func (service *service) CreateProduct(product canonical.Product) (canonical.Product, error) {
	product.Id = uuid.NewString()
	product.CreatedAt = time.Now()

	product, err := service.repo.CreateProduct(product)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to create a product")
		return canonical.Product{}, err
	}

	return product, nil
}

func (service *service) UpdateProduct(id string, product canonical.Product) (canonical.Product, error) {
	product, err := service.repo.UpdateProduct(id, product)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to update a product")
		return canonical.Product{}, err
	}

	return product, nil
}

func (service *service) DeleteProduct(id string) error {
	product, err := service.repo.GetProductById(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to get a product")
		return err
	}

	if product.Id == "" {
		return fmt.Errorf("product not found on db")
	}

	err = service.repo.DeleteProduct(id)
	if err != nil {
		logrus.WithError(err).Error("error occurred while trying to delete a product")
		return err
	}

	return nil
}
