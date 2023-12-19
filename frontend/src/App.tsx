import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { Product } from "./types";
import Header from "./Components/Header/Header";
import Footer from "./Components/Footer/Footer";
import { useState } from "react";
import SearchedProds from "./Components/SearchedProds/SearchedProds";
import { SearchContext } from "./Contexts/SearchContexts";
import Suggestions from "./Components/Suggestions/Suggestions";
function App() {
  const product1: Product = {
    product_url:
      "https://www.x-kom.pl/p/1180085-smartfon-telefon-apple-iphone-15-pro-max-256gb-black-titanium.html",
    product_id: "c20e654c-cbeb-40be-aa68-f0cba7ffcfdd",
    name: "Apple iPhone 15 PRO MAX",
    brand: "Apple",
    model: "iPhone 15 PRO MAX ",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/9/pr_2023_9_12_23_5_43_479_00.jpg",
    price: 7199,
    display: "6.7",
    processor: "Apple A17 Pro",
    ram: 0,
    storage: 262144,
    battery: 0,
  };
  const product2: Product = {
    product_url:
      "https://www.x-kom.pl/p/1180055-smartfon-telefon-apple-iphone-15-plus-2…",
    product_id: "ec98138c-613a-4aa3-9921-fadd39c3c972",
    name: "Apple iPhone 15 Plus",
    brand: "Apple",
    model: "iPhone 15 Plus ",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/9/pr_2023_9_13_0_36_38_213_00.jpg",
    price: 5599,
    display: "6.7",
    processor: "Apple A16 Bionic",
    ram: 0,
    storage: 262144,
    battery: 0,
  };
  const product3: Product = {
    product_url:
      "https://www.x-kom.pl/p/1158859-smartfon-telefon-samsung-galaxy-z-fold5…",
    product_id: "badd47eb-5934-4311-9053-cd50e1007de7",
    name: "Samsung Galaxy Z Fold5",
    brand: "Samsung",
    model: "Galaxy Z Fold5",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/7/pr_2023_7_18_11_12_29_471_00.jpg",
    price: 9799,
    display: "7.6 (ekran po rozłożeniu)6.2",
    processor:
      "Qualcomm Snapdragon 8 gen 2 (1x 3.2 GHz, X3 + 4x 2.8 GHz, A71 + 3x 2.0 GHz, A51)",
    ram: 12288,
    storage: 1048576,
    battery: 4400,
  };
  const [searchTerm, setSearchTerm] = useState("");
  const [inputValue, setInputValue] = useState("");
  return (
    <>
      <SearchContext.Provider
        value={{ searchTerm, setSearchTerm, inputValue, setInputValue }}
      >
        <Header />
        <center>
          <div className="fancy">
            <h1>Wybierz najtańszy telefon dla siebie</h1>
          </div>
        </center>
        <div className="searchContainer">
          <SearchBar />
          <div id="portal1">
            <Suggestions portalId="portal1" />
          </div>
        </div>
        <div className="searchedprods">
          <SearchedProds searchTerm={searchTerm} />
        </div>
        <center>
          <div className="waskie">
            <p>
              PhoneCompass to strona, która pomoże Ci znaleźć najlepsze telefony
              w najniższych cenach. Nasza strona oferuje łatwe i intuicyjne
              narzędzia, które pomogą Ci znaleźć telefon, który spełni Twoje
              wymagania. Dzięki naszej pomocy, nie musisz tracić czasu na
              szukanie najlepszych ofert.
            </p>
          </div>

          <h2 className="c">
            <i>Popularne produkty</i>
          </h2>
          <div className="products">
            <CardProd product={product1} />
            <CardProd product={product2} />
            <CardProd product={product3} />
          </div>
        </center>
        {/* <Footer /> */}
      </SearchContext.Provider>
    </>
  );
}

export default App;
