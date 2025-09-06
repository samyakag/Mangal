import React from 'react';
import type { OrderSuccess } from '../types';

interface OrderSuccessModalProps {
  orderSuccess: OrderSuccess | null;
  onClose: () => void;
}

const OrderSuccessModal: React.FC<OrderSuccessModalProps> = ({ orderSuccess, onClose }) => {
  if (!orderSuccess) return null;

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-md w-full p-8 text-center">
        <div className="text-6xl mb-4">✅</div>
        <h2 className="text-2xl font-bold text-green-600 mb-4">Order Placed Successfully!</h2>
        <p className="text-gray-600 mb-4">
          Thank you for your order. We'll contact you soon to confirm the details.
        </p>
        <p className="text-sm text-gray-500 mb-6">
          Order ID: <span className="font-mono">{orderSuccess.order_id}</span>
        </p>
        <button
          onClick={onClose}
          className="bg-orange-600 text-white py-2 px-6 rounded-lg hover:bg-orange-500 transition-colors duration-300"
        >
          Continue Shopping
        </button>
      </div>
    </div>
  );
};

export default OrderSuccessModal;
