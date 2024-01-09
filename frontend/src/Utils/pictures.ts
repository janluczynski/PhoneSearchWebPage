import mediaMarktLogo from "../Images/mediamarkt.png";
import komputronikLogo from "../Images/komputronik.png";
import xkomLogo from "../Images/x-kom.jpg";

export const getPicture= (picture: string) => {
    if (picture==="mediamarkt.pl") {
      return mediaMarktLogo;
    }
    else if (picture==="www.x-kom.pl") {
      return xkomLogo;
    }
    else if (picture==="www.komputronik.pl") {
      return komputronikLogo;
    }
  }