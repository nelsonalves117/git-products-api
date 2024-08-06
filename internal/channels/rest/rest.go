package rest

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nelsonalves117/go-products-api/internal/config"
	"github.com/nelsonalves117/go-products-api/internal/service"
)

type Rest interface {
	Start() error
}

type rest struct {
	service service.Service
}

func New() Rest {
	return &rest{
		service: service.New(),
	}
}

func (rest *rest) Start() error {
	router := echo.New()

	router.Use(middleware.Logger())

	router.GET("/products", rest.GetAllProducts)
	router.GET("/products/:id", rest.GetProductById)
	router.GET("/products/categories/:category", rest.GetProductsByCategory)
	router.POST("/products/create", rest.CreateProduct)
	router.PUT("/products/update/:id", rest.UpdateProduct)
	router.DELETE("/products/delete/:id", rest.DeleteProduct)

	return router.Start(":" + config.Get().Port)
}

func (rest *rest) GetAllProducts(c echo.Context) error {
	productSlice, err := rest.service.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, productSlice)
}

func (rest *rest) GetProductsByCategory(c echo.Context) error {
	category := c.Param("category")

	productSlice, err := rest.service.GetProductsByCategory(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, productSlice)
}

func (rest *rest) GetProductById(c echo.Context) error {
	id := c.Param("id")

	product, err := rest.service.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, product)
}

func (rest *rest) CreateProduct(c echo.Context) error {
	var product productRequest

	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	createdProduct, err := rest.service.CreateProduct(toCanonical(product))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusCreated, toResponse(createdProduct))
}

func (rest *rest) UpdateProduct(c echo.Context) error {
	var product productRequest

	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid data"))
	}

	id := c.Param("id")
	updatedProduct, err := rest.service.UpdateProduct(id, toCanonical(product))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, toResponse(updatedProduct))
}

func (rest *rest) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := rest.service.DeleteProduct(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.New("unexpected error occurred"))
	}

	return c.JSON(http.StatusOK, nil)
}
