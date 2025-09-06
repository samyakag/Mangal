package repositories

import (
	"context"

	"mangal-chai-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryInterface interface {
	CreateOrder(order models.Order) error
	GetOrder(id string) (*models.Order, error)
}

type OrderRepository struct {
	Collection *mongo.Collection
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	_, err := r.Collection.InsertOne(context.TODO(), order)
	return err
}

func (r *OrderRepository) GetOrder(id string) (*models.Order, error) {
	var order models.Order
	err := r.Collection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}