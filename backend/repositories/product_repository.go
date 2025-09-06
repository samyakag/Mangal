package repositories

import (
	"context"

	"mangal-chai-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryInterface interface {
	GetProducts() ([]models.Product, error)
	GetProduct(id string) (*models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetCategories() ([]string, error)
	SeedProducts(products []interface{}) error
}

type ProductRepository struct {
	Collection *mongo.Collection
}

func (r *ProductRepository) GetProducts() ([]models.Product, error) {
	var products []models.Product
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProduct(id string) (*models.Product, error) {
	var product models.Product
	err := r.Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	cursor, err := r.Collection.Find(context.TODO(), bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetCategories() ([]string, error) {
	categories, err := r.Collection.Distinct(context.TODO(), "category", bson.M{})
	if err != nil {
		return nil, err
	}
	var categoryStrings []string
	for _, category := range categories {
		categoryStrings = append(categoryStrings, category.(string))
	}
	return categoryStrings, nil

}

func (r *ProductRepository) SeedProducts(products []interface{}) error {
	count, err := r.Collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	if count == 0 {
		_, err := r.Collection.InsertMany(context.TODO(), products)
		if err != nil {
			return err
		}
	}
	return nil
}