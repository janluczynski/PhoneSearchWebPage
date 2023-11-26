import "./CardProd.css";
interface ProductProps{
name:string;
price: number;
imgSrc: string;
link: string;
};

const CardProd: React.FC<ProductProps> = ({name, price, imgSrc, link}) =>{
    return(
        <div className="product">
      <img src={imgSrc} alt={name} style={{ maxWidth: '100%', height: 'auto' }} />
      <div className="text">
      <h3>{name}</h3>
      <p>{price.toFixed(2)}zł</p>
      <p>
        <a href={link} target="_blank" rel="noopener noreferrer">
          View Product
        </a>
      </p>
      <button>Add to Cart</button>
      </div>
    </div>
    )
}
export default CardProd;