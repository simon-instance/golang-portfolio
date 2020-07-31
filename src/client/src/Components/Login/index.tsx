import React, { useState } from "react";
import {
  Flex,
  Box,
  IconButton,
  useColorMode,
  Heading,
  FormControl,
  FormLabel,
  Input,
  Button,
} from "@chakra-ui/core";

//
// ThemeSelector: button to toggle dark mode
//

const ThemeSelector: React.FC = () => {
  const { colorMode, toggleColorMode } = useColorMode();

  return (
    <Box textAlign="right" mt={8} mb={4}>
      <IconButton
        aria-label="ToggleDarkMode"
        icon={colorMode === "light" ? "moon" : "sun"}
        onClick={toggleColorMode}
      />
    </Box>
  );
};

//
// LoginHeader: text to inform user about what to do
//

const LoginHeader: React.FC = () => {
  return (
    <Box>
      <Heading as="h3">Log in met uw gegevens</Heading>
    </Box>
  );
};

//
// LoginForm: fields to let the user login
//

const LoginForm: React.FC = () => {
  const [username, setUsername] = useState<string | null>(null);
  const [password, setPassword] = useState<string | null>(null);

  const HandleSubmit: VoidFunction = async () => {
    if (username !== null && password !== null) {
      const requestOptions = {
        method: "POST",
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          username: "username",
          password: "password",
        }),
      };

      const response: any = await fetch(
        "http://127.0.0.1:8080/api/auth/login",
        requestOptions
      );
      const data: any = await response.json();

      console.log(await data);
    }
  };

  return (
    <Box my={8}>
      <FormControl>
        <FormLabel my={1}>Gebruikersnaam</FormLabel>
        <Input
          type="text"
          placeholder="Voer je gebruikersnaam in"
          onInput={(e: any) => setUsername(e.target.value)}
        />
      </FormControl>
      <FormControl mt={4}>
        <FormLabel my={1}>Wachtwoord</FormLabel>
        <Input
          type="password"
          placeholder="Voer je wachtwoord in"
          onInput={(e: any) => setPassword(e.target.value)}
        />
      </FormControl>
      <Button
        width="full"
        variantColor="green"
        mt={6}
        onClick={() => HandleSubmit()}
      >
        Inloggen
      </Button>
    </Box>
  );
};

//
// Login: main component
//

const Login: React.FC = () => {
  return (
    <Flex minHeight="100vh" width="full" align="center" justify="center">
      <Box borderRadius={4} borderWidth={1} px={8}>
        <ThemeSelector />
        <Box>
          <LoginHeader />
          <LoginForm />
        </Box>
      </Box>
    </Flex>
  );
};

export default Login;
