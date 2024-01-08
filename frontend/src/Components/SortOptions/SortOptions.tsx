import React from "react";
import {
  Radio,
  Stack,
  RadioGroup,
  Tabs,
  TabList,
  Tab,
  TabIndicator,
  TabPanel,
  TabPanels,
} from "@chakra-ui/react";
import { ArrowDownIcon, ArrowUpIcon } from "@chakra-ui/icons";
import { useState } from "react";
interface SortOptionsProps {
  sortedBy: string;
  setSortedBy: (value: string) => void;
  order: number;
  setOrder: (value: number) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({
  sortedBy,
  setSortedBy,
  order,
  setOrder,
}) => {
  const [tabIndex, setTabIndex] = useState(-1);
  const handleTabsChange = (index: number) => {
    switch (index) {
      case 0:
        setTabIndex(index);
        setSortedBy("price");
        setOrder(1);
        break;
      case 1:
        setTabIndex(index);
        setSortedBy("price");
        setOrder(-1);
        break;

      case 2:
        setTabIndex(index);
        setSortedBy("ram");
        setOrder(1);
        break;
      case 3:
        setTabIndex(index);
        setSortedBy("battery");
        setOrder(1);
        break;
      case 4:
        setTabIndex(index);
        setSortedBy("storage");
        setOrder(1);
        break;
      default:
        break;
    }
  };
  return (
    <div>
      <Tabs
        index={tabIndex}
        onChange={handleTabsChange}
        position="relative"
        variant="unstyled"
        background="#ab7b7b"
      >
        <TabList>
          <Tab _selected={{ color: "white", bg: "#860000" }}>
            Cena {<ArrowUpIcon />}
          </Tab>
          <Tab _selected={{ color: "white", bg: "#860000" }}>
            Cena {<ArrowDownIcon />}
          </Tab>
          <Tab _selected={{ color: "white", bg: "#860000" }}>RAM</Tab>
          <Tab _selected={{ color: "white", bg: "#860000" }}>Bateria</Tab>
          <Tab _selected={{ color: "white", bg: "#860000" }}>Pamięć</Tab>
        </TabList>
        <TabIndicator
          mt="-1.5px"
          height="2px"
          bg="#860000"
          borderRadius="1px"
        />
      </Tabs>
    </div>
  );
};

export default SortOptions;
