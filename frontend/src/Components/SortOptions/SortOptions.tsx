import React from "react";
import { Radio, Stack, RadioGroup } from "@chakra-ui/react";
interface SortOptionsProps {
  onSortChange: (sortOption: string) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({ onSortChange }) => {
  return (
    <div>
      <RadioGroup>
        <Stack direction="row">
          <Radio
            value="price"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Price
          </Radio>
          <Radio
            value="ram"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            RAM
          </Radio>
          <Radio
            value="battery"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Battery
          </Radio>
          <Radio
            value="storage"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Storage
          </Radio>
        </Stack>
      </RadioGroup>
    </div>
  );
};

export default SortOptions;
