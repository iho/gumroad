import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
// import * as serviceWorker from './serviceWorker';

import { RootStore, StoreContext } from "./models"
import { createHttpClient } from "mst-gql"

const rootStore = RootStore.create(undefined, {
    gqlHttpClient: createHttpClient("http://localhost:8080/query")
})

// 4
ReactDOM.render(
    <StoreContext.Provider value={rootStore}>
        <App />
    </StoreContext.Provider>,
    document.getElementById("root")
)

// @ts-ignore
window.store = rootStore

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
// serviceWorker.unregister();
