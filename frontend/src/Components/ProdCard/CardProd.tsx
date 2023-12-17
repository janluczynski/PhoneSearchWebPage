import "./CardProd.css";
import { Link } from "react-router-dom";
import { Product } from "../../types";
type ProductPageProps = {
  product: Product;
};
const CardProd: React.FC<ProductPageProps> = ({ product }) => {
  return (
    <Link to={`/product/${product.product_id}`} target="_blank">
      <div className="product">
        <img
          src={product.image}
          alt={product.model}
          style={{ maxWidth: "100%", height: "20vh" }}
        />
        <div className="text">
          <h3 className="product-model">
            {product.brand} {product.model}
          </h3>
          <h3>{product.price} zł</h3>
        </div>
      </div>
    </Link>
  );
};
export default CardProd;
