import Footer from "./components/hero/Footer";
import Header from "./components/hero/Header";
import ProductScreen from "./features/product/ProductScreen";

function App() {
  return (
    <div className="min-h-screen flex flex-col">
      {/* Header */}
      <Header />

      {/* Main */}
      <ProductScreen />

      {/* Footer */}
      <Footer />
    </div>
  );
}

export default App;
