
package services

import (
	"mangal-chai-backend/models"
	"mangal-chai-backend/repositories"
)

type ProductServiceInterface interface {
	GetProducts() ([]models.Product, error)
	GetProduct(id string) (*models.Product, error)
	GetProductsByCategory(category string) ([]models.Product, error)
	GetCategories() ([]string, error)
	SeedProducts() error
}

type ProductService struct {
	Repository repositories.ProductRepositoryInterface
}

func (s *ProductService) GetProducts() ([]models.Product, error) {
	return s.Repository.GetProducts()
}

func (s *ProductService) GetProduct(id string) (*models.Product, error) {
	return s.Repository.GetProduct(id)
}

func (s *ProductService) GetProductsByCategory(category string) ([]models.Product, error) {
	return s.Repository.GetProductsByCategory(category)
}

func (s *ProductService) GetCategories() ([]string, error) {
	return s.Repository.GetCategories()
}

func (s *ProductService) SeedProducts() error {
	sampleProducts := []interface{}{
		models.Product{ID: "a1b2c3d4-e5f6-7890-1234-567890abcdef", Name: "Premium Assam Black Tea", Description: "Rich, malty Assam tea with robust flavor. Perfect for morning tea with milk and sugar. Sourced from the finest tea gardens of Assam.", Price: 299.0, Category: "Black Tea", ImageURL: "https://images.unsplash.com/photo-1563822249366-3efb23b8e0c9", InStock: true, Weight: "100g"},
		models.Product{ID: "b2c3d4e5-f6a7-8901-2345-67890abcdef0", Name: "Darjeeling Muscatel", Description: "Delicate and aromatic Darjeeling tea with a distinctive muscatel flavor. Known as the 'Champagne of Teas'.", Price: 450.0, Category: "Black Tea", ImageURL: "https://images.pexels.com/photos/1793034/pexels-photo-1793034.jpeg", InStock: true, Weight: "100g"},
		models.Product{ID: "c3d4e5f6-a7b8-9012-3456-7890abcdef01", Name: "Traditional Masala Chai", Description: "Our signature blend of black tea with cardamom, cinnamon, cloves, and ginger. A 60-year-old family recipe.", Price: 199.0, Category: "Masala Chai", ImageURL: "https://images.pexels.com/photos/5947062/pexels-photo-5947062.jpeg", InStock: true, Weight: "200g"},
		models.Product{ID: "d4e5f6a7-b8c9-0123-4567-890abcdef012", Name: "Royal Jaipur Blend", Description: "A premium blend inspired by royal traditions of Jaipur. Mix of fine Assam tea with aromatic spices.", Price: 399.0, Category: "Special Blends", ImageURL: "https://images.unsplash.com/photo-1625033405953-f20401c7d848", InStock: true, Weight: "150g"},
		models.Product{ID: "e5f6a7b8-c9d0-1234-5678-90abcdef0123", Name: "Green Tea Classic", Description: "Pure green tea leaves with natural antioxidants. Light, refreshing taste perfect for health-conscious tea lovers.", Price: 349.0, Category: "Green Tea", ImageURL: "https://images.unsplash.com/photo-1521136492500-e18f107709f7", InStock: true, Weight: "100g"},
		models.Product{ID: "f6a7b8c9-d0e1-2345-6789-0abcdef01234", Name: "Cardamom Tea", Description: "Aromatic tea infused with premium green cardamom. A classic favorite for its warming and soothing properties.", Price: 259.0, Category: "Flavored Tea", ImageURL: "https://images.pexels.com/photos/3904035/pexels-photo-3904035.jpeg", InStock: true, Weight: "100g"},
	}
	return s.Repository.SeedProducts(sampleProducts)
}
