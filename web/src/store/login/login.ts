import { rules } from './../../views/login/account-config'
import { Agents } from './../../service/login/types'
import { Module } from 'vuex'
// import { IRootState } from '../types'
import { ILoginState } from './types'
import {
  accountLoginRequest,
  getAgents,
  getCurrentUserInfo,
  getMenus
} from '@/service/login/login'

import { getEnvCheckRule } from '@/service/deploy/deploy'

import localCache from '@/utils/cache'
import { mapMenuToRoutes } from '@/utils/map-menu'
import router from '@/router'
import { off } from 'process'
import store from '..'

const loginModule: Module<ILoginState, any> = {
  namespaced: true,
  state() {
    return {
      token: '',
      menuInfo: [],
      username: '',
      agents: [
        {
          metadata: {
            id: 0,
            createdAt: new Date(),
            updatedAt: new Date()
          },
          address: '',
          port: 0,
          is_app: false,
          is_middle: false
        }
      ]
    }
  },
  mutations: {
    changeToken(state, token: string) {
      state.token = token
    },
    changeMenu(state, menu: any) {
      state.menuInfo = menu
      const routes = mapMenuToRoutes(menu)
      routes.forEach((route) => {
        router.addRoute('main', route)
      })
    },
    changeUserInfo(state, username: string) {
      state.username = username
    },
    changeAgentInfo(state, agents: Agents[]) {
      state.agents = agents
    }
  },
  actions: {
    async accountLoginAction({ commit }, payload: any) {
      // 获取token,缓存到浏览器本地
      const loginResult = await accountLoginRequest(payload)
      const token = loginResult.token
      if (!token) {
        return
      }
      commit('changeToken', token)
      localCache.setCache('token', token)

      // 获取菜单信息
      const menuResult = await getMenus()
      const menu = menuResult.menu
      commit('changeMenu', menu)
      localCache.setCache('menu', menu)

      // 获取用户信息
      const userResult = await getCurrentUserInfo()
      const username = userResult.username
      commit('changeUserInfo', username)

      // 浏览器localstorage本地缓存
      localCache.setCache('username', username)
      localCache.setCache('userinfo', userResult)

      //获取agent信息
      const agentResult = await getAgents()
      localCache.setCache('agents', agentResult)
      commit('changeAgentInfo', agentResult)

      this.dispatch('check/getEnvCheckRule')
      router.push('/main/index')
    },

    // vuex: 重新获取用户相关信息
    loadLocalLogin({ commit }) {
      const token = localCache.getCache('token')
      if (token) {
        commit('changeToken', token)
      }

      const menu = localCache.getCache('menu')
      if (menu) {
        commit('changeMenu', menu)
      }

      const user = localCache.getCache('username')
      if (user) {
        commit('changeUserInfo', user)
      }

      const agents = localCache.getCache('agents')
      if (agents) {
        commit('changeAgentInfo', agents)
      }
    }
  },
  getters: {
    getAgents(state) {
      return state.agents
    }
  }
}

export default loginModule
