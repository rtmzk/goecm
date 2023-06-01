// 统一出口

import { BASE_URL, TIME_OUT } from './request/config'
import GoecmRequest from './request/'
import localCache from '@/utils/cache'

// 携带token的拦截器
const defaultRequest = new GoecmRequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT,
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

export default defaultRequest
