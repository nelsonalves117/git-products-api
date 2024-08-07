package rest

import (
	"time"

	"github.com/nelsonalves117/go-products-api/internal/canonical"
)

func toCanonical(product productRequest) canonical.Product {
	return canonical.Product{
		Name:     product.Name,
		Category: product.Category,
		Price:    product.Price,
		Stock:    product.Stock,
	}
}

func toResponse(product canonical.Product) productResponse {
	return productResponse{
		Id:        product.Id,
		Name:      product.Name,
		Category:  product.Category,
		Price:     product.Price,
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt.Format(time.RFC3339),
	}
}
