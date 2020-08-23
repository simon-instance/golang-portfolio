import React from "react";
import { Box } from "@chakra-ui/core";

interface WrapperProps {
    children: React.ReactNode;
    type?: "small" | "large";
}

const Wrapper: React.FC<WrapperProps> = ({ children, type = "small" }) => {
    return (
        <Box
            my="auto"
            mx="auto"
            w="100%"
            h={type === "small" ? "50vh" : "80vh"}
            maxW={type === "small" ? "400px" : "800px"}
        >
            {children}
        </Box>
    );
};

export default Wrapper;
