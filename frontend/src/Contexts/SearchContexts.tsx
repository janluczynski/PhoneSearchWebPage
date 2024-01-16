import React from "react";

export const SearchContext = React.createContext({
  searchTerm: "",
  setSearchTerm: (term: string) => {},
  inputValue: "",
  setInputValue: (value: string) => {},
  sortedBy: "",
  setSortedBy: (value: string) => {},
  order: 1,
  setOrder: (value: number) => {},
  sortValue: 0,
  setSortValue: (value: number) => {},
});
