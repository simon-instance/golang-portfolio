import React, { useState } from "react";

import { CSSReset, ChakraProvider, ColorModeProvider } from "@chakra-ui/core";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import theme from "@chakra-ui/theme";

import { UserForm, Header } from "./Components";

import { ColorModeProvider as MyColorModeProvider } from "./Providers/ColorModeProvider";

const App: React.FC = () => {
    return (
        <ChakraProvider theme={theme}>
            <ColorModeProvider>
                <CSSReset />
                <Router>
                    <MyColorModeProvider>
                        <Header />
                        <Switch>
                            <Route exact path="/login">
                                <UserForm type="login" />
                            </Route>
                            <Route exact path="/register">
                                <UserForm type="register" />
                            </Route>
                        </Switch>
                    </MyColorModeProvider>
                </Router>
            </ColorModeProvider>
        </ChakraProvider>
    );
};

export default App;
