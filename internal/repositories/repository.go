package repositories

import (
	"context"

	"github.com/nelsonalves117/go-products-api/internal/canonical"
	"github.com/nelsonalves117/go-products-api/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetAllProducts() ([]canonical.Product, error)
	GetProductsByCategory(category string) ([]canonical.Product, error)
	GetProductById(id string) (canonical.Product, error)
	CreateProduct(product canonical.Product) (canonical.Product, error)
	UpdateProduct(id string, product canonical.Product) (canonical.Product, error)
	DeleteProduct(id string) error
}

type repository struct {
	collection *mongo.Collection
}

func New() Repository {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Get().ConnectionString))
	if err != nil {
		panic(err)
	}

	return &repository{
		collection: client.Database("product_db").Collection("productSlice"),
	}
}

func (repo *repository) GetAllProducts() ([]canonical.Product, error) {
	var productSlice []canonical.Product

	res, err := repo.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	for res.Next(context.Background()) {
		var product canonical.Product

		err := res.Decode(&product)
		if err != nil {
			return nil, err
		}

		productSlice = append(productSlice, product)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return productSlice, nil
}

func (repo *repository) GetProductsByCategory(category string) ([]canonical.Product, error) {
	var productSlice []canonical.Product

	filter := bson.D{{Key: "category", Value: category}}

	res, err := repo.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	for res.Next(context.Background()) {
		var product canonical.Product

		err := res.Decode(&product)
		if err != nil {
			return nil, err
		}

		productSlice = append(productSlice, product)
	}

	if err := res.Err(); err != nil {
		return nil, err
	}

	return productSlice, nil
}

func (repo *repository) GetProductById(id string) (canonical.Product, error) {
	var product canonical.Product

	err := repo.collection.FindOne(context.Background(), bson.D{
		{
			Key:   "_id",
			Value: id,
		},
	}).Decode(&product)

	if err != nil {
		return canonical.Product{}, err
	}

	return product, nil
}

func (repo *repository) CreateProduct(product canonical.Product) (canonical.Product, error) {
	_, err := repo.collection.InsertOne(context.Background(), product)
	if err != nil {
		return canonical.Product{}, err
	}

	return product, nil
}

func (repo *repository) UpdateProduct(id string, product canonical.Product) (canonical.Product, error) {
	filter := bson.D{{Key: "_id", Value: id}}
	fields := bson.M{
		"$set": bson.M{
			"name":     product.Name,
			"category": product.Category,
			"price":    product.Price,
		},
	}

	_, err := repo.collection.UpdateOne(context.Background(), filter, fields)

	if err != nil {
		return canonical.Product{}, err
	}

	return product, nil
}

func (repo *repository) DeleteProduct(id string) error {
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
