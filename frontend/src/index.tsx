import React from "react";
import ReactDOM from "react-dom";
import App from "./App";

import { theme, ThemeProvider, CSSReset, Flex } from "@chakra-ui/core";
import { RootStore, StoreContext } from "./models";
import { createHttpClient } from "mst-gql";
import Header from "./components/header";
import { Router } from "react-router-dom";
import Footer from "./components/footer";
import { createBrowserHistory } from "history";
const rootStore = RootStore.create(undefined, {
  gqlHttpClient: createHttpClient("http://localhost:8080/query")
});

const breakpoints: any = ["360px", "768px", "1024px", "1440px"];
breakpoints.sm = breakpoints[0];
breakpoints.md = breakpoints[1];
breakpoints.lg = breakpoints[2];
breakpoints.xl = breakpoints[3];

const newTheme = {
  ...theme,
  breakpoints
};
const history = createBrowserHistory();

ReactDOM.render(
  <ThemeProvider theme={newTheme}>
    <StoreContext.Provider value={rootStore}>
      <CSSReset />
      <Router history={history}>
        <Flex flexDirection="column" minHeight="100vh">
          <Header />
          <App />
          <Footer />
        </Flex>
      </Router>
    </StoreContext.Provider>
  </ThemeProvider>,
  document.getElementById("root")
);

// @ts-ignore
window.store = rootStore;
