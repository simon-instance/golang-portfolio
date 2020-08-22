import React, { createContext, useEffect } from "react";

import { useColorMode } from "@chakra-ui/core";

export const DarkModeContext = createContext<Array<any> | undefined>(undefined);

export const DarkModeProvider = ({ children }) => {
    const { colorMode, toggleColorMode } = useColorMode();

    const providerValue: Array<any> = [colorMode, toggleColorMode];

    useEffect(() => {
        document.cookie = `isDarkMode=${colorMode === "dark"}`;
    }, [colorMode]);

    return (
        <DarkModeContext.Provider value={providerValue}>
            {children}
        </DarkModeContext.Provider>
    );
};
