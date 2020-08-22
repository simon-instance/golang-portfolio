import React, { lazy } from "react";

import { CSSReset, ThemeProvider, ColorModeProvider } from "@chakra-ui/core";
import Router from "./Router";

const UserForm = lazy(() =>
    import(/* webpackChunkName: "user_form" */ "./Components/UserForm")
);

const App: React.FC = () => {
    return (
        <ThemeProvider>
            <ColorModeProvider>
                <CSSReset />
                <Router>
                    {{
                        routeTo: "/login",
                        component: <UserForm type="login" />,
                    }}
                    {{
                        routeTo: "/register",
                        component: <UserForm type="register" />,
                    }}
                </Router>
            </ColorModeProvider>
        </ThemeProvider>
    );
};

export default App;
