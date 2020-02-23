import { Instance } from "mobx-state-tree"
import { TokenResponseModelBase } from "./TokenResponseModel.base"

/* The TypeScript type of an instance of TokenResponseModel */
export interface TokenResponseModelType extends Instance<typeof TokenResponseModel.Type> {}

/* A graphql query fragment builders for TokenResponseModel */
export { selectFromTokenResponse, tokenResponseModelPrimitives, TokenResponseModelSelector } from "./TokenResponseModel.base"

/**
 * TokenResponseModel
 */
export const TokenResponseModel = TokenResponseModelBase
  .actions(self => ({
    // This is an auto-generated example action.
    log() {
      console.log(JSON.stringify(self))
    }
  }))
