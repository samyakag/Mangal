import React from 'react';
import type { Product } from '../types';

interface ProductCardProps {
  product: Product;
  onAddToCart: (product: Product) => void;
  onViewDetails: (product: Product) => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onAddToCart, onViewDetails }) => {
  return (
    <div key={product.id} className="bg-white rounded-2xl shadow-lg overflow-hidden hover:shadow-2xl transition-all duration-300">
      <div className="relative h-64 bg-gradient-to-br from-orange-100 to-red-100">
        <img
          src={product.image_url}
          alt={product.name}
          className="w-full h-full object-cover"
        />
        <div className="absolute top-4 right-4 bg-orange-600 text-white px-3 py-1 rounded-full text-sm font-semibold">
          ₹{product.price}
        </div>
      </div>
      <div className="p-6">
        <h3 className="text-xl font-bold text-gray-800 mb-2">{product.name}</h3>
        <p className="text-gray-600 mb-4 text-sm leading-relaxed">{product.description}</p>
        <div className="flex justify-between items-center mb-4">
          <span className="text-sm text-orange-600 font-medium">{product.category}</span>
          <span className="text-sm text-gray-500">{product.weight}</span>
        </div>
        <div className="flex gap-2">
          <button
            onClick={() => onViewDetails(product)}
            className="flex-1 bg-gray-100 text-gray-700 py-2 px-4 rounded-lg hover:bg-gray-200 transition-colors duration-300 font-medium"
          >
            View Details
          </button>
          <button
            onClick={() => onAddToCart(product)}
            className="flex-1 bg-gradient-to-r from-orange-600 to-red-600 text-white py-2 px-4 rounded-lg hover:from-orange-500 hover:to-red-500 transition-all duration-300 font-semibold"
          >
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
