import React, { createContext, useEffect } from "react";

import { useColorMode } from "@chakra-ui/core";

export const DarkModeContext = createContext<Array<any> | undefined>(undefined);

interface Props {
    children: React.ReactNode;
}

export const DarkModeProvider: React.FC<Props> = ({ children }) => {
    const { colorMode, toggleColorMode } = useColorMode();

    const providerValue: Array<any> = [colorMode, toggleColorMode];

    // Sets cookie each time the dark mode is being changed by the user
    useEffect(() => {
        document.cookie = `isDarkMode=${colorMode === "dark"}`;
    }, [colorMode]);

    return (
        <DarkModeContext.Provider value={providerValue}>
            {children}
        </DarkModeContext.Provider>
    );
};
