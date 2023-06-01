import { EnvCheckOptionWrap, EnvCheckOption } from './types'
import { Module } from 'vuex'
import { getEnvCheckRule } from '@/service/deploy/deploy'
import localCache from '@/utils/cache'

const checkRuleModule: Module<EnvCheckOptionWrap, any> = {
  namespaced: true,
  state() {
    return {
      envc: [
        {
          name: '',
          description: '',
          func: '',
          status: '',
          message: ''
        }
      ]
    }
  },
  mutations: {
    changeEnvCheckRules(state, value: EnvCheckOption[]) {
      state.envc = value
    }
  },
  actions: {
    async getEnvCheckRule(state) {
      const rules = await getEnvCheckRule()
      if (rules.rules) {
        localCache.setCache('rules', rules.rules)
      }
      state.commit('changeEnvCheckRules', rules.rules)
    },
    loadLocalCheckRule({ commit }) {
      const rules = localCache.getCache('rules')
      if (rules) {
        commit('changeEnvCheckRules', rules)
      }
    }
  },
  getters: {
    getDeploySpec(state) {
      return state
    },
    getEnvCheckRule(state) {
      return state.envc
    }
  }
}

export default checkRuleModule
