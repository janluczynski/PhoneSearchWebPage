import mediaExpertLogo from "../Images/mediaexpert.jpg";
import mediaMarktLogo from "../Images/mediamarkt.png";
import komputronikLogo from "../Images/komputronik.png";

export const getPicture= (picture: string) => {
    if (picture==="mediamarkt.pl") {
      return mediaMarktLogo;
    }
    else if (picture==="mediaexpert.pl") {
      return mediaExpertLogo;
    }
    else if (picture==="komputronik.pl") {
      return komputronikLogo;
    }
  }