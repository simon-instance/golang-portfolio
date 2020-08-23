import React, { useState, useEffect, useRef, useContext } from "react";

import { Link as CLink, Box, Button, IconButton } from "@chakra-ui/core";

import { DarkModeContext } from "../Providers/DarkModeProvider";

import { AiFillGithub, AiFillLinkedin } from "react-icons/ai";

const ThemeSelector: React.FC = () => {
    const [darkMode, toggleDarkMode] = useContext(DarkModeContext);

    return (
        <Box>
            <IconButton
                aria-label="ToggleDarkMode"
                onClick={() => toggleDarkMode()}
                icon={darkMode === "light" ? "moon" : "sun"}
            />
        </Box>
    );
};

const SocialMediaButtons: React.FC = () => {
    const links = [
        {
            href: "https://github.com/scrummer123",
            icon: AiFillGithub,
            label: "My github",
        },
        {
            href: "https://www.linkedin.com/in/simon-peters-0834151a1",
            icon: AiFillLinkedin,
            label: "My Linkedin",
        },
    ] as const;

    return (
        <div>
            {links.map(({ href, icon, label }, i) => (
                <CLink key={i} href={href} isExternal ml="3">
                    <IconButton
                        variant="ghost"
                        aria-label={label}
                        icon={icon}
                    />
                </CLink>
            ))}
        </div>
    );
};

const Nav: React.FC<{ readonly first: boolean }> = ({ first }) => {
    const links = [
        {
            to: "/register",
            text: "Registreren",
        },
        {
            to: "/login",
            text: "Inloggen",
        },
        {
            to: "test",
            text: "test",
        },
    ] as const;

    const variant = first === true ? ("solid" as const) : ("ghost" as const);
    const d: Array<string> =
        first === true ? ["none", "block"] : ["block", "none"];

    return (
        <Box d="flex">
            {links.map(({ to, text }, i) => (
                <Button
                    key={i}
                    d={i < 2 ? d : first === true ? "none" : "block"}
                    variant={variant}
                    ml="3"
                >
                    {text}
                </Button>
            ))}
        </Box>
    );
};

const Header: React.FC<{ setHeaderHeight }> = ({ setHeaderHeight }) => {
    const headerWrapper = useRef<HTMLDivElement>(null);
    const [colorMode, setColorMode] = useContext(DarkModeContext);

    // scroll vars
    let prev = 0;
    const [py, setPy] = useState<number>(4);

    //
    // function for sweet nav animation
    //

    const handleNav: (e: any) => void = (e) => {
        const window = e.currentTarget;

        if (prev > window.scrollY) {
            setPy(3);
        } else if (prev < window.scrollY) {
            setPy(1);
        }
        prev = window.scrollY;
    };

    //
    // useEffect hook for handling the mobile nav
    //

    useEffect(() => {
        prev = window.scrollY;
        window.addEventListener("scroll", (e) => handleNav(e));
        const height: number = document.getElementById("header-root")
            .clientHeight;
        setHeaderHeight(height);
    });

    //
    // render variables
    //

    // [xs, sm, md, (lg)]...
    const sideFlex = [0, 2, 1, 3];
    const mainFlex = [1, 32, 15, 7];

    const currentColor = colorMode === "light" ? "gray.200" : "gray.700";
    const currentNavClass = colorMode === "light" ? "nav-light" : "nav-dark";

    const ButtonContainer: React.ReactNode = (
        <Box d="flex" w={["100%", "auto"]} justifyContent="space-between">
            <ThemeSelector />
            <SocialMediaButtons />
        </Box>
    );

    return (
        <Box w="100%" position="fixed" id="header-root">
            <Box
                d="flex"
                py={py + 2}
                className={"nav-default " + currentNavClass}
            >
                <Box flex={sideFlex} />
                <Box id="button-container-box" mx="5" flex={mainFlex}>
                    {ButtonContainer}
                    <Nav first={true} />
                </Box>
                <Box flex={sideFlex} />
            </Box>
            <Box
                className="nav-default"
                shadow="md"
                d="flex"
                py={py}
                bg={currentColor}
            >
                <Box flex={sideFlex} />
                <Box flex={mainFlex}>
                    <Nav first={false} />
                </Box>
                <Box flex={sideFlex} />
            </Box>
        </Box>
    );
};

export default Header;
