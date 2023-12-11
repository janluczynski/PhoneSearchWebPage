import { useState, useEffect } from "react";
import "./Header.css";
import Navbar from "../Navbar/Navbar.js";
import SearchBar from "../Searchbar/Searchbar.js";
import { Collapse } from "@chakra-ui/react";
import logo from "../../Images/logo.png";
import { debounce } from "lodash";
type SearchBarProps = {
  setSearchTerm: (searchTerm: string) => void;
};
const Header: React.FC<SearchBarProps> = ({ setSearchTerm }) => {
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
      <div className="headerContent">
        <img id="logo" src={logo} />
        <h1>GigaDeals</h1>
      </div>
      <div className="headerContent">
        <Collapse in={isVisibleNav && location.pathname === "/"}>
          <SearchBar setSearchTerm={setSearchTerm} />
        </Collapse>

        <Navbar />
      </div>
    </header>
  );
};

export default Header;
