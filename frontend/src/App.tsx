import "./App.css";
import SearchBar from "./Components/Searchbar/Searchbar";
import CardProd from "./Components/ProdCard/CardProd";
import { Product } from "./types";
import Header from "./Components/Header/Header";
import Footer from "./Components/Footer/Footer";
import { useState } from "react";
import SearchedProds from "./Components/SearchedProds/SearchedProds";
import { SearchContext } from "./Contexts/SearchContexts";
import Suggestions from "./Components/Suggestions/Suggestions";
import { Radio } from "@chakra-ui/react";
import SortOptions from "./Components/SortOptions/SortOptions";
import PopularProducts from "./Components/PopularProducts/PopularProducts";
function App() {
  const [searchTerm, setSearchTerm] = useState("");
  const [inputValue, setInputValue] = useState("");
  const [sortedBy, setSortedBy] = useState("price");
  const [order, setOrder] = useState(1);
  const [sortValue, setSortValue] = useState(0);
  return (
    <>
      <SearchContext.Provider
        value={{
          searchTerm,
          setSearchTerm,
          inputValue,
          setInputValue,
          sortedBy,
          setSortedBy,
          order,
          setOrder,
          sortValue,
          setSortValue,
        }}
      >
        <Header />
        <div className="mainPage" style={{ width: "100%" }}>
          <div className="sideBar">
            <div className="sort">
              <SortOptions />
            </div>
          </div>
          <div className="content">
            <div className="fancy">
              <h1>Wybierz najtańszy telefon dla siebie</h1>
            </div>

            <div className="searchContainer">
              <SearchBar width="40vw" height="10vh" fontSize="40px" />
              <div id="portal1">
                <Suggestions portalId="portal1" width="40vw" top="11.2vh" />
              </div>
            </div>
            <div className="searchedprods">
              <SearchedProds searchTerm={searchTerm} />
            </div>

            <div className="waskie">
              <p>
                PhoneCompass to strona, która pomoże Ci znaleźć najlepsze
                telefony w najniższych cenach. Nasza strona oferuje łatwe i
                intuicyjne narzędzia, które pomogą Ci znaleźć telefon, który
                spełni Twoje wymagania. Dzięki naszej pomocy, nie musisz tracić
                czasu na szukanie najlepszych ofert.
              </p>
            </div>

            <h2 className="c">
              <i>Popularne produkty</i>
            </h2>
            <div className="products">
              <PopularProducts />
            </div>
          </div>
        </div>
        <Footer />
      </SearchContext.Provider>
    </>
  );
}

export default App;
