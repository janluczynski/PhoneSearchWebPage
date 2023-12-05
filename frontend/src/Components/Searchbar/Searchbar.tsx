import React, { useState } from "react";
import "./Searchbar.css";
import { useQuery } from "@tanstack/react-query";
import CardProd from "../ProdCard/CardProd";
const SearchBar = () => {
  const [itemsToShow, setItemsToShow] = useState(5);

  const handleShowMore = () => {
    setItemsToShow(itemsToShow + 5);
  };
  const [inputValue, setInputValue] = useState("");
  const [searchTerm, setSearchTerm] = useState("");

  const searchQuery = useQuery({
    queryKey: ["search", searchTerm],
    queryFn: async () => {
      const response = await fetch("http://localhost:8080/search", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          SearchedPhrase: searchTerm,
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      console.log(data);
      return data;
    },
    enabled: searchTerm !== "", // The query will not run until searchTerm is not an empty string
  });

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setSearchTerm(inputValue);
  };

  return (
    <div>
      <form id="search-form" onSubmit={handleSubmit}>
        <input
          type="text"
          id="search-input"
          placeholder="Wyszukaj telefon..."
          value={inputValue}
          onChange={handleInputChange}
        />
      </form>
      {searchQuery.isLoading && <span>Loading...</span>}
      {searchQuery.error && <span>Error: {searchQuery.error.message}</span>}
      {searchQuery.data && (
        <div className="similarProd">
          {searchQuery.data.slice(0, itemsToShow).map((product: any) => (
            <CardProd key={product.product_id} product={product} />
          ))}
        </div>
      )}
      {location.pathname === "xd" && (
        <button className="c" onClick={handleShowMore}>
          Pokaż więcej
        </button>
      )}
    </div>
  );
};

export default SearchBar;
