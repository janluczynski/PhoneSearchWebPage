import React from "react";
import "./ProductOffers.css";
import { Link } from "@chakra-ui/react";
import { getPicture } from "../../Utils/pictures";
import { useQuery } from "@tanstack/react-query";
import { fetchSameProducts } from "../../API/Api";
import { Spinner } from "@chakra-ui/react";
type ProductOffersProps = {
  product_id: string;
};
const ProductOffers: React.FC<ProductOffersProps> = ({ product_id }) => {
  const sameProdQuery = useQuery({
    queryKey: ["sameProductsByID", product_id],
    enabled: product_id !== "",
    queryFn: () => {
      if (typeof product_id === "string") {
        return fetchSameProducts(product_id);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });
  if (sameProdQuery.error) {
    return <span>Error: {sameProdQuery.error.message}</span>;
  }
  if (sameProdQuery.isLoading) {
    return (
      <span>
        <Spinner color="#860000" />
      </span>
    );
  }
  return (
    <div className="productOffers">
      {sameProdQuery.data &&
        Object.values(sameProdQuery.data).map(
          (sameProduct: any, index: number) => (
            <Link href={sameProduct[0]} target="blank" key={index}>
              <div className="offer">
                <img src={getPicture(sameProduct[1])} />
                <p>{sameProduct[2]} zł</p>
              </div>
            </Link>
          ),
        )}
    </div>
  );
};

export default ProductOffers;
