import React, { useEffect, createContext } from "react";
import { useColorMode } from "@chakra-ui/core";

export const ColorMode = createContext<Array<any>>([]);

export const ColorModeProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <ColorMode.Provider value={[colorMode, toggleColorMode]}>
      {children}
    </ColorMode.Provider>
  );
};
