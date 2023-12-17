import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { Product } from "./types";
import Header from "./Components/Header/Header";
import Footer from "./Components/Footer/Footer";
import { useState } from "react";
import SearchedProds from "./Components/SearchedProds/SearchedProds";
function App() {
  const product1: Product = {
    product_url:
      "https://www.x-kom.pl/p/1180085-smartfon-telefon-apple-iphone-15-pro-max-256gb-black-titanium.html",
    product_id: "813e2688-f6ab-4c92-887f-510ad99fa483",
    brand: "Apple",
    model: "iPhone 15 PRO MAX ",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/9/pr_2023_9_12_23_5_43_479_00.jpg",
    price: 7199,
    display: "6.7",
    processor: "Apple A17 Pro",
    ram: 0,
    storage: 256,
    battery: 0,
  };
  const product2: Product = {
    product_url:
      "https://www.x-kom.pl/p/1180055-smartfon-telefon-apple-iphone-15-plus-2…",
    product_id: "c0b94f44-bcb2-49b9-9c34-16a98191cb4a",
    brand: "Apple",
    model: "iPhone 15 Plus ",
    image:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/9/pr_2023_9_13_0_36_38_213_00.jpg",
    price: 5599,
    display: "6.7",
    processor: "Apple A16 Bionic",
    ram: 0,
    storage: 256,
    battery: 0,
  };
  const product3: Product = {
    product_url:
      "https://www.x-kom.pl/p/1158859-smartfon-telefon-samsung-galaxy-z-fold5…",
    product_id: "c9c4a4d5-a4af-47f0-9c91-89fb25da92c6",
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
  return (
    <>
      <Header setSearchTerm={setSearchTerm} />
      <center>
        <h2>Wybierz najtańszy telefon dla siebie</h2>
      </center>
      <SearchBar setSearchTerm={setSearchTerm} />
      <div className="searchedprods">
        <SearchedProds searchTerm={searchTerm} />
      </div>
      <center>
        <div className="waskie">
          <p>
            Witaj na naszej porównywarce cen telefonów - miejscu, gdzie
            znajdziesz najkorzystniejsze oferty na najnowsze modele smartfonów!
            Jesteśmy tutaj po to, aby ułatwić Ci wybór idealnego telefonu,
            dostarczając kompleksowe porównania cenowe ze sprawdzonych sklepów
            internetowych. Dzięki naszej intuicyjnej platformie porównywawczej,
            możesz szybko i łatwo znaleźć najlepsze oferty na telefony mobilne,
            bez konieczności przeszukiwania wielu stron internetowych. Nasza
            baza danych jest regularnie aktualizowana, zapewniając Ci dostęp do
            najnowszych informacji o cenach i promocjach. Znajdziesz u nas pełne
            specyfikacje techniczne każdego modelu, recenzje użytkowników oraz
            profesjonalne opinie, co pozwoli Ci dokładnie zorientować się w
            dostępnych opcjach. Niezależnie od tego, czy szukasz najnowszego
            flagowca czy też bardziej przystępnego cenowo modelu, nasza
            porównywarka pomoże Ci znaleźć najlepszą ofertę na rynku. Dodatkowo,
            oferujemy informacje o dodatkowych promocjach, rabatach i darmowej
            dostawie, dzięki czemu możesz maksymalnie zaoszczędzić na zakupie
            swojego wymarzonego telefonu. Korzystaj z naszej porównywarki cen i
            ciesz się nowym smartfonem w atrakcyjnej cenie! Nie trać czasu na
            przeszukiwanie różnych stron - znajdź najlepszą ofertę już teraz i
            ułatw sobie zakupy dzięki naszej porównywarce cen telefonów!
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
      <Footer />
    </>
  );
}

export default App;
