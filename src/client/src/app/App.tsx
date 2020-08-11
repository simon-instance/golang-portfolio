import React, { useState } from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { UserForm, Header } from "./Components";

import { ColorModeProvider as MyColorModeProvider } from "./Providers/ColorModeProvider";

const App: React.FC = () => {
  return (
    <ThemeProvider>
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
    </ThemeProvider>
  );
};

export default App;
