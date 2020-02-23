/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { QueryBuilder } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { RootStoreType } from "./index"


/**
 * TokenResponseBase
 * auto generated base class for the model TokenResponseModel.
 */
export const TokenResponseModelBase = ModelBase
  .named('TokenResponse')
  .props({
    __typename: types.optional(types.literal("TokenResponse"), "TokenResponse"),
    token: types.union(types.undefined, types.string),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  }))

export class TokenResponseModelSelector extends QueryBuilder {
  get token() { return this.__attr(`token`) }
}
export function selectFromTokenResponse() {
  return new TokenResponseModelSelector()
}

export const tokenResponseModelPrimitives = selectFromTokenResponse().token
