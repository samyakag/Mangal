package tests

import (
	"testing"

	"mangal-chai-backend/models"
	"mangal-chai-backend/repositories"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestProductRepository(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("GetProducts", func(mt *mtest.T) {
		productRepository := &repositories.ProductRepository{Collection: mt.Coll}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: "1"}, {Key: "name", Value: "p1"}} )
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "id", Value: "2"}, {Key: "name", Value: "p2"}} )
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		products, err := productRepository.GetProducts()
		assert.Nil(t, err)
		assert.Len(t, products, 2)
	})

	mt.Run("GetProduct", func(mt *mtest.T) {
		productRepository := &repositories.ProductRepository{Collection: mt.Coll}
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "id", Value: "1"}, {Key: "name", Value: "p1"}} )
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, killCursors)

		product, err := productRepository.GetProduct("1")
		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, "p1", product.Name)
	})

	mt.Run("GetCategories", func(mt *mtest.T) {
		productRepository := &repositories.ProductRepository{Collection: mt.Coll}
		mt.AddMockResponses(mtest.CreateSuccessResponse(bson.E{Key: "values", Value: bson.A{"cat1", "cat2"}}))

		categories, err := productRepository.GetCategories()
		assert.Nil(t, err)
		assert.Len(t, categories, 2)
	})

	mt.Run("SeedProducts", func(mt *mtest.T) {
		productRepository := &repositories.ProductRepository{Collection: mt.Coll}
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{{Key: "n", Value: 0}}),
			mtest.CreateSuccessResponse(),
		)
		err := productRepository.SeedProducts([]interface{}{models.Product{ID: "1"}})
		assert.Nil(t, err)
	})
}