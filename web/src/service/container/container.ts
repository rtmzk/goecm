import GoecmRequest from '../request'
import { BASE_URL, TIME_OUT } from '../request/config'
import { deleteContainer } from './types'
import { CommonResponse } from '../common'
import localCache from '@/utils/cache'

const containerRequest = new GoecmRequest({
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

const enum containerApi {
  ContainerList = '/agent/getReportItems',
  ContainerDelete = '/agent/container/delete'
}

export function getContainerList(node: string = '') {
  return containerRequest.get({
    url: containerApi.ContainerList + `?n=${node}&t=container`
  })
}

export function deleteContainerRequest(data: deleteContainer[]) {
  return containerRequest.delete<CommonResponse>({
    url: containerApi.ContainerDelete,
    data: data
  })
}
