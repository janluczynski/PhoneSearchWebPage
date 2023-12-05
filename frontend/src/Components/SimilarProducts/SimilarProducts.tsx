import { useQuery } from "@tanstack/react-query";
import CardProd from "../ProdCard/CardProd";
import { useState } from "react";
type SimilarProductsProps = {
  name: string;
};

const SimilarProducts: React.FC<SimilarProductsProps> = ({ name }) => {
  const [itemsToShow, setItemsToShow] = useState(5);

  const handleShowMore = () => {
    setItemsToShow(itemsToShow + 5);
  };

  const similarProdQuery = useQuery({
    queryKey: ["productsByBrand", name], // using the brand as a key

    queryFn: async () => {
      const response = await fetch("http://localhost:8080/search", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          SearchedPhrase: `${name}`, // send the brand in the request body
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data;
    },
    enabled: name !== "", // only run the query if the brand is not an empty string
  });
  if (similarProdQuery.error) {
    return <span>Error: {similarProdQuery.error.message}</span>;
  }
  if (similarProdQuery.isLoading) {
    return <span>Loading...</span>;
  }
  console.log(similarProdQuery.data);
  return (
    <>
      <h2 className="c">Similar products</h2>
      <div className="similarProd">
        {similarProdQuery.data &&
          similarProdQuery.data
            .slice(0, itemsToShow)
            .map((product: any) => (
              <CardProd key={product.product_id} product={product} />
            ))}
      </div>
      <button className="c" onClick={handleShowMore}>
        Pokaż więcej
      </button>
    </>
  );
};
export default SimilarProducts;
