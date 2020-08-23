import React from "react";
import { Box } from "@chakra-ui/core";

interface WrapperProps {
    children: React.ReactNode;
    type?: "small" | "large";
}

const Wrapper: React.FC<WrapperProps> = ({ children, type = "small" }) => {
    return (
        <Box
            maxW={type === "small" ? "400px" : "800px"}
            minH={type === "small" ? "50vh" : "80vh"}
            mx="auto"
            w="100%"
            d="flex"
            alignItems="center"
        >
            <Box w="100%">{children}</Box>
        </Box>
    );
};

export default Wrapper;
