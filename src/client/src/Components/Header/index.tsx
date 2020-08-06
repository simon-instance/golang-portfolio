import React, { useEffect, useRef, useContext } from "react";

import { ColorMode } from "../../Providers/ColorModeProvider";

import { Box, Button, IconButton } from "@chakra-ui/core";
import { Link } from "react-router-dom";

//
// ThemeSelector: button to toggle dark mode
//

const ThemeSelector: React.FC = () => {
  const [colorMode, toggleColorMode] = useContext(ColorMode);

  return (
    <Box>
      <IconButton
        aria-label="ToggleDarkMode"
        icon={colorMode === "light" ? "moon" : "sun"}
        onClick={toggleColorMode}
      />
    </Box>
  );
};

const Nav: React.FC = () => (
  <div>
    <Link to="/register">
      <Button>Register</Button>
    </Link>
    <Link to="/login">
      <Button>Login</Button>
    </Link>
  </div>
);

const Header: React.FC<{ setNavHeight: Function }> = ({ setNavHeight }) => {
  const headerWrapper = useRef<HTMLDivElement>(null);
  const [colorMode] = useContext(ColorMode);

  //
  // useEffect hook for getting navbar height
  //

  useEffect(() => {
    if (headerWrapper.current && headerWrapper.current.clientHeight) {
      setNavHeight(headerWrapper.current.clientHeight);
    }
  }, [headerWrapper]);

  //
  // render variables
  //
  // [xs, sm, md, (lg)]...
  const sideFlex = [0, 1, 3];
  const mainFlex = [1, 3, 6];

  return (
    <Box
      d="flex"
      ref={headerWrapper}
      py="5"
      borderBottom="1px"
      borderColor={colorMode === "light" ? "gray.200" : "gray.700"}
    >
      <Box flex={sideFlex} />
      <Box
        d="flex"
        justifyContent="space-between"
        alignItems="center"
        mx="5"
        flex={mainFlex}
      >
        <ThemeSelector />
        <Nav />
      </Box>
      <Box flex={sideFlex} />
    </Box>
  );
};

export default Header;
