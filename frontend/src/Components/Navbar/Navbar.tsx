import "./Navbar.css";
import {
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Button,
  MenuGroup,
  MenuDivider,
  Stack,
  IconButton,
} from "@chakra-ui/react";
import { HamburgerIcon } from "@chakra-ui/icons";
const Navbar = () => {
  return (
    <div className="navbar">
      <Stack spacing={3} direction="row" align="center">
        <Menu>
          <MenuButton as={IconButton} icon={<HamburgerIcon />}>
            Menu
          </MenuButton>
          <MenuList>
            <MenuGroup title="Profile">
              <MenuItem>Log in</MenuItem>
              <MenuItem>Register</MenuItem>
              <MenuItem>Payments </MenuItem>
            </MenuGroup>
            <MenuDivider />
            <MenuGroup title="Help">
              <MenuItem>Docs</MenuItem>
              <MenuItem>FAQ</MenuItem>
            </MenuGroup>
          </MenuList>
        </Menu>
      </Stack>
    </div>
  );
};

export default Navbar;
