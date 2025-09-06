import React from 'react';
import type { Product } from '../types';

interface ProductModalProps {
  product: Product | null;
  onClose: () => void;
  onAddToCart: (product: Product) => void;
}

const ProductModal: React.FC<ProductModalProps> = ({ product, onClose, onAddToCart }) => {
  if (!product) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div className="relative">
          <img
            src={product.image_url}
            alt={product.name}
            className="w-full h-64 object-cover rounded-t-2xl"
          />
          <button
            onClick={onClose}
            className="absolute top-4 right-4 bg-white bg-opacity-90 hover:bg-opacity-100 w-10 h-10 rounded-full flex items-center justify-center text-xl font-bold transition-all duration-300"
          >
            ×
          </button>
        </div>
        <div className="p-6">
          <h2 className="text-2xl font-bold text-gray-800 mb-4">{product.name}</h2>
          <p className="text-gray-600 mb-4 leading-relaxed">{product.description}</p>
          <div className="grid grid-cols-2 gap-4 mb-6">
            <div>
              <span className="text-sm text-gray-500">Category</span>
              <p className="font-semibold text-orange-600">{product.category}</p>
            </div>
            <div>
              <span className="text-sm text-gray-500">Weight</span>
              <p className="font-semibold">{product.weight}</p>
            </div>
            <div>
              <span className="text-sm text-gray-500">Price</span>
              <p className="font-bold text-2xl text-green-600">₹{product.price}</p>
            </div>
            <div>
              <span className="text-sm text-gray-500">Availability</span>
              <p className="font-semibold text-green-600">In Stock</p>
            </div>
          </div>
          <button
            onClick={() => {
              onAddToCart(product);
              onClose();
            }}
            className="w-full bg-gradient-to-r from-orange-600 to-red-600 text-white py-3 px-6 rounded-lg hover:from-orange-500 hover:to-red-500 transition-all duration-300 font-semibold text-lg"
          >
            Add to Cart
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductModal;
