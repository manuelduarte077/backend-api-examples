import { FaSearch } from "react-icons/fa";

export default function SearchBar() {
  return (
    <div className="container mx-auto mt-8">
      <div className="relative flex items-center">
        <FaSearch className="absolute left-3 text-gray-500" />
        <input
          type="text"
          placeholder="Buscar productos..."
          className="w-full pl-10 p-2 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition duration-200"
        />
      </div>
    </div>
  );
}
