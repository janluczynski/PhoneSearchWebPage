import { useQuery } from "@tanstack/react-query";
import CardProd from "../ProdCard/CardProd";
import { useState } from "react";
import { Button, Spinner } from "@chakra-ui/react";
import { ChevronDownIcon, AddIcon } from "@chakra-ui/icons";
import { fetchSimilarProducts } from "../../API/Api";
import { Product } from "../../types";
type SimilarProductsProps = {
  brand: string;
  popularity: number;
};

const SimilarProducts: React.FC<SimilarProductsProps> = ({
  brand,
  popularity,
}) => {
  const [similarItemsToShow, setSimilarItemsToShow] = useState(5);

  const handleShowMore = () => {
    setSimilarItemsToShow((similarItemsToShow) => similarItemsToShow + 5);
  };

  const similarProdQuery = useQuery({
    queryKey: ["productsByBrand", name],
    enabled: brand !== "",
    queryFn: () => {
      if (typeof name === "string") {
        return fetchSimilarProducts(brand, popularity);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });
  if (similarProdQuery.error) {
    return <span>Error: {similarProdQuery.error.message}</span>;
  }
  if (similarProdQuery.isLoading) {
    return (
      <span>
        <Spinner color="#860000" />
      </span>
    );
  }

  return (
    <>
      <h2 className="c" style={{ marginTop: "50px" }}>
        Podobne Produkty
      </h2>
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
