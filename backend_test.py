import requests
import json
import unittest
import os
import sys

# Get the backend URL from the frontend .env file
def get_backend_url():
    with open('/app/frontend/.env', 'r') as f:
        for line in f:
            if line.startswith('REACT_APP_BACKEND_URL='):
                return line.strip().split('=')[1].strip('"\'')
    return None

# Base URL for API requests
BASE_URL = f"{get_backend_url()}/api"
print(f"Using backend URL: {BASE_URL}")

class MangalChaiBackendTests(unittest.TestCase):
    """Test suite for Mangal Chai backend API endpoints"""
    
    def setUp(self):
        """Setup before each test"""
        self.product_id = None  # Will be set after getting products
        self.order_id = None    # Will be set after creating an order
    
    def test_01_health_check(self):
        """Test the health check endpoint"""
        response = requests.get(f"{BASE_URL}/health")
        self.assertEqual(response.status_code, 200)
        data = response.json()
        self.assertEqual(data["status"], "healthy")
        print("✅ Health check endpoint working")
    
    def test_02_get_all_products(self):
        """Test getting all products"""
        response = requests.get(f"{BASE_URL}/products")
        self.assertEqual(response.status_code, 200)
        products = response.json()
        self.assertIsInstance(products, list)
        self.assertGreater(len(products), 0)
        
        # Store a product ID for later tests
        self.product_id = products[0]["id"]
        print(f"✅ Get all products endpoint working - Found {len(products)} products")
    
    def test_03_get_product_by_id(self):
        """Test getting a specific product by ID"""
        # First get all products to get a valid ID
        if not self.product_id:
            self.test_02_get_all_products()
        
        response = requests.get(f"{BASE_URL}/products/{self.product_id}")
        self.assertEqual(response.status_code, 200)
        product = response.json()
        self.assertEqual(product["id"], self.product_id)
        print(f"✅ Get product by ID endpoint working - Found product: {product['name']}")
        
        # Test with invalid product ID
        response = requests.get(f"{BASE_URL}/products/invalid-id")
        self.assertEqual(response.status_code, 404)
        print("✅ Get product by ID correctly handles invalid IDs")
    
    def test_04_get_categories(self):
        """Test getting all categories"""
        response = requests.get(f"{BASE_URL}/categories")
        self.assertEqual(response.status_code, 200)
        categories = response.json()
        self.assertIsInstance(categories, list)
        self.assertGreater(len(categories), 0)
        print(f"✅ Get categories endpoint working - Found {len(categories)} categories")
    
    def test_05_get_products_by_category(self):
        """Test getting products by category"""
        # First get categories
        response = requests.get(f"{BASE_URL}/categories")
        self.assertEqual(response.status_code, 200)
        categories = response.json()
        
        # Test with first category
        category = categories[0]
        response = requests.get(f"{BASE_URL}/products/category/{category}")
        self.assertEqual(response.status_code, 200)
        products = response.json()
        self.assertIsInstance(products, list)
        
        # Verify all products are of the requested category
        for product in products:
            self.assertEqual(product["category"], category)
        
        print(f"✅ Get products by category endpoint working - Found {len(products)} products in category '{category}'")
    
    def test_06_create_order(self):
        """Test creating a new order"""
        # First get a product ID
        if not self.product_id:
            self.test_02_get_all_products()
        
        # Create order data
        order_data = {
            "customer_info": {
                "name": "Raj Sharma",
                "phone": "9876543210",
                "email": "raj.sharma@example.com",
                "address": "123 Gandhi Road, Jaipur, Rajasthan"
            },
            "items": [
                {
                    "product_id": self.product_id,
                    "quantity": 2
                }
            ]
        }
        
        response = requests.post(f"{BASE_URL}/orders", json=order_data)
        self.assertEqual(response.status_code, 200)
        result = response.json()
        self.assertIn("order_id", result)
        self.assertIn("total_amount", result)
        
        # Store order ID for next test
        self.order_id = result["order_id"]
        print(f"✅ Create order endpoint working - Order ID: {self.order_id}")
        
        # Test with invalid product ID
        invalid_order_data = {
            "customer_info": {
                "name": "Raj Sharma",
                "phone": "9876543210",
                "email": "raj.sharma@example.com",
                "address": "123 Gandhi Road, Jaipur, Rajasthan"
            },
            "items": [
                {
                    "product_id": "invalid-id",
                    "quantity": 2
                }
            ]
        }
        
        response = requests.post(f"{BASE_URL}/orders", json=invalid_order_data)
        self.assertEqual(response.status_code, 404)
        print("✅ Create order correctly handles invalid product IDs")
        
        # Test with missing customer info
        incomplete_order_data = {
            "items": [
                {
                    "product_id": self.product_id,
                    "quantity": 2
                }
            ]
        }
        
        response = requests.post(f"{BASE_URL}/orders", json=incomplete_order_data)
        self.assertNotEqual(response.status_code, 200)
        print("✅ Create order correctly handles missing customer info")
    
    def test_07_get_order_by_id(self):
        """Test getting an order by ID"""
        # First create an order if we don't have an order ID
        if not self.order_id:
            self.test_06_create_order()
        
        response = requests.get(f"{BASE_URL}/orders/{self.order_id}")
        self.assertEqual(response.status_code, 200)
        order = response.json()
        self.assertEqual(order["id"], self.order_id)
        print(f"✅ Get order by ID endpoint working - Found order for: {order['customer_info']['name']}")
        
        # Test with invalid order ID
        response = requests.get(f"{BASE_URL}/orders/invalid-id")
        self.assertEqual(response.status_code, 404)
        print("✅ Get order by ID correctly handles invalid IDs")

if __name__ == "__main__":
    # Run the tests
    unittest.main(argv=['first-arg-is-ignored'], exit=False)
