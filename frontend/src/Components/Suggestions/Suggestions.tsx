import { useQuery } from "@tanstack/react-query";
import { fetchSuggestions } from "../../API/Api";
import { Product } from "../../types";
import { debounce } from "lodash";
import { useState, useEffect } from "react";
import { SearchContext } from "../../Contexts/SearchContexts";
import "./Suggestions.css";
import { useContext } from "react";
import ReactDOM from "react-dom";

interface SuggestionsProps {
  portalId: string;
}

const Suggestions: React.FC<SuggestionsProps> = ({ portalId }) => {
  const { setSearchTerm, inputValue, setInputValue } =
    useContext(SearchContext);
  const [debouncedInputValue, setDebouncedInputValue] = useState(inputValue);
  const debouncedUpdate = debounce(
    () => setDebouncedInputValue(inputValue),
    500,
  );
  useEffect(() => {
    debouncedUpdate();
    return debouncedUpdate.cancel;
  }, [inputValue]);

  const searchQuery = useQuery({
    queryKey: ["search", debouncedInputValue],
    enabled: debouncedInputValue.length > 2,
    queryFn: async () => {
      return fetchSuggestions(debouncedInputValue);
    },
  });
  const portalElement = document.getElementById(portalId);
  if (portalElement === null) {
    return null;
  }
  return ReactDOM.createPortal(
    <div>
      {debouncedInputValue.length > 2 &&
        searchQuery.data &&
        searchQuery.data.length > 0 && (
          <div className="dropdown">
            {searchQuery.data
              .filter(
                (product: Product, index: number, self: Product[]) =>
                  index === self.findIndex((p) => p.model === product.model),
              )
              .slice(0, 5)
              .map((product: Product) => (
                <div
                  key={product.product_id}
                  onClick={() => {
                    setSearchTerm(product.model);
                    setInputValue("");
                    window.scrollTo(0, 0);
                  }}
                  className="suggestion"
                >
                  <p>{product.model}</p>
                </div>
              ))}
          </div>
        )}
    </div>,
    portalElement,
  );
};

export default Suggestions;
