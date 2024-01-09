import React from "react";
import "./ProductOffers.css";
import { Link } from "@chakra-ui/react";
import mediaExpertLogo from "../../Images/mediaexpert.jpg";
import mediaMarktLogo from "../../Images/mediamarkt.png";
import komputronikLogo from "../../Images/komputronik.png";
const ProductOffers: React.FC = () => {
  return (
    <div className="productOffers">
      <Link href="https://www.mediaexpert.pl/" target="blank">
        <div className="offer">
          <img src={mediaExpertLogo} alt="Komputronik" />
          <p>Cena: 3000zł</p>
        </div>
      </Link>
      <Link href="https://www.mediamarkt.pl/" target="blank">
        <div className="offer">
          <img src={mediaMarktLogo} alt="Mediamartk" />
          <p>Cena: 4000zł</p>
        </div>
      </Link>
      <Link href="https://www.komputronik.pl/" target="blank">
        <div className="offer">
          <img src={komputronikLogo} alt="Mediamartk" />
          <p>Cena: 3300zł</p>
        </div>
      </Link>
    </div>
  );
};

export default ProductOffers;
