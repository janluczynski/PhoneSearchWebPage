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
  value: number,
) => {
  const response = await fetch(
    `http://localhost:8080/search?searchedPhrase=${searchTerm}&sortBy=${sortedBy}&order=${order}&value=${value}`,
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
  brand: string,
  popularity?: number,
) => {
  const response = await fetch(
    `http://localhost:8080/similar?brand=${brand}&popularity=${popularity}`,
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
export const increaseProductViews = async (product_id: string) => {
  const response = await fetch(
    `http://localhost:8080/increment?id=${product_id}`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
    },
  );
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
};
export const fetchTopProducts = async () => {
  const response = await fetch(`http://localhost:8080/top-products`, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
    },
  });
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }
  const data = await response.json();
  return data;
};
