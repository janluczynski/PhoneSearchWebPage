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
        <Button colorScheme="facebook" size="sm">
          Login
        </Button>
        <Button colorScheme="facebook" size="sm">
          Register
        </Button>
        <Menu>
          <MenuButton
            as={IconButton}
            colorScheme="facebook"
            icon={<HamburgerIcon />}
          >
            Menu
          </MenuButton>
          <MenuList backgroundColor="black">
            <MenuGroup title="Profile" backgroundColor="black">
              <MenuItem>My Account</MenuItem>
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
