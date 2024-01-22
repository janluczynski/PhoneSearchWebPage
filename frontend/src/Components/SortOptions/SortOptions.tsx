import React from "react";
import { ChevronDownIcon, ChevronUpIcon } from "@chakra-ui/icons";
import { SearchContext } from "../../Contexts/SearchContexts";
import { useContext } from "react";
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
import { NotAllowedIcon } from "@chakra-ui/icons";
const SortOptions = () => {
  const { setOrder, setSortValue, setSortedBy } = useContext(SearchContext);
  return (
    <div>
      <h2>Sortuj po:</h2>
      <RadioGroup defaultValue="1" marginTop="7%">
        <Stack>
          <Radio value="1" onChange={() => setOrder(1)}>
            Cena rosnąco
          </Radio>
          <Radio value="2" onChange={() => setOrder(-1)}>
            Cena malejąco
          </Radio>
        </Stack>
      </RadioGroup>
      <Divider my={4} borderWidth="2px" />
      <h2>Filtry</h2>
      {/* <p>Cena</p> */}
      {/* <RangeSlider defaultValue={[120, 240]} min={0} max={300} step={30}>
        <RangeSliderTrack bg="red.100">
          <RangeSliderFilledTrack bg="tomato" />
        </RangeSliderTrack>
        <RangeSliderThumb boxSize={6} index={0} />
        <RangeSliderThumb boxSize={6} index={1} />
      </RangeSlider> */}
      {/* <Flex>
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
      </Flex> */}

      <RadioGroup marginTop="7%" defaultValue="0">
        <Radio
          value="0"
          onChange={() => {
            setSortedBy("price");
            setSortValue(0);
          }}
        >
          <NotAllowedIcon />
        </Radio>
        <p>Pamięc</p>
        <Stack>
          <Radio
            value="1"
            onChange={() => {
              setSortedBy("storage");
              setSortValue(128 * 1024);
            }}
          >
            128 GB
          </Radio>
          <Radio
            value="2"
            onChange={() => {
              setSortedBy("storage");
              setSortValue(256 * 1024);
            }}
          >
            256 GB
          </Radio>
          <Radio
            value="3"
            onChange={() => {
              setSortedBy("storage");
              setSortValue(512 * 1024);
            }}
          >
            512 GB
          </Radio>
          <Radio
            value="4"
            onChange={() => {
              setSortedBy("storage");
              setSortValue(1024 * 1024);
            }}
          >
            1 TB
          </Radio>
          <p>Ram</p>
          <Radio
            value="5"
            onChange={() => {
              setSortedBy("ram");
              setSortValue(4 * 1024);
            }}
          >
            4 GB
          </Radio>
          <Radio
            value="6"
            onChange={() => {
              setSortedBy("ram");
              setSortValue(6 * 1024);
            }}
          >
            6 GB
          </Radio>
          <Radio
            value="7"
            onChange={() => {
              setSortedBy("ram");
              setSortValue(8 * 1024);
            }}
          >
            8 GB
          </Radio>
          <Radio
            value="8"
            onChange={() => {
              setSortedBy("ram");
              setSortValue(12 * 1024);
            }}
          >
            12 GB
          </Radio>
          <Radio
            value="9"
            onChange={() => {
              setSortedBy("ram");
              setSortValue(16 * 1024);
            }}
          >
            16 GB
          </Radio>
        </Stack>
      </RadioGroup>
    </div>
  );
};

export default SortOptions;
