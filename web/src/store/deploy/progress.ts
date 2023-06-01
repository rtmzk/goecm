import { DeployProgress, EnvCheckOption } from './types'
import { Module } from 'vuex'
import { getEnvCheckRule } from '@/service/deploy/deploy'
import localCache from '@/utils/cache'

const checkRuleModule: Module<DeployProgress, any> = {
  namespaced: true,
  state() {
    return {
      taskack: false,
      percent: '0%',
      message: ''
    }
  },
  mutations: {
    changeTaskACK(state, payload: boolean) {
      state.taskack = payload
    },
    changeProgress(state, payload: DeployProgress) {
      state.percent = payload.percent
      state.message = payload.message
    }
  },
  actions: {},
  getters: {
    getProgress(state) {
      return state
    }
  }
}

export default checkRuleModule
