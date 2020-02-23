/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { QueryBuilder } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { RootStoreType } from "./index"


/**
 * BoolResponseBase
 * auto generated base class for the model BoolResponseModel.
 */
export const BoolResponseModelBase = ModelBase
  .named('BoolResponse')
  .props({
    __typename: types.optional(types.literal("BoolResponse"), "BoolResponse"),
    isSuccess: types.union(types.undefined, types.boolean),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  }))

export class BoolResponseModelSelector extends QueryBuilder {
  get isSuccess() { return this.__attr(`isSuccess`) }
}
export function selectFromBoolResponse() {
  return new BoolResponseModelSelector()
}

export const boolResponseModelPrimitives = selectFromBoolResponse().isSuccess
