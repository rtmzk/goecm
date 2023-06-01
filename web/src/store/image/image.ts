import { Module } from 'vuex'
import { ImageSpecState } from './types'
import { imagePull, getPullOrLoadStatus } from '@/service/image/image'
import { ImagePullOrLoadStatus } from '@/service/image/types'

const imageModule: Module<ImageSpecState, any> = {
  namespaced: true,
  state: {
    pull_items: {
      imageName: '',
      node: ''
    },
    pull_or_load_status: {
      taskid: '',
      file_name: '',
      image_name: '',
      percent: ''
    }
  },
  mutations: {
    changePullOrLoadStatus(state, data: ImagePullOrLoadStatus) {
      state.pull_or_load_status = data
    }
  },
  actions: {
    async pullImage({ dispatch, getters }, payload) {
      const imagePullResult = await imagePull(payload)
      if (imagePullResult.taskid) {
        var itsTimeCleanUp = 0
        var pullInterval = setInterval(() => {
          dispatch('pullState', imagePullResult.taskid)
          const pullStatus = getters['getPullOrLoadPercent']
          if (pullStatus == '100') {
            itsTimeCleanUp = 1
          }
        }, 2000)

        while (true) {
          if (itsTimeCleanUp == 1) {
            clearInterval(pullInterval)
          }
          await new Promise((r) => setTimeout(r, 2000))
        }
      }
    },
    async pullState({ commit }, payload: string) {
      const imagePullResult = await getPullOrLoadStatus(payload)
      if (imagePullResult.percent) {
        commit('changePullOrLoadStatus', imagePullResult)
      }
    }
  },
  getters: {
    getPullOrLoadPercent(state) {
      return state.pull_or_load_status.percent
    }
  }
}

export default imageModule
