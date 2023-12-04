import Layout from "./Components/Layout/Layout";
import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { useEffect, useState } from "react";
import ProductPage from "./Components/ProductPage/ProductPage";
import { useMutation, useQuery } from "@tanstack/react-query";

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
    product_id: "1",
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
  const productsQuery = useQuery({
    queryKey: ["products"],
    queryFn: async () => {
      const response = await fetch("http://localhost:8080/parse/product", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          product_id: "d096efb3-9289-4d2d-8889-ab5af2a7d2f6",
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      console.log(data.product_url);
      console.log(data.product_id);
      return response.json();
    },
  });

  // fetch("http://localhost:8080/parse/product",{
  //   method: 'POST',
  //   headers: {
  //     'Content-Type': 'application/json',
  //   },
  //   body: JSON.stringify({
  //     product_id: 'd096efb3-9289-4d2d-8889-ab5af2a7d2f6',
  //   }),

  // })
  //   .then((response) => {
  //     if (!response.ok) {
  //       throw new Error(`HTTP error! status: ${response.status}`);
  //     }
  //     return response.json();
  //   })
  //   .then((data) => {
  //     console.log(data);
  //     const transformedData = data.map((item: any) => ({
  //       product_id: item.product_id,
  //       brand: item.brand,
  //       model: item.model,
  //       price: item.price,
  //       display: item.display,
  //       processor: item.processor,
  //       ram: item.ram,
  //       storage: item.storage,
  //       battery: item.battery,
  //       imageURL: item.imageURL,
  //       product_url: item.product_url,
  //     }));
  //       setProducts(transformedData);

  //   })
  // if(productsQuery.isLoading) return <div>Loading...</div>

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

      {/* {products.map((product) => (
  <div key={product.product_id}>
    <h2>{product.brand} {product.model}</h2>
    <img src={product.imageURL} alt={product.model} />
    <p>Price: {product.price}</p>
    <p>Display: {product.display}</p>
    <p>Processor: {product.processor}</p>
    <p>RAM: {product.ram}</p>
    <p>Storage: {product.storage}</p>
    <p>Battery: {product.battery}</p>
    <a href={product.product_url}>Buy now</a>
  </div>
))} */}
    </Layout>
  );
}

export default App;
