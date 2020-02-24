import { RootStoreBase } from "./RootStore.base";
import { Instance, types, getEnv, flow } from "mobx-state-tree";
import { ExtendedUserModel } from "./ExtendedUserModel";
import { useQuery } from "../models";
import { TokenResponseModelType } from "./TokenResponseModel";

export interface RootStoreType extends Instance<typeof RootStore.Type> {}
export const RootStore = RootStoreBase.props({
  currentToken: "",
  currentUser: types.maybe(types.reference(ExtendedUserModel))
})
  .actions(self => ({
    log() {
      console.log(JSON.stringify(self));
    },
    setToken: flow(function* setToken(token: string) {
      self.currentToken = token;
      getEnv(self).gqlHttpClient.setHeaders({
        Authorization: `BEARER ${token}`
      });
      const data = yield self.queryMe();
      self.currentUser = data?.me;
    })
  }))
  .actions(self => ({
    login(values: any){
      return  self.mutateLogin(values, undefined);
      
      // if ("response" in data && "errors" in data.response){
      //   setError("email", "error", data.response.errors[0]);
      //   return
      // }
      // const token = data?.login?.token;
      // if (token !== undefined) {
      //   self.setToken(token);
      // }
    }
  }));

