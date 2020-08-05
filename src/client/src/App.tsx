import React, { useState, useEffect } from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import { UserForm, Header } from "./Components";

const App: React.FC = () => {
  const [navHeight, setNavHeight] = useState<number>(0);

  return (
    <ThemeProvider>
      <ColorModeProvider>
        <CSSReset />
        <Router>
          <Header setNavHeight={setNavHeight} />
          <Switch>
            <Route exact path="/login">
              <UserForm type="login" height={navHeight} />
            </Route>
            <Route exact path="/register">
              <UserForm type="register" height={navHeight} />
            </Route>
          </Switch>
        </Router>
      </ColorModeProvider>
    </ThemeProvider>
  );
};

export default App;
