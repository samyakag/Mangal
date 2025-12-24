
import { CartItem } from "../types";

declare global {
  interface Window {
    Razorpay: any;
  }
}

const useRazorpay = () => {
  const openRazorpay = async (cartItems: CartItem[]) => {
    try {
      const itemsToOrder = cartItems.map(item => ({
        product_id: item.id,
        quantity: item.quantity,
      }));

      const apiBaseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8001/api';
      const response = await fetch(`${apiBaseUrl}/payments/create-order`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "include",
        body: JSON.stringify({ items: itemsToOrder }),
      });

      if (!response.ok) {
        throw new Error("Failed to create Razorpay order");
      }

      const orderDetails = await response.json();

      const options = {
        key: orderDetails.key_id,
        amount: orderDetails.amount,
        currency: orderDetails.currency,
        name: "Mangal Chai",
        description: "Test Transaction",
        order_id: orderDetails.order_id,
        handler: function (response: any) {
          alert(`Payment successful. Payment ID: ${response.razorpay_payment_id}`)
          // You can handle the successful payment here
          // e.g., redirect to a success page, clear the cart, etc.
        },
        prefill: {
          name: "Test User",
          email: "test.user@example.com",
          contact: "9999999999",
        },
        notes: {
          address: "Test Address",
        },
        theme: {
          color: "#3399cc",
        },
      };

      const rzp = new window.Razorpay(options);
      rzp.open();
    } catch (error) {
      console.error("Error opening Razorpay checkout:", error);
      alert("Failed to open Razorpay checkout.");
    }
  };

  return { openRazorpay };
};

export default useRazorpay;
