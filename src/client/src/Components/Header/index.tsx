import React, { useEffect, useRef } from "react";

import { Box, Grid, IconButton, useColorMode } from "@chakra-ui/core";

//
// ThemeSelector: button to toggle dark mode
//

const ThemeSelector: React.FC = () => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <Box textAlign="right" mt={8} mb={4}>
      <IconButton
        aria-label="ToggleDarkMode"
        icon={colorMode === "light" ? "moon" : "sun"}
        onClick={toggleColorMode}
      />
    </Box>
  );
};

const Header: React.FC<{ setNavHeight: Function }> = ({ setNavHeight }) => {
  const headerWrapper = useRef<HTMLDivElement>(null);

  //
  // useEffect hook for getting navbar height
  //

  useEffect(() => {
    if (headerWrapper.current && headerWrapper.current.clientHeight) {
      setNavHeight(headerWrapper.current.clientHeight);
    }
    //setNavHeight(headerWrapper.current);
  }, [headerWrapper]);

  //
  // render variables
  //
  // [xs, sm, md, (lg)]...
  const sideFlex = [0, 1, 2];
  const mainFlex = [1, 3, 4];

  return (
    <Box d="flex" ref={headerWrapper}>
      <Box flex={sideFlex}>
        <h1>test</h1>
      </Box>
      {/*<ThemeSelector />*/}
      <Box w="60%" flex={mainFlex}>
        <h1>test</h1>
      </Box>
      <Box w="20%" flex={sideFlex}>
        <h1>test</h1>
      </Box>
    </Box>
  );
};

export default Header;
