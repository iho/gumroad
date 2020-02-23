import { Instance } from "mobx-state-tree"
import { PayResponseModelBase } from "./PayResponseModel.base"

/* The TypeScript type of an instance of PayResponseModel */
export interface PayResponseModelType extends Instance<typeof PayResponseModel.Type> {}

/* A graphql query fragment builders for PayResponseModel */
export { selectFromPayResponse, payResponseModelPrimitives, PayResponseModelSelector } from "./PayResponseModel.base"

/**
 * PayResponseModel
 */
export const PayResponseModel = PayResponseModelBase
  .actions(self => ({
    // This is an auto-generated example action.
    log() {
      console.log(JSON.stringify(self))
    }
  }))
