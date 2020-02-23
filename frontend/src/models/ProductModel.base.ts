/* This is a mst-gql generated file, don't modify it manually */
/* eslint-disable */
/* tslint:disable */

import { types } from "mobx-state-tree"
import { MSTGQLRef, QueryBuilder, withTypedRefs } from "mst-gql"
import { ModelBase } from "./ModelBase"
import { UserModel, UserModelType } from "./UserModel"
import { UserModelSelector } from "./UserModel.base"
import { RootStoreType } from "./index"


/* The TypeScript type that explicits the refs to other models in order to prevent a circular refs issue */
type Refs = {
  user: UserModelType;
}

/**
 * ProductBase
 * auto generated base class for the model ProductModel.
 */
export const ProductModelBase = withTypedRefs<Refs>()(ModelBase
  .named('Product')
  .props({
    __typename: types.optional(types.literal("Product"), "Product"),
    id: types.identifier,
    user: types.union(types.undefined, MSTGQLRef(types.late((): any => UserModel))),
    price: types.union(types.undefined, types.integer),
    name: types.union(types.undefined, types.string),
    description: types.union(types.undefined, types.null, types.string),
    summary: types.union(types.undefined, types.null, types.string),
    callToAction: types.union(types.undefined, types.null, types.string),
    coverImage: types.union(types.undefined, types.null, types.string),
    slug: types.union(types.undefined, types.null, types.string),
    isPablished: types.union(types.undefined, types.null, types.boolean),
    receipt: types.union(types.undefined, types.null, types.string),
    content: types.union(types.undefined, types.null, types.string),
  })
  .views(self => ({
    get store() {
      return self.__getStore<RootStoreType>()
    }
  })))

export class ProductModelSelector extends QueryBuilder {
  get id() { return this.__attr(`id`) }
  get price() { return this.__attr(`price`) }
  get name() { return this.__attr(`name`) }
  get description() { return this.__attr(`description`) }
  get summary() { return this.__attr(`summary`) }
  get callToAction() { return this.__attr(`callToAction`) }
  get coverImage() { return this.__attr(`coverImage`) }
  get slug() { return this.__attr(`slug`) }
  get isPablished() { return this.__attr(`isPablished`) }
  get receipt() { return this.__attr(`receipt`) }
  get content() { return this.__attr(`content`) }
  user(builder?: string | UserModelSelector | ((selector: UserModelSelector) => UserModelSelector)) { return this.__child(`user`, UserModelSelector, builder) }
}
export function selectFromProduct() {
  return new ProductModelSelector()
}

export const productModelPrimitives = selectFromProduct().price.name.description.summary.callToAction.coverImage.slug.isPablished.receipt.content
