import { Instance } from "mobx-state-tree"
import { ExtendedUserModelBase } from "./ExtendedUserModel.base"

/* The TypeScript type of an instance of ExtendedUserModel */
export interface ExtendedUserModelType extends Instance<typeof ExtendedUserModel.Type> {}

/* A graphql query fragment builders for ExtendedUserModel */
export { selectFromExtendedUser, extendedUserModelPrimitives, ExtendedUserModelSelector } from "./ExtendedUserModel.base"

/**
 * ExtendedUserModel
 */
export const ExtendedUserModel = ExtendedUserModelBase
  .actions(self => ({
    // This is an auto-generated example action.
    log() {
      console.log(JSON.stringify(self))
    }
  }))
