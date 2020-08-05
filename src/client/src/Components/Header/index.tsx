import React, { useEffect, useRef } from "react";

import { Box } from "@chakra-ui/core";

const Header: React.FC<{ setNavHeight: Function }> = ({ setNavHeight }) => {
  const headerWrapper = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (headerWrapper.current && headerWrapper.current.clientHeight) {
      setNavHeight(headerWrapper.current.clientHeight);
    }
    //setNavHeight(headerWrapper.current);
  }, [headerWrapper]);

  return (
    <Box d="flex" ref={headerWrapper}>
      <h1>test</h1>
    </Box>
  );
};

export default Header;
