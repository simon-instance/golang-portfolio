import React, { Suspense } from "react";

import { BrowserRouter, Switch, Route } from "react-router-dom";

import { ColorModeProvider } from "./Providers/ColorModeProvider";

import Header from "./Components/Header";

interface Props {
    children: Array<Child>;
}

interface Child {
    routeTo: string;
    component: JSX.Element;
}

const Router: React.FC<Props> = ({ children }) => (
    <BrowserRouter>
        <ColorModeProvider>
            <Header />
            <Suspense fallback={<div>loading...</div>}>
                <Switch>
                    {children.map(({ routeTo, component }) => (
                        <Route exact path={routeTo} render={() => component} />
                    ))}
                </Switch>
            </Suspense>
        </ColorModeProvider>
    </BrowserRouter>
);

export default Router;
