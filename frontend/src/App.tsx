import React from 'react';
import logo from './logo.svg';
import { observer } from "mobx-react-lite";
import './App.css';

import { useQuery } from "./models/reactUtils"

// import gql from 'graphql-tag';
import {Box, Button, Spinner} from "@chakra-ui/core";
const App: React.FunctionComponent<{}> = observer(() => {
  const {  error, loading, data } = useQuery(store => store.queryProducts({after: 25}))
  if (error) return <Box>{error.message}</Box>
  if (loading) return <Spinner />


  return (
    <ul>
      {data?.products.map(product => (
        <Box key={product.id}>{product.name}</Box>
      ))}
    </ul>
  )
});

export default App;
