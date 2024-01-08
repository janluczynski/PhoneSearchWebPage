import { useQuery } from "@tanstack/react-query";
import CardProd from "../ProdCard/CardProd";
import { useState } from "react";
import { Button } from "@chakra-ui/react";
import { ChevronDownIcon, AddIcon } from "@chakra-ui/icons";
import { fetchSimilarProducts } from "../../API/Api";
import { Product } from "../../types";
type SimilarProductsProps = {
  name: string;
  ram: number;
  storage: number;
};

const SimilarProducts: React.FC<SimilarProductsProps> = ({
  name,
  ram,
  storage,
}) => {
  const [similarItemsToShow, setSimilarItemsToShow] = useState(5);

  const handleShowMore = () => {
    setSimilarItemsToShow((similarItemsToShow) => similarItemsToShow + 5);
  };

  const similarProdQuery = useQuery({
    queryKey: ["productsByBrand", name],
    enabled: name !== "",
    queryFn: () => {
      if (typeof name === "string") {
        return fetchSimilarProducts(name, ram, storage);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });
  if (similarProdQuery.error) {
    return <span>Error: {similarProdQuery.error.message}</span>;
  }
  if (similarProdQuery.isLoading) {
    return <span>Loading...</span>;
  }

  return (
    <>
      <h2 className="c">Podobne Produkty</h2>
      <div className="similarProd">
        {similarProdQuery.data &&
          similarProdQuery.data
            .slice(0, similarItemsToShow)
            .map((product: Product) => (
              <CardProd key={product.product_id} product={product} />
            ))}
      </div>
      {similarProdQuery.data &&
        similarProdQuery.data.length > similarItemsToShow && (
          <Button className="c" onClick={handleShowMore} leftIcon={<AddIcon />}>
            Pokaż więcej
          </Button>
        )}
    </>
  );
};
export default SimilarProducts;
