import React, { type Dispatch, type SetStateAction } from 'react';
import useRazorpay from '../hooks/useRazorpay';
import type { CartItem, CustomerInfo } from '../types';

interface CheckoutModalProps {
  showCheckout: boolean;
  customerInfo: CustomerInfo;
  cart: CartItem[];
  onClose: () => void;
  onCustomerInfoChange: Dispatch<SetStateAction<CustomerInfo>>;
  getTotalAmount: () => number;
}

const CheckoutModal: React.FC<CheckoutModalProps> = ({
  showCheckout,
  customerInfo,
  cart,
  onClose,
  onCustomerInfoChange,
  getTotalAmount,
}) => {
  const { openRazorpay } = useRazorpay();

  if (!showCheckout) return null;

  const handleCheckout = () => {
    openRazorpay(cart);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div className="bg-white rounded-2xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div className="p-6 border-b">
          <div className="flex justify-between items-center">
            <h2 className="text-2xl font-bold text-gray-800">Checkout</h2>
            <button
              onClick={onClose}
              className="text-gray-500 hover:text-gray-700 text-2xl font-bold"
            >
              ×
            </button>
          </div>
        </div>
        <div className="p-6">
          <div className="mb-6">
            <h3 className="text-lg font-semibold mb-4">Customer Information</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <input
                type="text"
                placeholder="Full Name"
                value={customerInfo.name}
                onChange={(e) => onCustomerInfoChange({...customerInfo, name: e.target.value})}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              />
              <input
                type="tel"
                placeholder="Phone Number"
                value={customerInfo.phone}
                onChange={(e) => onCustomerInfoChange({...customerInfo, phone: e.target.value})}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              />
              <input
                type="email"
                placeholder="Email Address"
                value={customerInfo.email}
                onChange={(e) => onCustomerInfoChange({...customerInfo, email: e.target.value})}
                className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent"
              />
            </div>
            <textarea
              placeholder="Delivery Address"
              value={customerInfo.address}
              onChange={(e) => onCustomerInfoChange({...customerInfo, address: e.target.value})}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-orange-500 focus:border-transparent mt-4"
              rows={3}
            />
          </div>
          
          <div className="mb-6">
            <h3 className="text-lg font-semibold mb-4">Order Summary</h3>
            {cart.map(item => (
              <div key={item.id} className="flex justify-between items-center py-2">
                <span>{item.name} × {item.quantity}</span>
                <span>₹{item.price * item.quantity}</span>
              </div>
            ))}
            <div className="border-t pt-2 mt-2">
              <div className="flex justify-between items-center text-xl font-bold">
                <span>Total: ₹{getTotalAmount()}</span>
              </div>
            </div>
          </div>

          <button
            onClick={handleCheckout}
            disabled={!customerInfo.name || !customerInfo.phone || !customerInfo.address}
            className="w-full bg-gradient-to-r from-green-600 to-green-700 text-white py-3 px-6 rounded-lg hover:from-green-500 hover:to-green-600 transition-all duration-300 font-semibold disabled:opacity-50 disabled:cursor-not-allowed"
          >
            Place Order
          </button>
        </div>
      </div>
    </div>
  );
};

export default CheckoutModal;
