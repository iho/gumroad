/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { QueryBuilder } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { RootStoreType } from "./index"


/**
 * PayResponseBase
 * auto generated base class for the model PayResponseModel.
 */
export const PayResponseModelBase = ModelBase
  .named('PayResponse')
  .props({
    __typename: types.optional(types.literal("PayResponse"), "PayResponse"),
    url: types.union(types.undefined, types.string),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  }))

export class PayResponseModelSelector extends QueryBuilder {
  get url() { return this.__attr(`url`) }
}
export function selectFromPayResponse() {
  return new PayResponseModelSelector()
}

export const payResponseModelPrimitives = selectFromPayResponse().url
