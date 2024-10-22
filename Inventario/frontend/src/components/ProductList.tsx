import { useEffect, useState } from "react";
import { getProducts } from "../api/api";

interface Producto {
  id: number;
  nombre: string;
  precio: number;
  stock: number;
}

function ProductList() {
  const [products, setProducts] = useState<Producto[]>([]);

  useEffect(() => {
    getProducts()
      .then((response) => {
        setProducts(response.data);
      })
      .catch((error) => {
        console.error("Error al obtener productos:", error);
      });
  }, []);

  console.log("Productos:", products);
  console.log("Cantidad de productos:", products.length);

  return (
    <div>
      <h2>Lista de Productos</h2>
      <ul>
        {products.map((product) => (
          <li key={product.id}>
            {product.nombre} - ${product.precio} - Stock: {product.stock}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ProductList;
