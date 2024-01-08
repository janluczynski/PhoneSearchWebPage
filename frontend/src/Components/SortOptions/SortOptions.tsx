import React from "react";
import { Tabs, TabList, Tab, TabIndicator, Flex } from "@chakra-ui/react";
import { ChevronDownIcon, ChevronUpIcon } from "@chakra-ui/icons";
import { useState } from "react";

interface SortOptionsProps {
  sortedBy: string;
  setSortedBy: (value: string) => void;
  order: number;
  setOrder: (value: number) => void;
}

const SortOptions: React.FC<SortOptionsProps> = ({ setSortedBy, setOrder }) => {
  const [tabIndex, setTabIndex] = useState(0);
  const [sortIndex, setSortIndex] = useState(0);
  const handleSortChange = (index: number) => {
    switch (index) {
      case 0:
        setSortIndex(index);
        setOrder(1);
        break;
      case 1:
        setSortIndex(index);
        setOrder(-1);
        break;
      default:
        break;
    }
  };
  const handleTabsChange = (index: number) => {
    switch (index) {
      case 0:
        setTabIndex(index);
        setSortedBy("price");
        break;
      case 1:
        setTabIndex(index);
        setSortedBy("ram");
        break;
      case 2:
        setTabIndex(index);
        setSortedBy("battery");
        break;
      case 3:
        setTabIndex(index);
        setSortedBy("storage");
        break;
      default:
        break;
    }
  };
  return (
    <div>
      <Flex>
        <Tabs
          defaultIndex={0}
          index={tabIndex}
          onChange={handleTabsChange}
          position="relative"
          variant="unstyled"
          background="#ab7b7b"
        >
          <TabList>
            <Tab _selected={{ color: "white", bg: "#860000" }} height="5vh">
              Cena
            </Tab>
            <Tab _selected={{ color: "white", bg: "#860000" }} height="5vh">
              RAM
            </Tab>
            <Tab _selected={{ color: "white", bg: "#860000" }} height="5vh">
              Bateria
            </Tab>
            <Tab _selected={{ color: "white", bg: "#860000" }} height="5vh">
              Pamięć
            </Tab>
          </TabList>
          <TabIndicator
            mt="-1.5px"
            height="2px"
            bg="#860000"
            borderRadius="1px"
          />
        </Tabs>
        <Tabs
          index={sortIndex}
          onChange={handleSortChange}
          position="relative"
          variant="unstyled"
          background="#ab7b7b"
        >
          <TabList>
            <Tab _selected={{ color: "white", bg: "#860000" }}>
              {<ChevronUpIcon height="3.3vh" />}
            </Tab>
            <Tab _selected={{ color: "white", bg: "#860000" }}>
              {<ChevronDownIcon height="3.3vh" />}
            </Tab>
          </TabList>
          <TabIndicator
            mt="-1.5px"
            height="2px"
            bg="#860000"
            borderRadius="1px"
          />
        </Tabs>
      </Flex>
    </div>
  );
};

export default SortOptions;
