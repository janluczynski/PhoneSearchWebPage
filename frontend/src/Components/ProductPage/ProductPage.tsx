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
            <div className="col40">
              <img src={product.image} alt={product.model} />
            </div>
            <div className="col60">
              <h2>
                {product.brand} {product.model}
              </h2>
              <div className="productDetails">
                <ul>
                  <li>Cena: {product.price} zł</li>
                  <li>Wyświetlacz: {product.display}</li>
                  <li>Procesor: {product.processor}</li>
                  {product.ram ? (
                    <li>
                      RAM:{" "}
                      {product.ram > 1024
                        ? product.ram / 1024 + "GB"
                        : product.ram + "MB"}
                    </li>
                  ) : null}

                  <li>
                    Pamięć:{" "}
                    {product.storage > 1048576
                      ? product.storage / 1048576 + "TB"
                      : product.storage > 1024
                        ? product.storage / 1024 + "GB"
                        : product.storage + "MB"}
                  </li>
                  <li>Bateria: {product.battery} mAh</li>
                  <li>
                    <Link href={product.product_url} target="_blank">
                      <Button leftIcon={<ExternalLinkIcon />}>Buy now</Button>
                    </Link>
                  </li>
                </ul>
              </div>
            </div>
            <div>
              <SimilarProducts name={product.model} />
            </div>
          </>
        )}
      </div>
    </Layout>
  );
};

export default ProductPage;
