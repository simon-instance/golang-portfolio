import React from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { UserForm } from "./Components";

const App: React.FC = () => {
  return (
    <ThemeProvider>
      <ColorModeProvider>
        <CSSReset />
        <Router>
          <Switch>
            <Route exact path="/login">
              <UserForm type="login" />
            </Route>
            <Route exact path="/register">
              <UserForm type="register" />
            </Route>
          </Switch>
        </Router>
      </ColorModeProvider>
    </ThemeProvider>
  );
};

export default App;
