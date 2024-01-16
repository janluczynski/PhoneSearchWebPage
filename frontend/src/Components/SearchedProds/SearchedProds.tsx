import React, { useEffect, useState } from "react";
import { Product } from "../../types";
import CardProd from "../ProdCard/CardProd";
import { useQuery } from "@tanstack/react-query";
import { fetchProductsSearch } from "../../API/Api";
import { Button, Spinner } from "@chakra-ui/react";
import { AddIcon } from "@chakra-ui/icons";
import { SearchContext } from "../../Contexts/SearchContexts";
import { useContext } from "react";
import SortOptions from "../SortOptions/SortOptions";
type SearchedProdsProps = {
  searchTerm: string;
};

const SearchedProds: React.FC<SearchedProdsProps> = ({ searchTerm }) => {
  const [itemsToShow, setItemsToShow] = useState(5);
  const [isReady, setIsReady] = useState(false);
  const { order, sortValue, sortedBy } = useContext(SearchContext);
  useEffect(() => {
    setItemsToShow(5);
  }, [searchTerm]);

  const handleShowMore = () => {
    setItemsToShow((itemsToShow) => itemsToShow + 5);
  };
  const searchQuery = useQuery({
    queryKey: [
      "search",
      {
        searchedPhrase: searchTerm,
        sortBy: sortedBy,
        order: order,
        value: sortValue,
      },
    ],
    enabled: isReady,
    queryFn: async () => {
      if (typeof searchTerm === "string") {
        return fetchProductsSearch(searchTerm, sortedBy, order, sortValue);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });

  useEffect(() => {
    console.log(sortValue);
    if (searchTerm !== "") {
      setIsReady(true);
    }
  }, [searchTerm, sortValue, sortedBy, order]);
  return (
    <>
      {searchQuery.isLoading && (
        <span>
          <Spinner color="#860000" />
        </span>
      )}
      {searchQuery.error && <span>Error: {searchQuery.error.message}</span>}
      {searchQuery.data && (
        <>
          <div className="similarProd">
            {searchQuery.data
              .filter((product: Product) => product.price > 0)
              .slice(0, itemsToShow)
              .map((product: Product) => (
                <CardProd key={product.product_id} product={product} />
              ))}
          </div>
        </>
      )}
      {searchQuery.data && searchQuery.data.length > itemsToShow && (
        <Button className="c" onClick={handleShowMore} leftIcon={<AddIcon />}>
          Pokaż więcej
        </Button>
      )}
    </>
  );
};

export default SearchedProds;
