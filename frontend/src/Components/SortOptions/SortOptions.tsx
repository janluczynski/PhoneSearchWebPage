import React from "react";
import { Radio, Stack, RadioGroup } from "@chakra-ui/react";
import { ArrowDownIcon, ArrowUpIcon } from "@chakra-ui/icons";
interface SortOptionsProps {
  onSortChange: (sortOption: string) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({ onSortChange }) => {
  return (
    <div>
      <RadioGroup>
        <Stack direction="row">
          <Radio
            value="priceAsc"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Cena ({<ArrowUpIcon />})
          </Radio>
          {/* <Radio
            value="priceDesc"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Cena ({<ArrowDownIcon />})
          </Radio> */}
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
            Bateria
          </Radio>
          <Radio
            value="storage"
            name="sort"
            onChange={(e) => onSortChange(e.target.value)}
          >
            Pamięć
          </Radio>
        </Stack>
      </RadioGroup>
    </div>
  );
};

export default SortOptions;
