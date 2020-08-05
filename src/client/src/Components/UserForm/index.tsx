import React, { useState } from "react";
import {
  TryCatchHandler,
  TryCatchInterface,
  TryCatchDataInterface,
} from "../../Handlers";
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
  useToast,
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

const LoginHeader: React.FC<{ type: string }> = ({ type }) => {
  return (
    <Box>
      <Heading as="h3">
        {type === "login" ? "Log in" : "Registreer"} met uw gegevens
      </Heading>
    </Box>
  );
};

//
// LoginForm: fields to let the user login
//

const LoginForm: React.FC<{ type: string }> = ({ type }) => {
  const [username, setUsername] = useState<string | null>(null);
  const [password, setPassword] = useState<string | null>(null);

  const TCHandler: TryCatchInterface = new TryCatchHandler();
  const toast: any = useToast();

  const HandleSubmit: VoidFunction = async () => {
    if (username !== null && password !== null) {
      const requestOptions = {
        method: "POST",
        body: JSON.stringify({
          Username: username,
          Password: password,
        }),
      };

      let data: TryCatchDataInterface = TCHandler.Data;
      let toastConfig: object = {
        title: null,
        description: null,
      };

      try {
        const response: any = await fetch(
          `http://127.0.0.1:8080/api/auth/${
            type === "login" ? type : "register"
          }`,
          requestOptions
        );

        data.response.data = await response.json();
        data.response.status = response.status;

        const error: Error | void = TCHandler.handleData(data);
        if (typeof error !== "undefined") throw error;

        toastConfig = {
          title: "Gelukt",
          description: `Je bent nu ${
            type === "login" ? "ingelogd" : "geregistreed"
          }`,
        };
      } catch (e) {
        data.error.status = true;
        data.error.message = e;
        toastConfig = {
          title: "Foutmelding",
          description: e.message,
        };
      } finally {
        console.log(data.error.status);
        TCHandler.Data = data;
        toast({
          position: "top-left",
          status: data.error.status === true ? "error" : "success",
          duration: 7000,
          ...toastConfig,
          isClosable: true,
        });
      }
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

const UserForm: React.FC<{ type: string }> = ({ type }) => {
  return (
    <Flex minHeight="100vh" width="full" align="center" justify="center">
      <Box borderWidth={1} px={8} mx={4}>
        <ThemeSelector />
        <Box>
          <LoginHeader type={type} /> <LoginForm type={type} />
        </Box>
      </Box>
    </Flex>
  );
};

export default UserForm;
