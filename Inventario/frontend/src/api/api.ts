import axios from "axios";

const API_URL = "http://localhost:8080/api/v1/productos";

export const getProducts = async () => {
  console.log("API_URL", API_URL);
  return await axios.get(API_URL);
};

export const getProduct = async (id: string) => {
  return await axios.get(`${API_URL}/${id}`);
};
