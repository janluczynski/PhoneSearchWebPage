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
          src="https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2022/7/pr_2022_7_4_13_46_24_503_05.jpg"
          alt={product.model}
          style={{ maxWidth: "100%", height: "20vh" }}
        />
        <div className="text">
          <h3 className="product-model">{product.model}</h3>
          <h3>{product.sale_price}</h3>
        </div>
      </div>
    </Link>
  );
};
export default CardProd;
