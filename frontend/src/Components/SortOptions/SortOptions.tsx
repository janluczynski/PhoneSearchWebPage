import React from "react";

interface SortOptionsProps {
  onSortChange: (sortOption: string) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({ onSortChange }) => {
  return (
    <div>
      <label>
        <input
          type="radio"
          value="price"
          name="sort"
          onChange={(e) => onSortChange(e.target.value)}
        />
        Price
      </label>
      <label>
        <input
          type="radio"
          value="ram"
          name="sort"
          onChange={(e) => onSortChange(e.target.value)}
        />
        RAM
      </label>
      <label>
        <input
          type="radio"
          value="storage"
          name="sort"
          onChange={(e) => onSortChange(e.target.value)}
        />
        Storage
      </label>
      <label>
        <input
          type="radio"
          value="battery"
          name="sort"
          onChange={(e) => onSortChange(e.target.value)}
        />
        Battery
      </label>
    </div>
  );
};

export default SortOptions;
