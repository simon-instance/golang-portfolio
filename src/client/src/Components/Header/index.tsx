import React, { useState, useEffect, useRef, useContext } from "react";

import { ColorMode } from "../../Providers/ColorModeProvider";

import { Link as CLink, Box, Button, IconButton } from "@chakra-ui/core";
import { Link } from "react-router-dom";
import { AiFillGithub, AiFillLinkedin } from "react-icons/ai";

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
  ];

  return (
    <div>
      {links.map(({ href, icon, label }) => (
        <CLink href={href} isExternal ml="3">
          <IconButton variant="ghost" aria-label={label} icon={icon} />
        </CLink>
      ))}
    </div>
  );
};

const Nav: React.FC<{ first: boolean }> = ({ first }) => {
  const links: Array<{ to: string; text: string }> = [
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
  ];

  const variant = first === true ? "solid" : "ghost";
  const d = first === true ? ["none", "block"] : ["block", "none"];

  return (
    <Box d="flex">
      {links.map(({ to, text }, i) => (
        <Link to={to}>
          <Button
            d={i < 2 ? d : first === true ? "none" : "block"}
            variant={variant}
            ml="3"
          >
            {text}
          </Button>
        </Link>
      ))}
    </Box>
  );
};

const Header: React.FC = () => {
  const headerWrapper = useRef<HTMLDivElement>(null);
  const [colorMode] = useContext(ColorMode);

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
  });

  //
  // render variables
  //

  // [xs, sm, md, (lg)]...
  const sideFlex = [0, 2, 1, 3];
  const mainFlex = [1, 32, 15, 7];

  const currentColor = colorMode === "light" ? "gray.200" : "gray.700";
  const currentNavBg =
    colorMode === "light" ? "rgba(255,255,255,.8)" : "rgba(26,32,44, .8)";

  const ButtonContainer: React.ReactNode = (
    <Box d="flex" w={["100%", "auto"]} justifyContent="space-between">
      <ThemeSelector />
      <SocialMediaButtons />
    </Box>
  );

  return (
    <Box w="100%" position="fixed">
      <Box
        d="flex"
        transition=".25s ease all"
        py={py + 2}
        borderBottom="1px"
        borderColor={currentColor}
        className={"filter"}
        bg={currentNavBg}
        style={{ backdropFilter: "blur(8px)" }}
      >
        <Box flex={sideFlex} />
        <Box
          d="flex"
          justifyContent="space-between"
          alignItems="center"
          mx="5"
          flex={mainFlex}
        >
          {ButtonContainer}
          <Nav first={true} />
        </Box>
        <Box flex={sideFlex} />
      </Box>
      <Box
        transition=".25s ease all"
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
