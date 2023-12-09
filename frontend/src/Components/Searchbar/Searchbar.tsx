import React, { useState } from "react";
import "./Searchbar.css";
type SearchBarProps = {
  setSearchTerm: (searchTerm: string) => void;
};
const SearchBar: React.FC<SearchBarProps> = ({ setSearchTerm }) => {
  const [itemsToShow, setItemsToShow] = useState(5);

  const [inputValue, setInputValue] = useState("");

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    setSearchTerm(inputValue);
    setItemsToShow(5);
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
        />
      </form>
    </div>
  );
};

export default SearchBar;
