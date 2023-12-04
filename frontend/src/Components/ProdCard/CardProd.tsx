import "./CardProd.css";
import { useNavigate } from "react-router-dom";
import { Link } from "react-router-dom";
type Product = {
  product_url: string;
  product_id: string;
  brand: string;
  model: string;
  imageURL: string;
  price: string;
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
    <div className="product">
      <img
        src="https://cdn.x-kom.pl/i/setup/images/prod/big/product-new-big,,2023/4/pr_2023_4_12_13_2_59_532_01.jpg"
        alt={product.model}
        style={{ maxWidth: "100%", height: "20vh" }}
      />
      <div className="text">
        <h3>{product.model}</h3>
        <Link to={`/product/${product.product_id}`} target="_blank">Check details</Link>
      </div>
    </div>
  );
};
export default CardProd;
