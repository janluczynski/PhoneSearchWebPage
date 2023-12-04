import Layout from "./Components/Layout/Layout";
import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { useEffect, useState } from "react";
import ProductPage from "./Components/ProductPage/ProductPage";
import { useMutation, useQuery } from "@tanstack/react-query";
import searchQuery from "./Components/Searchbar/Searchbar";

function App() {
  type Product = {
    product_url: string;
    product_id: string;
    brand: string;
    model: string;
    imageURL: string;
    price: string;
    display: string;
    processor: string;
    ram: string;
    storage: string;
    battery: string;
  };
  const product: Product = {
    product_url: "https://www.ceneo.pl/143460739#tag=pp1",
    product_id: "d096efb3-9289-4d2d-8889-ab5af2a7d2f6",
    brand: "test",
    model: "testmodel",
    imageURL:
      "https://image.ceneostatic.pl/data/products/143460739/f-haxe-elektryczna-banka-antycellulitowa-hx801.jpg",
    price: "2000",
    display: "6'",
    processor: "dobry",
    ram: "5gb",
    storage: "2gb",
    battery: "1000",
  };

  return (
    <Layout>
      <SearchBar />

      <center>
        <div className="waskie">
          <p>
            WWitaj na naszej porównywarce cen telefonów - miejscu, gdzie
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
        <div className="products">
          <CardProd product={product} />
        </div>
        </center>


    </Layout>
  );
}

export default App;
