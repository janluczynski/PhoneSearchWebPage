import React from "react";
import "./Header.css";
import Navbar from "../Navbar/Navbar.js";
import SearchBar from "../Searchbar/Searchbar.js";

const Header = () => {
    return(
        <header>
        <h1>SKLEP</h1>
        <div className="header-content">
            <SearchBar />
            <Navbar />
        </div>
        </header>
    )
    
}

export default Header;