import React from "react";

export const SearchContext = React.createContext({
  searchTerm: "",
  setSearchTerm: (term: string) => {},
  inputValue: "",
  setInputValue: (value: string) => {},
});
