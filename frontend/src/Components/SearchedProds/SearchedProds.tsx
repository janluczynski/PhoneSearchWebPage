import React, { useEffect, useState } from "react";
import { Product } from "../../types";
import CardProd from "../ProdCard/CardProd";
import { useQuery } from "@tanstack/react-query";
import { fetchProductsSearch } from "../../API/Api";
import { Button, Spinner } from "@chakra-ui/react";
import { AddIcon } from "@chakra-ui/icons";
import SortOptions from "../SortOptions/SortOptions";
type SearchedProdsProps = {
  searchTerm: string;
};

const SearchedProds: React.FC<SearchedProdsProps> = ({ searchTerm }) => {
  const [sortedBy, setSortedBy] = useState("price");
  const [order, setOrder] = useState(1);
  const [itemsToShow, setItemsToShow] = useState(5);
  const [isReady, setIsReady] = useState(false);

  useEffect(() => {
    setItemsToShow(5);
  }, [searchTerm]);

  const handleShowMore = () => {
    setItemsToShow((itemsToShow) => itemsToShow + 5);
  };
  const searchQuery = useQuery({
    queryKey: [
      "search",
      { searchedPhrase: searchTerm, sortBy: sortedBy, order: order },
    ],
    enabled: isReady,
    queryFn: async () => {
      if (typeof searchTerm === "string") {
        return fetchProductsSearch(searchTerm, sortedBy, order);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });

  useEffect(() => {
    if (searchTerm !== "") {
      setIsReady(true);
    }
  }, [searchTerm]);
  return (
    <>
      {/* <SortOptions
        sortedBy={sortedBy}
        setSortedBy={setSortedBy}
        order={order}
        setOrder={setOrder}
      /> */}
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
