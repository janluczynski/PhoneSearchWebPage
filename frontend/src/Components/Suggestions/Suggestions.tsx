import { useQuery } from "@tanstack/react-query";
import { fetchProductsSearch } from "../../API/Api";
import { Product } from "../../types";
import { debounce } from "lodash";
import { useState, useEffect } from "react";
type SuggestionsProps = {
  inputValue: string;
  setSearchTerm: (searchTerm: string) => void;
};
const Suggestions: React.FC<SuggestionsProps> = ({
  inputValue,
  setSearchTerm,
}) => {
  const [debouncedInputValue, setDebouncedInputValue] = useState(inputValue);

  useEffect(() => {
    const debouncedUpdate = debounce(
      () => setDebouncedInputValue(inputValue),
      500,
    );
    debouncedUpdate();
    return debouncedUpdate.cancel;
  }, [inputValue]);

  const searchQuery = useQuery({
    queryKey: ["search", debouncedInputValue],
    enabled: debouncedInputValue !== "" && debouncedInputValue.length > 2,
    queryFn: async () => {
      if (typeof debouncedInputValue === "string") {
        return fetchProductsSearch(debouncedInputValue);
      } else {
        throw new Error(`Search term is undefined`);
      }
    },
  });
  console.log(debouncedInputValue.length);
  return (
    <div>
      {debouncedInputValue.length > 2 &&
        searchQuery &&
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
                  }}
                  className="suggestion"
                >
                  <p>{product.model}</p>
                </div>
              ))}
          </div>
        )}
    </div>
  );
};

export default Suggestions;
