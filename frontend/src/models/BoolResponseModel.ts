import { Instance } from "mobx-state-tree"
import { BoolResponseModelBase } from "./BoolResponseModel.base"

/* The TypeScript type of an instance of BoolResponseModel */
export interface BoolResponseModelType extends Instance<typeof BoolResponseModel.Type> {}

/* A graphql query fragment builders for BoolResponseModel */
export { selectFromBoolResponse, boolResponseModelPrimitives, BoolResponseModelSelector } from "./BoolResponseModel.base"

/**
 * BoolResponseModel
 */
export const BoolResponseModel = BoolResponseModelBase
  .actions(self => ({
    // This is an auto-generated example action.
    log() {
      console.log(JSON.stringify(self))
    }
  }))
