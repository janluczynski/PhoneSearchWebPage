import React, { useState } from "react";
import "./Searchbar.css";
import Suggestions from "../Suggestions/Suggestions";
import { useContext } from "react";
import { SearchContext } from "../../Contexts/SearchContexts";

const SearchBar: React.FC = () => {
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
    <div>
      <form id="search-form" onSubmit={handleSubmit}>
        <input
          type="text"
          id="search-input"
          placeholder="Wyszukaj telefon..."
          value={inputValue}
          onChange={handleInputChange}
          autoComplete="off"
        />
      </form>
      <center>
        <Suggestions inputValue={inputValue} setSearchTerm={setSearchTerm} />
      </center>
    </div>
  );
};

export default SearchBar;
