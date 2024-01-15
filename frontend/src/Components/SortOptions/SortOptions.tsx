import React from "react";
import { ChevronDownIcon, ChevronUpIcon } from "@chakra-ui/icons";
import { useState } from "react";
import {
  Radio,
  RadioGroup,
  Stack,
  RangeSlider,
  RangeSliderTrack,
  RangeSliderFilledTrack,
  RangeSliderThumb,
  Input,
  Divider,
  Flex,
  Box,
} from "@chakra-ui/react";

interface SortOptionsProps {
  sortedBy: string;
  setSortedBy: (value: string) => void;
  order: number;
  setOrder: (value: number) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({ setSortedBy, setOrder }) => {
  const [tabIndex, setTabIndex] = useState(0);
  const [sortIndex, setSortIndex] = useState(0);
  return (
    <div>
      <h2>Sortuj po:</h2>
      <RadioGroup defaultValue="1" marginTop="7%">
        <Stack>
          <Radio value="1">Cena rosnąco</Radio>
          <Radio value="2">Cena malejąco</Radio>
        </Stack>
      </RadioGroup>
      <Divider my={4} borderWidth="2px" />
      <h2>Filtry</h2>
      <p>Cena</p>
      {/* <RangeSlider defaultValue={[120, 240]} min={0} max={300} step={30}>
        <RangeSliderTrack bg="red.100">
          <RangeSliderFilledTrack bg="tomato" />
        </RangeSliderTrack>
        <RangeSliderThumb boxSize={6} index={0} />
        <RangeSliderThumb boxSize={6} index={1} />
      </RangeSlider> */}
      <Flex>
        <Input
          width="30%"
          placeholder="od"
          boxShadow="0px 4px 4px rgba(0, 0, 0, 0.1)"
          focusBorderColor="#860000"
        />
        <Box
          border="1px solid"
          borderColor="gray.200"
          width="1px"
          alignSelf="center"
          mx={2}
          paddingLeft="20%"
        />
        <Input
          width="30%"
          placeholder="do"
          boxShadow="0px 4px 4px rgba(0, 0, 0, 0.1)"
          focusBorderColor="#860000"
        />
      </Flex>
      <p>Pamięc</p>
      <RadioGroup marginTop="7%">
        <Stack>
          <Radio value="1">128 GB</Radio>
          <Radio value="2">256 GB</Radio>
          <Radio value="3">512 GB</Radio>
          <Radio value="4">1 TB</Radio>
          <p>Ram</p>
          <Radio value="5">4 GB</Radio>
          <Radio value="6">6 GB</Radio>
          <Radio value="7">8 GB</Radio>
          <Radio value="8">12 GB</Radio>
          <Radio value="9">16 GB</Radio>
        </Stack>
      </RadioGroup>
    </div>
  );
};

export default SortOptions;
