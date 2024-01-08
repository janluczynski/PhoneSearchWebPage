import { useState, useEffect } from "react";
import "./Header.css";
import Navbar from "../Navbar/Navbar.js";
import SearchBar from "../Searchbar/Searchbar.js";
import { Collapse } from "@chakra-ui/react";
import logo from "../../Images/logo.png";
import { debounce } from "lodash";
import Suggestions from "../Suggestions/Suggestions.js";

const Header: React.FC = () => {
  const [isVisibleNav, setIsVisibleNav] = useState(false);
  useEffect(() => {
    const handleScroll = debounce(() => {
      const scrollY = window.scrollY || document.documentElement.scrollTop;
      if (scrollY > 150) {
        setIsVisibleNav(true);
      } else {
        setIsVisibleNav(false);
      }
    }, 400);
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);
  return (
    <header>
      <div className="headerContent headerH">
        <img id="logo" src={logo} />
        <h1>PhoneCompass</h1>
      </div>
      <div className="headerContent">
        <Collapse
          in={isVisibleNav && location.pathname === "/"}
          style={{ overflow: "visible" }}
        >
          <div className="searchContainer">
            <SearchBar width="20vw" height="4vh" fontSize="16px" />
            <div id="portal2">
              <Suggestions portalId="portal2" width="20vw" top="5.13vh" />
            </div>
          </div>
        </Collapse>

        <Navbar />
      </div>
    </header>
  );
};

export default Header;
