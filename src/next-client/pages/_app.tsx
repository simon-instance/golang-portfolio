import { AppProps } from "next/app";
import cookies from "next-cookies";

import { ThemeProvider, CSSReset, ColorModeProvider } from "@chakra-ui/core";

import { DarkModeProvider } from "../Providers/DarkModeProvider";

import Layout from "../Components/Layout";

interface InitialAppProps extends AppProps {
    initialColorMode: "dark" | "light";
}

const App = ({ Component, pageProps, initialColorMode }: InitialAppProps) => {
    return (
        <ThemeProvider>
            <ColorModeProvider value={initialColorMode}>
                <DarkModeProvider>
                    <CSSReset />
                    <Layout>
                        <Component {...pageProps} />
                    </Layout>
                </DarkModeProvider>
            </ColorModeProvider>
        </ThemeProvider>
    );
};

App.getInitialProps = async ({ Component, ctx }) => {
    let pageProps = {};
    if (Component.getInitialProps) {
        pageProps = await Component.getInitialProps(ctx);
    }
    const { isDarkMode = "false" } = cookies(ctx);
    return {
        pageProps,
        initialColorMode: isDarkMode === "true" ? "dark" : "light",
    };
};

export default App;
