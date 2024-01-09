import React from "react";
import "./ProductOffers.css";
import { Link } from "@chakra-ui/react";
import { getPicture } from "../../Utils/pictures";
const ProductOffers: React.FC = () => {
  return (
    <div className="productOffers">
      <Link href="https://www.mediaexpert.pl/" target="blank">
        <div className="offer">
          <img src={getPicture("mediaexpert.pl")} alt="Komputronik" />
          <p>Cena: 3000zł</p>
        </div>
      </Link>
      <Link href="https://www.mediamarkt.pl/" target="blank">
        <div className="offer">
          <img src={getPicture("mediamarkt.pl")} alt="Mediamartk" />
          <p>Cena: 4000zł</p>
        </div>
      </Link>
      <Link href="https://www.komputronik.pl/" target="blank">
        <div className="offer">
          <img src={getPicture("komputronik.pl")} alt="Mediamartk" />
          <p>Cena: 3300zł</p>
        </div>
      </Link>
    </div>
  );
};

export default ProductOffers;
