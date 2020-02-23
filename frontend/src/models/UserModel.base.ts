/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { QueryBuilder } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { RootStoreType } from "./index"


/**
 * UserBase
 * auto generated base class for the model UserModel.
 */
export const UserModelBase = ModelBase
  .named('User')
  .props({
    __typename: types.optional(types.literal("User"), "User"),
    id: types.identifier,
    username: types.union(types.undefined, types.string),
    name: types.union(types.undefined, types.string),
    bio: types.union(types.undefined, types.string),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  }))

export class UserModelSelector extends QueryBuilder {
  get id() { return this.__attr(`id`) }
  get username() { return this.__attr(`username`) }
  get name() { return this.__attr(`name`) }
  get bio() { return this.__attr(`bio`) }
}
export function selectFromUser() {
  return new UserModelSelector()
}

export const userModelPrimitives = selectFromUser().username.name.bio
