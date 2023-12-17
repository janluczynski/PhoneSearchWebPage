import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { Product } from "./types";
import Header from "./Components/Header/Header";
import Footer from "./Components/Footer/Footer";
import { useState } from "react";
import SearchedProds from "./Components/SearchedProds/SearchedProds";
function App() {
  const product: Product = {
    product_url: "https://www.ceneo.pl/143460739#tag=pp1",
    product_id: "d096efb3-9289-4d2d-8889-ab5af2a7d2f6",
    brand: "test",
    model: "testmodel",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2022/7/pr_2022_7_4_13_46_24_503_05.jpg",
    price: 2000,
    display: "6'",
    processor: "dobry",
    ram: 5,
    storage: 2,
    battery: 1000,
  };
  const [searchTerm, setSearchTerm] = useState("");
  return (
    <>
      <Header setSearchTerm={setSearchTerm} />
      <center>
        <div className="mainHeader">
          <h1>Wybierz najtańszy telefon dla siebie</h1>
        </div>
      </center>
      <SearchBar setSearchTerm={setSearchTerm} />
      <div className="searchedprods">
        <SearchedProds searchTerm={searchTerm} />
      </div>
      <center>
        <div className="waskie">
          <p>
            Witaj w naszej porównywarce cen telefonów, gdzie znajdziesz
            najlepsze oferty na najnowsze modele smartfonów z różnych
            sprawdzonych sklepów internetowych. Nasza intuicyjna platforma
            porównawcza umożliwia szybkie i łatwe znalezienie atrakcyjnych ofert
            bez konieczności przeszukiwania wielu stron. Skorzystaj z regularnie
            aktualizowanej bazy danych, pełnych specyfikacji technicznych,
            recenzji użytkowników i profesjonalnych opinii, aby maksymalnie
            zaoszczędzić na zakupie i cieszyć się nowym smartfonem w atrakcyjnej
            cenie!
          </p>
        </div>

        <h2 className="c">
          <i>Popularne produkty</i>
        </h2>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
      </center>
      <Footer />
    </>
  );
}

export default App;
