import Layout from "../Layout/Layout";
import "./ProductPage.css";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { Button, Link, Spinner } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import SimilarProducts from "../SimilarProducts/SimilarProducts";
import { fetchProduct } from "../../API/Api";
import { Product } from "../../types";
import { formatMemory } from "../../Utils/converters";
import ProductOffers from "../ProductOffers/ProductOffers";
const ProductPage = () => {
  const { product_id } = useParams();

  const productsQuery = useQuery({
    queryKey: ["products", product_id],
    queryFn: () => {
      if (typeof product_id === "string") {
        return fetchProduct(product_id);
      } else {
        throw new Error(`Product ID is undefined`);
      }
    },
  });
  const product: Product = productsQuery.data;
  return (
    <Layout>
      <div>
        {productsQuery.isLoading && <span>{<Spinner color="#860000" />}</span>}
        {productsQuery.error && (
          <span>Error: {productsQuery.error.message}</span>
        )}
        {!productsQuery.isLoading && !productsQuery.error && (
          <>
            <div className="productContainer">
              <div className="col40">
                <img src={product.image} alt={product.model} />
              </div>
              <div className="col60">
                <h2>
                  <b>{product.name}</b>
                </h2>
                <div className="productDetails">
                  <ul>
                    <li>
                      <b>Wyświetlacz:</b> {product.display}
                    </li>
                    <li>
                      <b>Procesor:</b> {product.processor}
                    </li>
                    {product.ram ? (
                      <li>
                        <b>RAM:</b> {formatMemory(product.ram)}
                      </li>
                    ) : null}
                    <li>
                      <b>Pamięć:</b> {formatMemory(product.storage)}
                    </li>
                    {product.battery ? (
                      <li>
                        <b>Bateria:</b> {product.battery} mAh
                      </li>
                    ) : null}
                    <li></li>
                  </ul>
                </div>
              </div>
              <div>
                <ProductOffers
                  product_id={product.product_id}
                  product_name={product.name}
                />
              </div>
              <div>
                <SimilarProducts
                  name={product.model}
                  ram={product.ram}
                  storage={product.storage}
                />
              </div>
            </div>
          </>
        )}
      </div>
    </Layout>
  );
};

export default ProductPage;
