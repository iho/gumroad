import React from "react";
import { Heading } from "@chakra-ui/core";
import { observer } from "mobx-react-lite";
import { useForm } from "react-hook-form";
import { useState } from "react";
import { useQuery, RootStore, StoreContext } from "../models";
import { useContext } from "react";
import {
  SimpleGrid,
  Flex,
  InputGroup,
  InputLeftAddon,
  Input,
  InputRightElement,
  Icon,
  Button
} from "@chakra-ui/core";

type FormData = {
  email: string;
  password: string;
  username: string;
  name: string;
};
const SignupPage = observer(() => {
  const { handleSubmit, errors, register, setError } = useForm<FormData>();
  const { setQuery, loading, error, data } = useQuery();
  console.log(data);
  return (
    <SimpleGrid
      columns={2}
      spacingX={1}
      spacingY={1}
      display="flex"
      width="full"
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
        <Heading> Реєстрація </Heading>

        <form
          onSubmit={handleSubmit((values: FormData) => {
            setQuery(store => store.mutateSignup(values, undefined));
          })}
        >
          <InputGroup mt={2} width="full">
            <InputLeftAddon width={1 / 4}>Email</InputLeftAddon>
            <Input name="email" ref={register} isRequired={true} />
            <InputRightElement>
              <Icon name="email" />
            </InputRightElement>
          </InputGroup>
          <InputGroup mt={2} width="full">
            <InputLeftAddon width={1 / 4}>Логін</InputLeftAddon>
            <Input name="username" ref={register} isRequired={true} />
            <InputRightElement>
              <Icon name="view" />
            </InputRightElement>
          </InputGroup>
          <InputGroup mt={2} width="full">
            <InputLeftAddon width={1 / 4}>Ім`я</InputLeftAddon>
            <Input name="name" ref={register} isRequired={true} />
            <InputRightElement>
              <Icon name="edit" />
            </InputRightElement>
          </InputGroup>
          <InputGroup mt={2} width="full">
            <InputLeftAddon width={1 / 4}>Пароль</InputLeftAddon>
            <Input
              name="password"
              ref={register}
              type="password"
              isRequired={true}
            />
            <InputRightElement>
              <Icon name="lock" />
            </InputRightElement>
          </InputGroup>
          <Button mt={3} width="full" type="submit">
            Зареєструватися
          </Button>
        </form>
      </Flex>
    </SimpleGrid>
  );
});

export default SignupPage;
