import "./CardProd.css";
import { Link } from "react-router-dom";
import { Product } from "../../types";
import { formatMemory } from "../../Utils/converters";
import { useQuery } from "@tanstack/react-query";
import { increaseProductViews } from "../../API/Api";
import { useMutation } from "@tanstack/react-query";

type ProductPageProps = {
  product: Product;
};

const CardProd: React.FC<ProductPageProps> = ({ product }) => {
  const mutation = useMutation({
    mutationFn: () => increaseProductViews(product.product_id),
  });
  return (
    <Link
      to={`/product/${product.product_id}`}
      target="_blank"
      onClick={() => mutation.mutate()}
    >
      <div className="product">
        <img
          src={product.image}
          alt={product.model}
          style={{ maxWidth: "100%", height: "20vh" }}
        />
        <div className="text">
          <h3 className="productModel">
            {product.brand} {product.model}
          </h3>
          <h3>
            <b>{product.price} zł</b>
          </h3>
          <h3>
            {""}
            {product.ram === 0 ? "" : formatMemory(product.ram) + "/"}
            {product.storage === 0 ? "" : formatMemory(product.storage)}
          </h3>
        </div>
      </div>
    </Link>
  );
};
export default CardProd;
