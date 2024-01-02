import React, { useState } from "react";
import "./Searchbar.css";
import { useContext } from "react";
import { SearchContext } from "../../Contexts/SearchContexts";

interface SearchBarProps {
  width?: string;
  height?: string;
  fontSize?: string;
}

const SearchBar: React.FC<SearchBarProps> = ({ width = 'auto', height = 'auto', fontSize = '1em' }) => {
  const { setSearchTerm, inputValue, setInputValue } =
    useContext(SearchContext);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setSearchTerm(inputValue);
    setInputValue("");
  };

  return (
    <form id="search-form" onSubmit={handleSubmit}>
      <input
        type="text"
        id="search-input"
        placeholder="Wyszukaj telefon..."
        value={inputValue}
        onChange={handleInputChange}
        autoComplete="off"
        style={{ width: width, height: height, fontSize: fontSize }}
      />
    </form>
  );
};

export default SearchBar;
