import axios from "axios";

import { config } from "./config";

export const TOKEN_KEY = "hospital_token";

export const getToken = (): string | null => {
  if (typeof window === "undefined") {
    return null;
  }
  return window.localStorage.getItem(TOKEN_KEY);
};

export const setToken = (token: string): void => {
  window.localStorage.setItem(TOKEN_KEY, token);
  document.cookie = `${TOKEN_KEY}=${encodeURIComponent(token)}; path=/`;
};

export const clearToken = (): void => {
  window.localStorage.removeItem(TOKEN_KEY);
  document.cookie = `${TOKEN_KEY}=; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT`;
};

const client = axios.create({
  baseURL: config.api.baseUrl,
  timeout: config.api.timeout,
  headers: {
    "Content-Type": "application/json",
  },
});

client.interceptors.request.use((req) => {
  const token = getToken();
  if (token) {
    req.headers.Authorization = `Bearer ${token}`;
  }
  return req;
});

client.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      clearToken();
    }
    return Promise.reject(error);
  },
);

export default client;
