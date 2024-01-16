import { Product } from "../types";
export const fetchProduct = async (product_id: string) => {
  const response = await fetch(
    `http://localhost:8080/parse/product?product_id=${product_id}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};

export const fetchProductsSearch = async (
  searchTerm: string,
  sortedBy: string,
  order: number,
) => {
  const response = await fetch(
    `http://localhost:8080/search?searchedPhrase=${searchTerm}&sortBy=${sortedBy}&order=${order}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();

  return data as unknown as Product[];
};
export const fetchSuggestions = async (searchTerm: string) => {
  const response = await fetch(
    `http://localhost:8080/searchbar?name=${searchTerm}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();

  return data as unknown as Product[];
};
export const fetchSimilarProducts = async (
  name: string,
  ram: number,
  storage: number,
) => {
  const response = await fetch(
    `http://localhost:8080/similar?name=${name}&ram=${ram}&storage=${storage}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};
export const fetchSameProducts = async (product_id: string) => {
  const response = await fetch(
    `http://localhost:8080/same/product?product_id=${product_id}`,
    {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};
