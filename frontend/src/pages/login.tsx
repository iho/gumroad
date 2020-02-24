import React from "react";
import {
  SimpleGrid,
  Flex,
  FormControl,
  FormLabel,
  Input,
  FormHelperText,
  FormErrorMessage,
  Button
} from "@chakra-ui/core";
import { Heading } from "@chakra-ui/core";
import { observer } from "mobx-react-lite";
import { useForm } from "react-hook-form";
import { useState } from "react";
import { useQuery, RootStore, StoreContext } from "../models";
import { useContext } from "react";
type FormData = {
  email: string;
  password: string;
};
const LoginPage = observer(() => {
  const { handleSubmit, errors, register, setError } = useForm<FormData>();
  const { setQuery, loading, error, data } = useQuery();
  console.log(data);
  return (
    <SimpleGrid
      columns={3}
      spacingX={1}
      spacingY={1}
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
    >
      <Flex
        display="flex"
        flexDirection="column"
        alignItems="flex-start"
        justifyContent="flex-start"
      >
        <Heading> Вхід </Heading>
        <form
          onSubmit={handleSubmit((values: any) => {
            setQuery(store => store.login(values));
          })}
        >
          <FormControl isInvalid={errors.email !== undefined || error}>
            <FormLabel>Email</FormLabel>
            <Input
              size="md"
              as="input"
              variant="outline"
              isFullWidth
              focusBorderColor="blue.500"
              errorBorderColor="red.500"
              name="email"
              ref={register}
              isRequired={true}
            />
            {/* <FormHelperText>Helper message</FormHelperText> */}
            <FormErrorMessage>
              {error && "Користувач з такими данними відсутній"}
            </FormErrorMessage>
          </FormControl>
          <FormControl isInvalid={errors.password !== undefined}>
            <FormLabel>Пароль</FormLabel>
            <Input
              size="md"
              as="input"
              variant="outline"
              isFullWidth
              focusBorderColor="blue.500"
              errorBorderColor="red.500"
              name="password"
              type="password"
              ref={register}
              isRequired={true}
            />
            {/* <FormHelperText>Helper message</FormHelperText> */}
            <FormErrorMessage>
              {errors.password && errors.password.message}
            </FormErrorMessage>
          </FormControl>
          {/* <br/> */}
          <Button variantColor="teal" isLoading={loading} type="submit" mt={2}>
            Увійти
          </Button>
        </form>
      </Flex>
    </SimpleGrid>
  );
});

export default LoginPage;
