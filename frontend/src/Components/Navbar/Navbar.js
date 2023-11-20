import React from 'react';
import { HamburgerIcon } from '@chakra-ui/icons';
import "./Navbar.css";
import {
    Menu,
    MenuButton,
    MenuList,
    MenuItem,
    IconButton,
  } from '@chakra-ui/react'

const Navbar = () => {

    return(
        <div class="navbar">
        <Menu>
  <MenuButton
    as={IconButton}
    ariaLabel='Options'
    icon={<HamburgerIcon color="black" height="3vh" width="3vh"/>}
    variant='outline'
    height="4vh"
    width="4vh"
    borderRadius= "5px"
  />
  <MenuList color="black">
    <MenuItem height="3.5vh">
      New Tab
    </MenuItem>
    <MenuItem height="3.5vh">
      New Window
    </MenuItem>
    <MenuItem height="3.5vh">
      Open Closed Tab
    </MenuItem>
    <MenuItem height="3.5vh">
      Open File...
    </MenuItem>
  </MenuList>
</Menu>
        	</div> 
    )
}

export default Navbar;