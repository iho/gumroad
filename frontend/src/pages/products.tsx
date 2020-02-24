import React from 'react';
import { observer } from "mobx-react-lite";
import './App.css';

import { useQuery } from "../models/reactUtils"

import {Box, Button, Spinner} from "@chakra-ui/core";

const Products: React.FunctionComponent<{}> = observer(() => {
    const {  error, loading, data } = useQuery(store => store.queryProducts({}))
    if (error) return <Box>{error.message}</Box>
    if (loading) return <Spinner />
    return (
      <ul style={{"flex": "1"}}>
        {data?.products.map(product => (
          <Box key={product.id}>{product.name}</Box>
        ))}
      </ul>
    )
  });
  
  export default Products;
  