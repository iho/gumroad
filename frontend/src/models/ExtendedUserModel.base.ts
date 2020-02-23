/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { QueryBuilder } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { RootStoreType } from "./index"


/**
 * ExtendedUserBase
 * auto generated base class for the model ExtendedUserModel.
 */
export const ExtendedUserModelBase = ModelBase
  .named('ExtendedUser')
  .props({
    __typename: types.optional(types.literal("ExtendedUser"), "ExtendedUser"),
    id: types.identifier,
    username: types.union(types.undefined, types.string),
    name: types.union(types.undefined, types.string),
    bio: types.union(types.undefined, types.string),
    balance: types.union(types.undefined, types.integer),
    email: types.union(types.undefined, types.string),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  }))

export class ExtendedUserModelSelector extends QueryBuilder {
  get id() { return this.__attr(`id`) }
  get username() { return this.__attr(`username`) }
  get name() { return this.__attr(`name`) }
  get bio() { return this.__attr(`bio`) }
  get balance() { return this.__attr(`balance`) }
  get email() { return this.__attr(`email`) }
}
export function selectFromExtendedUser() {
  return new ExtendedUserModelSelector()
}

export const extendedUserModelPrimitives = selectFromExtendedUser().username.name.bio.balance.email
