import { fetchTopProducts } from "../../API/Api";
import { useQuery } from "@tanstack/react-query";
import { Product } from "../../types";
import CardProd from "../ProdCard/CardProd";
import { Spinner } from "@chakra-ui/react";

const PopularProducts = () => {
  const PopularProductsQuery = useQuery({
    queryKey: ["topProducts"],
    queryFn: () => fetchTopProducts(),
  });
  return (
    <>
      {PopularProductsQuery.isLoading && (
        <span>
          <Spinner color="#860000" />
        </span>
      )}
      {PopularProductsQuery.error && (
        <span>Error: {PopularProductsQuery.error.message}</span>
      )}
      {PopularProductsQuery.data && (
        <div className="similarProd">
          {PopularProductsQuery.data.map((product: Product) => (
            <CardProd key={product.product_id} product={product} />
          ))}
        </div>
      )}
    </>
  );
};
export default PopularProducts;
