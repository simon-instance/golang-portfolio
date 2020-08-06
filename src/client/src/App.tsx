import React, { useState } from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { UserForm, Header } from "./Components";

import { ColorModeProvider as MyColorModeProvider } from "./Providers/ColorModeProvider";

const App: React.FC = () => {
  const [navHeight, setNavHeight] = useState<number>(0);

  return (
    <ThemeProvider>
      <ColorModeProvider>
        <CSSReset />
        <Router>
          <MyColorModeProvider>
            <Header setNavHeight={setNavHeight} />
            <Switch>
              <Route exact path="/login">
                <UserForm type="login" height={navHeight} />
              </Route>
              <Route exact path="/register">
                <UserForm type="register" height={navHeight} />
              </Route>
            </Switch>
          </MyColorModeProvider>
        </Router>
      </ColorModeProvider>
    </ThemeProvider>
  );
};

export default App;
