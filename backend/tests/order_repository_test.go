package tests

import (
	"testing"
	"time"

	"mangal-chai-backend/models"
	"mangal-chai-backend/repositories"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestOrderRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("CreateOrder", func(mt *mtest.T) {
		orderRepository := &repositories.OrderRepository{Collection: mt.Coll}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		order := models.Order{
			ID: "test_order_id",
			CustomerInfo: models.CustomerInfo{Name: "John Doe"},
			Items:        []models.CartItem{{ProductID: "prod1", Quantity: 1}},
			TotalAmount:  10.0,
			Status:       "pending",
			OrderDate:    time.Now(),
		}
		err := orderRepository.CreateOrder(order)
		assert.Nil(t, err)
	})

	mt.Run("GetOrder", func(mt *mtest.T) {
		orderRepository := &repositories.OrderRepository{Collection: mt.Coll}

		// Define nested BSON documents/arrays separately
		customerInfoDoc := bson.D{{Key: "name", Value: "John Doe"}}
		itemDoc := bson.D{{Key: "product_id", Value: "prod1"}, {Key: "quantity", Value: 1}}
		itemsArray := bson.A{itemDoc}

		expectedOrder := bson.D{
			{Key: "id", Value: "test_order_id"},
			{Key: "customer_info", Value: customerInfoDoc},
			{Key: "items", Value: itemsArray},
			{Key: "total_amount", Value: 10.0},
			{Key: "status", Value: "pending"},
			{Key: "order_date", Value: time.Now()},
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, expectedOrder))

		order, err := orderRepository.GetOrder("test_order_id")
		assert.Nil(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, "test_order_id", order.ID)
	})
}