export interface Product {
  id: string;
  name: string;
  description: string;
  price: number;
  category: string;
  image_url: string;
  weight: string;
}

export interface CartItem extends Product {
  quantity: number;
}

export interface CustomerInfo {
  name: string;
  phone: string;
  email: string;
  address: string;
}

export interface OrderItem {
  product_id: string;
  quantity: number;
}

export interface OrderSuccess {
  order_id: string;
  message: string;
}
