import axios from 'axios';

const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${localStorage.getItem('token')}`,
  },
});

export const getClients = () => api.get('/clients').then(res => res.data);
export const addClient = (username) => api.post('/clients', { username }).then(res => res.data);
export const deleteClient = (clientId) => api.delete(`/clients/${clientId}`);
export const downloadConfig = (clientId) => api.get(`/clients/${clientId}/config`).then(res => res.data);