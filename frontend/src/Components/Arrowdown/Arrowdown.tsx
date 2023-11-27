import { useState, useEffect } from "react";
import "./Arrowdown.css";
import { ArrowDownIcon } from "@chakra-ui/icons";
const Arrowdown = () => {
  const [isVisible, setIsVisible] = useState(true);
  let arrow;
  if (isVisible) {
    arrow = (
      <div className={`arrow-down ${isVisible ? "visible" : "hidden"}`}>
        <ArrowDownIcon />
      </div>
    );
  } else {
    arrow = <></>;
  }
  useEffect(() => {
    const handleScroll = () => {
      const scrollY = window.scrollY || window.pageYOffset;

      if (scrollY > 300) {
        setIsVisible(false);
      } else {
        setIsVisible(true);
      }
    };
    window.addEventListener("scroll", handleScroll);
    return () => {
      window.removeEventListener("scroll", handleScroll);
    };
  }, []);
  return <>{arrow}</>;
};
export default Arrowdown;
