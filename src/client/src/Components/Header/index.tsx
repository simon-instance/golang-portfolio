import React, { useEffect, useRef, useContext } from "react";

import { ColorMode } from "../../Providers/ColorModeProvider";

import { Link as CLink, Box, Button, IconButton } from "@chakra-ui/core";
import { Link } from "react-router-dom";
import { AiFillGithub, AiFillLinkedin } from "react-icons/ai";

//
// ThemeSelector: button to toggle dark mode
//

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
    },
    {
      href: "https://www.linkedin.com/in/simon-peters-0834151a1",
      icon: AiFillLinkedin,
    },
  ];

  return (
    <div>
      {links.map(({ href, icon }) => (
        <CLink href={href} isExternal>
          <IconButton aria-label="My Github profile" icon={icon} />
        </CLink>
      ))}
    </div>
  );
};

const Nav: React.FC = () => {
  const links = [
    {
      to: "/register",
      text: "Registreren",
    },
    {
      to: "login",
      text: "Inloggen",
    },
  ];

  return (
    <div>
      {links.map(({ to, text }) => (
        <Link to={to}>
          <Button ml="3">{text}</Button>
        </Link>
      ))}
    </div>
  );
};

const Header: React.FC<{ setNavHeight: Function }> = ({ setNavHeight }) => {
  const headerWrapper = useRef<HTMLDivElement>(null);
  const [colorMode] = useContext(ColorMode);

  //
  // useEffect hook for getting navbar height
  //

  useEffect(() => {
    if (headerWrapper.current && headerWrapper.current.clientHeight) {
      setNavHeight(headerWrapper.current.clientHeight);
    }
  }, [headerWrapper]);

  //
  // render variables
  //
  // [xs, sm, md, (lg)]...
  const sideFlex = [0, 1, 3];
  const mainFlex = [1, 3, 6];

  return (
    <Box
      d="flex"
      ref={headerWrapper}
      py="5"
      borderBottom="1px"
      borderColor={colorMode === "light" ? "gray.200" : "gray.700"}
    >
      <Box flex={sideFlex} />
      <Box
        d="flex"
        justifyContent="space-between"
        alignItems="center"
        mx="5"
        flex={mainFlex}
      >
        <ThemeSelector />
        <Nav />
      </Box>
      <Box flex={sideFlex} />
    </Box>
  );
};

export default Header;
