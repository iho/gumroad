import React from 'react';
import logo from './logo.svg';
import { observer } from "mobx-react-lite";
import './App.css';

import { useQuery } from "./models/reactUtils"

import { Box, Button, Spinner } from "@chakra-ui/core";
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Link
} from "react-router-dom";
import LoginPage from './pages/login';
import SignupPage from './pages/signup';

const About: React.FunctionComponent<{}> = observer(() => {
  return <div>about </div>
});
const Users: React.FunctionComponent<{}> = observer(() => {
  return <div>users</div>
});
const Home: React.FunctionComponent<{}> = observer(() => {
  return <div>Home </div>
});
const App: React.FunctionComponent<{}> = observer(() => {
  return (
    <div style={{ "flex": "1" }}>
      <Switch>
        <Route path="/about">
          <About />
        </Route>
        <Route path="/users">
          <Users />
        </Route>
        <Route path="/login">
          <LoginPage />
        </Route>
        <Route path="/signup">
          <SignupPage />
        </Route>
        <Route path="/">
          <Home />
        </Route>
      </Switch>
    </div>
  )
});

export default App;
