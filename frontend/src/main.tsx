import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import ProductPage from "./Components/ProductPage/ProductPage.tsx";
import "./index.css";
import { ChakraProvider } from "@chakra-ui/react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { BrowserRouter, Routes, Route } from "react-router-dom";
const queryClient = new QueryClient();
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
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <BrowserRouter>
      <QueryClientProvider client={queryClient}>
        <ChakraProvider>
          <Routes>
            <Route path="/" element={<App />} />
            <Route path="/product/:id" element={<ProductPage />} />
          </Routes>
          <ReactQueryDevtools />
        </ChakraProvider>
      </QueryClientProvider>
    </BrowserRouter>
  </React.StrictMode>,
);
