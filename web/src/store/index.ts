import { createStore, Store, useStore as useVueStore } from 'vuex'
import { IStoreType } from './type'

import login from './login/login'
import deploy from './deploy/deploy'
import image from './image/image'
import check from './deploy/check'
import progress from './deploy/progress'

const store = createStore<any>({
  mutations: {},
  getters: {},
  actions: {},
  modules: {
    login,
    deploy,
    image,
    check,
    progress
  }
})

export function setupStore() {
  store.dispatch('login/loadLocalLogin')
  store.dispatch('check/loadLocalCheckRule')
}

export function useStore(): Store<IStoreType> {
  return useVueStore()
}

export default store
