import { useState, useEffect } from "react";
import "./Header.css";
import Navbar from "../Navbar/Navbar.js";
import SearchBar from "../Searchbar/Searchbar.js";
import { Collapse } from "@chakra-ui/react";
import logo from "../../Images/logo.png";

const Header = () => {
  const [isVisibleNav, setIsVisibleNav] = useState(false);
  useEffect(() => {
    const handleScroll = () => {
      const scrollY = window.scrollY || window.pageYOffset;

      if (scrollY > 150) {
        setIsVisibleNav(true);
      } else {
        setIsVisibleNav(false);
      }
    };
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);
  return (
    <header>
      <img id="logo" src={logo} />
      <h1>GigaDeals</h1>

      <div className="header-content">
        <Collapse in={isVisibleNav}>
          <div className="HeadNav">
            <SearchBar />
          </div>
        </Collapse>

        <Navbar />
      </div>
    </header>
  );
};

export default Header;
