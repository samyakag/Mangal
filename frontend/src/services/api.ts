import axios from 'axios';

const backendUrl = process.env.REACT_APP_BACKEND_URL || 'http://localhost:8001';

const api = axios.create({
  baseURL: backendUrl,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const fetchProducts = async () => {
  const response = await api.get('/api/products');
  return response.data;
};

export const fetchCategories = async () => {
  const response = await api.get('/api/categories');
  return response.data;
};

export const createOrder = async (orderData: any) => {
  const response = await api.post('/api/orders', orderData);
  return response.data;
};

export default api;
