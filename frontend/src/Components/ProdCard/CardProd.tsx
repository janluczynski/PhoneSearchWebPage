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
          <h3 className="productModel">
            {product.brand} {product.model}
          </h3>
          <h3>
            <b>{product.price} zł</b>
          </h3>
          <h3>
            {""}
            {product.ram === 0
              ? ""
              : product.ram > 1024
                ? product.ram / 1024 + "GB/"
                : product.ram + "MB/"}
            {product.storage === 0
              ? ""
              : product.storage >= 1048576
                ? product.storage / 1048576 + "TB"
                : product.storage >= 1024
                  ? product.storage / 1024 + "GB"
                  : product.storage + "MB"}
          </h3>
        </div>
      </div>
    </Link>
  );
};
export default CardProd;
