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
    imageURL:
      "https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2022/7/pr_2022_7_4_13_46_24_503_05.jpg",
    sale_price: "2000 zł",
    display: "6'",
    processor: "dobry",
    ram: "5gb",
    storage: "2gb",
    battery: "1000",
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
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
          <CardProd product={product} />
        </div>
        <div className="products">
          <CardProd product={product} />
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
