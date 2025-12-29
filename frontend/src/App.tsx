import './App.css';

import type { CartItem, CustomerInfo, OrderSuccess, Product } from './types';
import { fetchCategories, fetchProducts } from './services/api';
import { useQuery } from '@tanstack/react-query';

import CartModal from './components/CartModal';
import CategoryFilter from './components/CategoryFilter';
import CheckoutModal from './components/CheckoutModal';
import HeroSection from './components/HeroSection';
import OrderSuccessModal from './components/OrderSuccessModal';
import ProductCard from './components/ProductCard';
import ProductModal from './components/ProductModal';
import { useState } from 'react';
import MainLayout from './layouts/MainLayout';

const App = () => {
  const [selectedCategory, setSelectedCategory] = useState<string>('all');
  const [cart, setCart] = useState<CartItem[]>([]);
  const [showCart, setShowCart] = useState<boolean>(false);
  const [showCheckout, setShowCheckout] = useState<boolean>(false);
  const [selectedProduct, setSelectedProduct] = useState<Product | null>(null);
  const [customerInfo, setCustomerInfo] = useState<CustomerInfo>({
    name: '',
    phone: '',  
    email: '',
    address: ''
  });
  const [orderSuccess, setOrderSuccess] = useState<OrderSuccess | null>(null);

  const { data: products, isLoading: isLoadingProducts, error: productsError } = useQuery<Product[]>({ queryKey: ['products'], queryFn: fetchProducts });
  const { data: categories, isLoading: isLoadingCategories, error: categoriesError } = useQuery<string[]>({ queryKey: ['categories'], queryFn: fetchCategories });

  const filteredProducts = selectedCategory === 'all' 
    ? products 
    : products?.filter(product => product.category === selectedCategory);

  const addToCart = (product: Product) => {
    const existingItem = cart.find(item => item.id === product.id);
    if (existingItem) {
      setCart(cart.map(item => 
        item.id === product.id 
          ? { ...item, quantity: item.quantity + 1 }
          : item
      ));
    } else {
      setCart([...cart, { ...product, quantity: 1 }]);
    }
  };

  const removeFromCart = (productId: string) => {
    setCart(cart.filter(item => item.id !== productId));
  };

  const updateQuantity = (productId: string, newQuantity: number) => {
    if (newQuantity === 0) {
      removeFromCart(productId);
    } else {
      setCart(cart.map(item => 
        item.id === productId 
          ? { ...item, quantity: newQuantity }
          : item
      ));
    }
  };

  const getTotalAmount = () => {
    return cart.reduce((total, item) => total + (item.price * item.quantity), 0);
  };

  if (isLoadingProducts || isLoadingCategories) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-orange-50 to-red-50 flex items-center justify-center">
        <div className="text-2xl text-orange-800 font-semibold">Loading Mangal Chai...</div>
      </div>
    );
  }

  if (productsError || categoriesError) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-orange-50 to-red-50 flex items-center justify-center">
        <div className="text-2xl text-red-800 font-semibold">Error loading data.</div>
      </div>
    );
  }

  return (
    <MainLayout cartLength={cart.length} onShowCart={() => setShowCart(true)} showCart={showCart} setShowCart={setShowCart}>
      <HeroSection />
      <CategoryFilter
        categories={categories || []}
        selectedCategory={selectedCategory}
        onSelectCategory={setSelectedCategory}
      />

      <section className="py-16">
        <div className="container mx-auto px-4">
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {filteredProducts?.map(product => (
              <ProductCard
                key={product.id}
                product={product}
                onAddToCart={addToCart}
                onViewDetails={setSelectedProduct}
              />
            ))}
          </div>
        </div>
      </section>

      <ProductModal
        product={selectedProduct}
        onClose={() => setSelectedProduct(null)}
        onAddToCart={addToCart}
      />

      <CartModal
        cart={cart}
        onClose={() => setShowCart(false)}
        onRemoveFromCart={removeFromCart}
        onUpdateQuantity={updateQuantity}
        onProceedToCheckout={() => {
          setShowCart(false);
          setShowCheckout(true);
        }}
        getTotalAmount={getTotalAmount}
      />

      <CheckoutModal
        showCheckout={showCheckout}
        customerInfo={customerInfo}
        cart={cart}
        onClose={() => setShowCheckout(false)}
        onCustomerInfoChange={setCustomerInfo}
        getTotalAmount={getTotalAmount}
      />

      <OrderSuccessModal
        orderSuccess={orderSuccess}
        onClose={() => setOrderSuccess(null)}
      />
    </MainLayout>
  );
};

export default App;
