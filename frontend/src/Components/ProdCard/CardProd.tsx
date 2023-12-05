import "./CardProd.css";
import { Link } from "react-router-dom";
type Product = {
  product_url: string;
  product_id: string;
  brand: string;
  model: string;
  imageURL: string;
  sale_price: string;
  display: string;
  processor: string;
  ram: string;
  storage: string;
  battery: string;
};
type ProductPageProps = {
  product: Product;
};
const CardProd: React.FC<ProductPageProps> = ({ product }) => {
  return (
    <Link to={`/product/${product.product_id}`} target="_blank">
      <div className="product">
        <img
          src="https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/4/pr_2023_4_12_13_2_59_532_01.jpg"
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
