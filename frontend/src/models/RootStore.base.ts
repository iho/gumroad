/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */
import { ObservableMap } from "mobx"
import { types } from "mobx-state-tree"
import { MSTGQLStore, configureStoreMixin, QueryOptions, withTypedRefs } from "mst-gql"

import { ProductModel, ProductModelType } from "./ProductModel"
import { productModelPrimitives, ProductModelSelector } from "./ProductModel.base"
import { UserModel, UserModelType } from "./UserModel"
import { userModelPrimitives, UserModelSelector } from "./UserModel.base"
import { ExtendedUserModel, ExtendedUserModelType } from "./ExtendedUserModel"
import { extendedUserModelPrimitives, ExtendedUserModelSelector } from "./ExtendedUserModel.base"
import { PayResponseModel, PayResponseModelType } from "./PayResponseModel"
import { payResponseModelPrimitives, PayResponseModelSelector } from "./PayResponseModel.base"
import { TokenResponseModel, TokenResponseModelType } from "./TokenResponseModel"
import { tokenResponseModelPrimitives, TokenResponseModelSelector } from "./TokenResponseModel.base"
import { BoolResponseModel, BoolResponseModelType } from "./BoolResponseModel"
import { boolResponseModelPrimitives, BoolResponseModelSelector } from "./BoolResponseModel.base"


export type BuyProduct = {
  productId: string
}
export type NewProduct = {
  name: string
  price: number
  description: string
  summary: string
  callToAction: string
  coverImage: string
  slug: string
  isPablished: boolean
  receipt: string
  content: string
}
export type PublishProduct = {
  productId: string
  slug: string
}
/* The TypeScript type that explicits the refs to other models in order to prevent a circular refs issue */
type Refs = {
  products: ObservableMap<string, ProductModelType>,
  users: ObservableMap<string, UserModelType>,
  extendedusers: ObservableMap<string, ExtendedUserModelType>
}

/**
* Store, managing, among others, all the objects received through graphQL
*/
export const RootStoreBase = withTypedRefs<Refs>()(MSTGQLStore
  .named("RootStore")
  .extend(configureStoreMixin([['Product', () => ProductModel], ['User', () => UserModel], ['ExtendedUser', () => ExtendedUserModel], ['PayResponse', () => PayResponseModel], ['TokenResponse', () => TokenResponseModel], ['BoolResponse', () => BoolResponseModel]], ['Product', 'User', 'ExtendedUser']))
  .props({
    products: types.optional(types.map(types.late((): any => ProductModel)), {}),
    users: types.optional(types.map(types.late((): any => UserModel)), {}),
    extendedusers: types.optional(types.map(types.late((): any => ExtendedUserModel)), {})
  })
  .actions(self => ({
    queryProduct(variables: { username: string, slug: string }, resultSelector: string | ((qb: ProductModelSelector) => ProductModelSelector) = productModelPrimitives.toString(), options: QueryOptions = {}) {
      return self.query<{ product: ProductModelType}>(`query product($username: String!, $slug: String!) { product(username: $username, slug: $slug) {
        ${typeof resultSelector === "function" ? resultSelector(new ProductModelSelector()).toString() : resultSelector}
      } }`, variables, options)
    },
    queryProducts(variables: { username?: string, count?: number, after?: number }, resultSelector: string | ((qb: ProductModelSelector) => ProductModelSelector) = productModelPrimitives.toString(), options: QueryOptions = {}) {
      return self.query<{ products: ProductModelType[]}>(`query products($username: String, $count: Int, $after: Int) { products(username: $username, count: $count, after: $after) {
        ${typeof resultSelector === "function" ? resultSelector(new ProductModelSelector()).toString() : resultSelector}
      } }`, variables, options)
    },
    queryMyProducts(variables: { count?: number, after?: number }, resultSelector: string | ((qb: ProductModelSelector) => ProductModelSelector) = productModelPrimitives.toString(), options: QueryOptions = {}) {
      return self.query<{ myProducts: ProductModelType[]}>(`query myProducts($count: Int, $after: Int) { myProducts(count: $count, after: $after) {
        ${typeof resultSelector === "function" ? resultSelector(new ProductModelSelector()).toString() : resultSelector}
      } }`, variables, options)
    },
    queryMe(variables?: {  }, resultSelector: string | ((qb: ExtendedUserModelSelector) => ExtendedUserModelSelector) = extendedUserModelPrimitives.toString(), options: QueryOptions = {}) {
      return self.query<{ me: ExtendedUserModelType}>(`query me { me {
        ${typeof resultSelector === "function" ? resultSelector(new ExtendedUserModelSelector()).toString() : resultSelector}
      } }`, variables, options)
    },
    mutateBuyProduct(variables: { input?: BuyProduct }, resultSelector: string | ((qb: PayResponseModelSelector) => PayResponseModelSelector) = payResponseModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ buyProduct: PayResponseModelType}>(`mutation buyProduct($input: BuyProduct) { buyProduct(input: $input) {
        ${typeof resultSelector === "function" ? resultSelector(new PayResponseModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutateCreateProduct(variables: { input: NewProduct }, resultSelector: string | ((qb: ProductModelSelector) => ProductModelSelector) = productModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ createProduct: ProductModelType}>(`mutation createProduct($input: NewProduct!) { createProduct(input: $input) {
        ${typeof resultSelector === "function" ? resultSelector(new ProductModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutatePublishProduct(variables: { input: PublishProduct }, resultSelector: string | ((qb: ProductModelSelector) => ProductModelSelector) = productModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ publishProduct: ProductModelType}>(`mutation publishProduct($input: PublishProduct!) { publishProduct(input: $input) {
        ${typeof resultSelector === "function" ? resultSelector(new ProductModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutateSignup(variables: { email: string, password: string, username: string, name?: string }, resultSelector: string | ((qb: TokenResponseModelSelector) => TokenResponseModelSelector) = tokenResponseModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ signup: TokenResponseModelType}>(`mutation signup($email: String!, $password: String!, $username: String!, $name: String) { signup(email: $email, password: $password, username: $username, name: $name) {
        ${typeof resultSelector === "function" ? resultSelector(new TokenResponseModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutateLogin(variables: { email: string, password: string }, resultSelector: string | ((qb: TokenResponseModelSelector) => TokenResponseModelSelector) = tokenResponseModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ login: TokenResponseModelType}>(`mutation login($email: String!, $password: String!) { login(email: $email, password: $password) {
        ${typeof resultSelector === "function" ? resultSelector(new TokenResponseModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutateForgotPassword(variables: { email?: string }, resultSelector: string | ((qb: BoolResponseModelSelector) => BoolResponseModelSelector) = boolResponseModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ forgotPassword: BoolResponseModelType}>(`mutation forgotPassword($email: String) { forgotPassword(email: $email) {
        ${typeof resultSelector === "function" ? resultSelector(new BoolResponseModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
    mutateChangePassword(variables: { password: string }, resultSelector: string | ((qb: BoolResponseModelSelector) => BoolResponseModelSelector) = boolResponseModelPrimitives.toString(), optimisticUpdate?: () => void) {
      return self.mutate<{ changePassword: BoolResponseModelType}>(`mutation changePassword($password: String!) { changePassword(password: $password) {
        ${typeof resultSelector === "function" ? resultSelector(new BoolResponseModelSelector()).toString() : resultSelector}
      } }`, variables, optimisticUpdate)
    },
  })))
