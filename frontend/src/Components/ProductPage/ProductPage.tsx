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
              <div className="productWLinks">
                <div className="productMain">
                  <div className="col40">
                    <img src={product.image} alt={product.model} />
                  </div>
                  <div className="col60">
                    <h2>{product.name}</h2>
                    <div className="productDetails">
                      <ul>
                        <li>Przekątna wyświetlacza: {product.display}</li>
                        <li>Procesor: {product.processor}</li>
                        {product.ram ? (
                          <li>Pamięć RAM: {formatMemory(product.ram)}</li>
                        ) : null}
                        <li>
                          Pamięć wewnętrzna: {formatMemory(product.storage)}
                        </li>
                        {product.battery ? (
                          <li>Pojemność baterii: {product.battery} mAh</li>
                        ) : null}
                        <li></li>
                      </ul>
                    </div>
                  </div>
                </div>
                <div style={{ width: "40%", overflow: "auto", height: "60vh" }}>
                  <ProductOffers
                    product_id={product.product_id}
                    product_name={product.site_name}
                  />
                </div>
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
