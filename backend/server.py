from fastapi import FastAPI, HTTPException, Depends
from fastapi.middleware.cors import CORSMiddleware
from pymongo import MongoClient
from pydantic import BaseModel
from typing import List, Optional
import os
import uuid
from datetime import datetime

app = FastAPI()

# CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# MongoDB connection
mongo_url = os.environ.get('MONGO_URL', 'mongodb://localhost:27017/')
client = MongoClient(mongo_url)
db = client.mangal_chai_db

# Collections
products_collection = db.products
orders_collection = db.orders

# Pydantic models
class Product(BaseModel):
    id: str
    name: str
    description: str
    price: float
    category: str
    image_url: str
    in_stock: bool = True
    weight: str = "100g"

class CartItem(BaseModel):
    product_id: str
    quantity: int

class CustomerInfo(BaseModel):
    name: str
    phone: str
    email: str
    address: str

class Order(BaseModel):
    id: str
    customer_info: CustomerInfo
    items: List[CartItem]
    total_amount: float
    status: str = "pending"
    order_date: datetime
    notes: Optional[str] = None

# Initialize sample products on startup
@app.on_event("startup")
async def startup_event():
    # Check if products already exist
    if products_collection.count_documents({}) == 0:
        sample_products = [
            {
                "id": str(uuid.uuid4()),
                "name": "Premium Assam Black Tea",
                "description": "Rich, malty Assam tea with robust flavor. Perfect for morning tea with milk and sugar. Sourced from the finest tea gardens of Assam.",
                "price": 299.0,
                "category": "Black Tea",
                "image_url": "https://images.unsplash.com/photo-1563822249366-3efb23b8e0c9",
                "in_stock": True,
                "weight": "100g"
            },
            {
                "id": str(uuid.uuid4()),
                "name": "Darjeeling Muscatel",
                "description": "Delicate and aromatic Darjeeling tea with a distinctive muscatel flavor. Known as the 'Champagne of Teas'.",
                "price": 450.0,
                "category": "Black Tea",
                "image_url": "https://images.pexels.com/photos/1793034/pexels-photo-1793034.jpeg",
                "in_stock": True,
                "weight": "100g"
            },
            {
                "id": str(uuid.uuid4()),
                "name": "Traditional Masala Chai",
                "description": "Our signature blend of black tea with cardamom, cinnamon, cloves, and ginger. A 60-year-old family recipe.",
                "price": 199.0,
                "category": "Masala Chai",
                "image_url": "https://images.pexels.com/photos/5947062/pexels-photo-5947062.jpeg",
                "in_stock": True,
                "weight": "200g"
            },
            {
                "id": str(uuid.uuid4()),
                "name": "Royal Jaipur Blend",
                "description": "A premium blend inspired by royal traditions of Jaipur. Mix of fine Assam tea with aromatic spices.",
                "price": 399.0,
                "category": "Special Blends",
                "image_url": "https://images.unsplash.com/photo-1625033405953-f20401c7d848",
                "in_stock": True,
                "weight": "150g"
            },
            {
                "id": str(uuid.uuid4()),
                "name": "Green Tea Classic",
                "description": "Pure green tea leaves with natural antioxidants. Light, refreshing taste perfect for health-conscious tea lovers.",
                "price": 349.0,
                "category": "Green Tea",
                "image_url": "https://images.unsplash.com/photo-1521136492500-e18f107709f7",
                "in_stock": True,
                "weight": "100g"
            },
            {
                "id": str(uuid.uuid4()),
                "name": "Cardamom Tea",
                "description": "Aromatic tea infused with premium green cardamom. A classic favorite for its warming and soothing properties.",
                "price": 259.0,
                "category": "Flavored Tea",
                "image_url": "https://images.pexels.com/photos/3904035/pexels-photo-3904035.jpeg",
                "in_stock": True,
                "weight": "100g"
            }
        ]
        products_collection.insert_many(sample_products)
        print("Sample products inserted successfully")

# API Routes
@app.get("/api/products")
async def get_products():
    products = list(products_collection.find({}, {"_id": 0}))
    return products

@app.get("/api/products/{product_id}")
async def get_product(product_id: str):
    product = products_collection.find_one({"id": product_id}, {"_id": 0})
    if not product:
        raise HTTPException(status_code=404, detail="Product not found")
    return product

@app.get("/api/products/category/{category}")
async def get_products_by_category(category: str):
    products = list(products_collection.find({"category": category}, {"_id": 0}))
    return products

@app.post("/api/orders")
async def create_order(order_data: dict):
    # Validate products exist and calculate total
    total_amount = 0
    validated_items = []
    
    for item in order_data["items"]:
        product = products_collection.find_one({"id": item["product_id"]}, {"_id": 0})
        if not product:
            raise HTTPException(status_code=404, detail=f"Product {item['product_id']} not found")
        if not product["in_stock"]:
            raise HTTPException(status_code=400, detail=f"Product {product['name']} is out of stock")
        
        item_total = product["price"] * item["quantity"]
        total_amount += item_total
        validated_items.append({
            "product_id": item["product_id"],
            "product_name": product["name"],
            "price": product["price"],
            "quantity": item["quantity"],
            "subtotal": item_total
        })
    
    # Create order
    order = {
        "id": str(uuid.uuid4()),
        "customer_info": order_data["customer_info"],
        "items": validated_items,
        "total_amount": total_amount,
        "status": "pending",
        "order_date": datetime.now(),
        "notes": order_data.get("notes", "")
    }
    
    orders_collection.insert_one(order)
    return {"message": "Order placed successfully", "order_id": order["id"], "total_amount": total_amount}

@app.get("/api/orders/{order_id}")
async def get_order(order_id: str):
    order = orders_collection.find_one({"id": order_id}, {"_id": 0})
    if not order:
        raise HTTPException(status_code=404, detail="Order not found")
    return order

@app.get("/api/categories")
async def get_categories():
    categories = products_collection.distinct("category")
    return categories

@app.get("/api/health")
async def health_check():
    return {"status": "healthy", "message": "Mangal Chai API is running"}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8001)