import React from "react";

import Header from "./Header";

interface Props {
    children: Array<JSX.Element> | JSX.Element;
}

const Layout: React.FC<Props> = ({ children }) => (
    <div>
        <Header />
        {children}
    </div>
);

export default Layout;
