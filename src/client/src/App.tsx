import React from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";

const App: React.FC = ({ children }) => {
  return (
    <ThemeProvider>
      <ColorModeProvider>
        <CSSReset />
        {children}
      </ColorModeProvider>
    </ThemeProvider>
  );
};

export default App;
