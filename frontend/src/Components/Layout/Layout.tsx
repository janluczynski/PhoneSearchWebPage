import Footer from "../Footer/Footer";
import Header from "../Header/Header";

import React, { ReactNode } from "react";

interface LayoutProps {
  children: ReactNode;
}
import "./Layout.css";
const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <Header />
      <main>{children}</main>

      <Footer />
    </>
  );
};
export default Layout;
