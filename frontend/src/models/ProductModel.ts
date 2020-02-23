import { Instance } from "mobx-state-tree"
import { ProductModelBase } from "./ProductModel.base"

/* The TypeScript type of an instance of ProductModel */
export interface ProductModelType extends Instance<typeof ProductModel.Type> {}

/* A graphql query fragment builders for ProductModel */
export { selectFromProduct, productModelPrimitives, ProductModelSelector } from "./ProductModel.base"

/**
 * ProductModel
 */
export const ProductModel = ProductModelBase
  .actions(self => ({
    // This is an auto-generated example action.
    log() {
      console.log(JSON.stringify(self))
    }
  }))
