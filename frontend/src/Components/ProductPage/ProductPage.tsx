import Layout from "../Layout/Layout";
import "./ProductPage.css";
import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { Button, Link } from "@chakra-ui/react";
import { ExternalLinkIcon } from "@chakra-ui/icons";
import SimilarProducts from "../SimilarProducts/SimilarProducts";
import { fetchProduct } from "../../API/Api";
import { Product } from "../../types";
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
        {productsQuery.isLoading && <span>Loading...</span>}
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
                      <b>Cena:</b> {product.price} zł
                    </li>
                    <li>
                      <b>Wyświetlacz:</b> {product.display}
                    </li>
                    <li>
                      <b>Procesor:</b> {product.processor}
                    </li>
                    {product.ram ? (
                      <li>
                        <b>RAM:</b>{" "}
                        {product.ram > 1024
                          ? product.ram / 1024 + "GB"
                          : product.ram + "MB"}
                      </li>
                    ) : null}

                    <li>
                      <b>Pamięć:</b>{" "}
                      {product.storage >= 1048576
                        ? product.storage / 1048576 + "TB"
                        : product.storage >= 1024
                          ? product.storage / 1024 + "GB"
                          : product.storage + "MB"}
                    </li>
                    <li>
                      <b>Bateria:</b> {product.battery} mAh
                    </li>
                    <li>
                      <Link href={product.product_url} target="_blank">
                        <Button leftIcon={<ExternalLinkIcon />}>
                          Kup teraz
                        </Button>
                      </Link>
                    </li>
                  </ul>
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
