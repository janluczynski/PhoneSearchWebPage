import React, { useState } from "react";
import "./Searchbar.css";
import { useContext } from "react";
import { SearchContext } from "../../Contexts/SearchContexts";
import {
  Input,
  InputGroup,
  InputLeftAddon,
  InputLeftElement,
} from "@chakra-ui/react";
import { Search2Icon } from "@chakra-ui/icons";

interface SearchBarProps {
  width?: string;
  height?: string;
  fontSize?: string;
}

const SearchBar: React.FC<SearchBarProps> = ({
  width = "auto",
  height = "auto",
  fontSize = "1em",
}) => {
  const { setSearchTerm, inputValue, setInputValue } =
    useContext(SearchContext);

  const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (inputValue.trim() !== "") {
      setSearchTerm(inputValue);
      setInputValue("");
    }
  };

  return (
    <form id="search-form" onSubmit={handleSubmit}>
      <InputGroup display="flex" justifyContent="center" alignItems="center">
        <InputLeftElement height={height} paddingLeft="5%" color="gray.300">
          <Search2Icon height={fontSize} width={fontSize} />
        </InputLeftElement>
        <Input
          type="text"
          id="search-input"
          placeholder="Wyszukaj telefon..."
          value={inputValue}
          onChange={handleInputChange}
          autoComplete="off"
          width={width}
          height={height}
          fontSize={fontSize}
          borderRadius="30px"
          paddingLeft="10%"
          backgroundColor={"white"}
          boxShadow="0px 4px 4px rgba(0, 0, 0, 0.25)"
          focusBorderColor="#860000"
        />
      </InputGroup>
    </form>
  );
};

export default SearchBar;
