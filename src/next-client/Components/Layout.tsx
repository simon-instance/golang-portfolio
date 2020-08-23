import React, { useState } from "react";

import { Box } from "@chakra-ui/core";

import Header from "./Header";

const Layout: React.FC<{ children: React.ReactNode }> = ({ children }) => {
    const [headerHeight, setHeaderHeight] = useState(0);

    return (
        <div>
            <Header setHeaderHeight={setHeaderHeight} />
            <Box height={headerHeight + "px"}></Box>
            {children}
        </div>
    );
};

export default Layout;
