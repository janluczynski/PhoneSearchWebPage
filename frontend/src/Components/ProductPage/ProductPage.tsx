import Layout from "../Layout/Layout";
import "./ProductPage.css";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { Button, Link } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import SimilarProducts from "../SimilarProducts/SimilarProducts";
const ProductPage = () => {
  const { product_id } = useParams();

  const productsQuery = useQuery({
    queryKey: ["products", product_id],
    queryFn: async () => {
      const response = await fetch("http://localhost:8080/parse/product", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          product_id: `${product_id}`,
        }),
      });
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data = await response.json();
      return data;
    },
  });
  if (productsQuery.error) {
    return <span>Error: {productsQuery.error.message}</span>;
  }
  if (productsQuery.isLoading) {
    return <span>Loading...</span>;
  }
  const product = productsQuery.data;
  return (
    <Layout>
      <div className="col40">
        <img
          src="https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/4/pr_2023_4_12_13_2_59_532_01.jpg"
          alt={product.model}
        />
      </div>
      <div className="col60">
        <h2>
          {product.brand} {product.model}
        </h2>
        <div className="productDetails">
          <ul>
            <li>Price: {product.sale_price}</li>
            <li>Display: {product.display}</li>
            <li>Processor: {product.processor}</li>
            <li>RAM: {product.ram}</li>
            <li>Storage: {product.storage}</li>
            <li>Battery: {product.battery}</li>
            <li>
              <Link href={product.product_url} target="_blank">
                <Button leftIcon={<ExternalLinkIcon />}>Buy now</Button>
              </Link>
            </li>
          </ul>
        </div>
      </div>
      <div>
        <SimilarProducts name={product.brand} />
      </div>
    </Layout>
  );
};

export default ProductPage;
