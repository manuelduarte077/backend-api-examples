import ProductList from "../../components/Productos/ProductList";
import SearchBar from "./SearchBar";

function ProductScreen() {
  return (
    <main className="flex-grow">
      {/* Search Bar */}
      <SearchBar />

      {/* Product List */}
      <ProductList />

    </main>
  );
}

export default ProductScreen;
