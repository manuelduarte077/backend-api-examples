import { useEffect, useState } from "react";
import { getProducts } from "../../api/api";
import Producto from "../../interface/Producto";
import { FaSpinner } from "react-icons/fa";

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

  return products.length === 0 ? (
    <div className="flex flex-col items-center justify-center h-[70vh]">
      <FaSpinner className="text-4xl text-gray-500 animate-spin mb-4" />
      <p className="text-gray-500 text-lg">Cargando productos...</p>
    </div>
  ) : (
    <div className="container mx-auto py-8">
      <h2 className="text-2xl font-bold mb-6">
        Productos que te pueden interesar
      </h2>
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        {products.map((product) => (
          <div
            key={product.id}
            className="bg-white rounded-lg shadow-lg p-4 hover:shadow-xl transition-shadow duration-300"
          >
            {/* Imagen del producto */}
            <div className="relative mb-4">
              <img
                src={"https://placehold.co/600x400.png"}
                alt={product.nombre}
                className="w-full h-48 object-cover rounded-lg"
              />
              <span className="absolute bottom-2 right-2 bg-gray-900 text-white text-lg font-bold px-2 py-1 rounded-lg">
                ${product.precio}
              </span>
            </div>

            <h3 className="text-xl font-semibold mb-1">{product.nombre}</h3>
            <p className="text-gray-500 mb-4">{product.descripcion}</p>

            <button className="w-full bg-gray-200 text-gray-700 py-2 rounded-lg hover:bg-gray-300 transition-colors duration-200">
              Add to bag
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ProductList;
