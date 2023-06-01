import { Agents } from './../login/types'
import GoecmRequest from '../request'
import { BASE_URL, TIME_OUT } from '../request/config'
import { rule, rules, DeploySpec, DeployProgress } from './types'
import localCache from '@/utils/cache'

const deployRequest = new GoecmRequest({
  baseURL: BASE_URL,
  timeout: 0,
  interceptors: {
    requestInterceptor: (config) => {
      const token = localCache.getCache('token')
      if (token) {
        if (config.headers) {
          config.headers.Authorization = `Bearer ${token}`
        }
      }
      return config
    }
  }
})

const enum deployApi {
  getEnvCheckRule = '/agent/envc/rules/get',
  doEnvCheck = '/agent/envc/result/get',
  doDeploy = '/agent/deploy',
  getDeployProcess = '/agent/deploy/getDeployProgress'
}

export function getEnvCheckRule() {
  return deployRequest.get<rules>({
    url: deployApi.getEnvCheckRule
  })
}

export function doEnvCheck(data: Agents[]) {
  return deployRequest.post<rule[]>({
    url: deployApi.doEnvCheck,
    data: data
  })
}

export function doDeployRequest(data: DeploySpec) {
  return deployRequest.post({
    url: deployApi.doDeploy,
    data: data
  })
}

export function getDeployProcessRequest() {
  return deployRequest.get<DeployProgress>({
    url: deployApi.getDeployProcess
  })
}
