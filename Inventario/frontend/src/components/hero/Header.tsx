import { FaShoppingCart } from "react-icons/fa";

export default function Header() {
  return (
    <header className="bg-blue-600 text-white p-4 shadow-md">
      <div className="container mx-auto flex justify-between items-center">
        <h1 className="text-2xl font-bold">WonderStore</h1>
        <nav className="space-x-4 flex items-center">
          <a
            href="/"
            className="bg-white text-blue-600 py-2 px-4 rounded shadow hover:bg-gray-100"
          >
            Productos
          </a>

          <div className="relative group">
            <FaShoppingCart className="text-white text-2xl transition-transform duration-300 group-hover:scale-125 group-hover:text-yellow-300" />
            <span className="absolute -top-1 -right-3 bg-red-500 text-white rounded-full text-sm px-2">
              3{" "}
            </span>
          </div>
        </nav>
      </div>
    </header>
  );
}
