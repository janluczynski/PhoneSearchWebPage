import "./CardProd.css";
import { useNavigate } from "react-router-dom";
interface ProductProps {
  name: string;
  price: number;
  imgSrc: string;
  link: string;
  id: string;
}

const CardProd: React.FC<ProductProps> = ({ name, price, imgSrc, link, id }) => {
const navigate = useNavigate();

  return (
    <div className="product">
        <img
        src={imgSrc}
        alt={name}
        style={{ maxWidth: "100%", height: "20vh" }}
      />
      <div className="text">
        <h3>{name}</h3>
        <button onClick={() => navigate(`/CardProd/${id}`)}>ccccc</button>
      </div>
    </div>
  );
};
export default CardProd;
