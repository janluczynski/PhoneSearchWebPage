import React, { useEffect, useState } from "react";
import { Product } from "../../types";
import CardProd from "../ProdCard/CardProd";
import { useQuery } from "@tanstack/react-query";
import { fetchProductsSearch } from "../../API/Api";
import { Button } from "@chakra-ui/react";
import { AddIcon } from "@chakra-ui/icons";
import SortOptions from "../SortOptions/SortOptions";
type SearchedProdsProps = {
  searchTerm: string;
};

const SearchedProds: React.FC<SearchedProdsProps> = ({ searchTerm }) => {
  const [sortOption, setSortOption] = useState("");
  const [itemsToShow, setItemsToShow] = useState(5);

  useEffect(() => {
    setItemsToShow(5);
  }, [searchTerm]);

  const handleShowMore = () => {
    setItemsToShow((itemsToShow) => itemsToShow + 5);
  };
  const searchQuery = useQuery({
    queryKey: ["search", searchTerm],
    enabled: searchTerm !== "",
    queryFn: async () => {
      if (typeof searchTerm === "string") {
        return fetchProductsSearch(searchTerm);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });
  return (
    <>
      <SortOptions onSortChange={setSortOption} />
      {searchQuery.isLoading && <span>Loading...</span>}
      {searchQuery.error && <span>Error: {searchQuery.error.message}</span>}
      {searchQuery.data && (
        <div className="similarProd">
          {searchQuery.data
            .sort((a: Product, b: Product) => {
              switch (sortOption) {
                case "price":
                  return a.price - b.price;
                case "ram":
                  return a.ram - b.ram;
                case "storage":
                  return a.storage - b.storage;
                case "battery":
                  return a.battery - b.battery;
                default:
                  return 0;
              }
            })
            .slice(0, itemsToShow)
            .map((product: Product) => (
              <CardProd key={product.product_id} product={product} />
            ))}
        </div>
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
