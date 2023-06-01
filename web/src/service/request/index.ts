import axios from 'axios'
import type { AxiosInstance } from 'axios'
import type { RequestInterceptors, RequestConfig } from './type'

const DEFAULT_LOADING = false

// 该类的对象可以有拦截器
class GoecmRequest {
  instance: AxiosInstance
  interceptors?: RequestInterceptors
  //   loading?: ILoadingInstance
  showLoading?: boolean

  constructor(config: RequestConfig) {
    //创建Axios实例
    this.instance = axios.create(config)

    //保存基本信息
    this.interceptors = config.interceptors
    this.showLoading = config.showLoading ?? true

    // 使用拦截器
    // 从config中取出拦截器对应的实例的拦截器
    this.instance.interceptors.request.use(
      this.interceptors?.requestInterceptor,
      this.interceptors?.requestInterceptorCatch
    )
    this.instance.interceptors.response.use(
      this.interceptors?.responseinterceptor,
      this.interceptors?.responseInterceptorCatch
    )

    // 添加所有实例都有的拦截器
    // this.instance.interceptors.request.use((config) => {
    //   if (this.showLoading == true) {
    //     this.loading = ElLoading.service({
    //       lock: true,
    //       text: '正在请求数据...',
    //       background: 'rgba(0, 0 ,0, 0.5)'
    //     })
    //   }
    //   return config
    // })

    this.instance.interceptors.response.use(
      (res) => {
        // 移除loading
        // setTimeout(() => {
        //   this.loading?.close()
        // }, 1000)

        const data = res.data
        return data
      },
      (err) => {
        return err
      }
    )
  }

  request<T>(config: RequestConfig<T>): Promise<T> {
    return new Promise((resolve, reject) => {
      if (config.interceptors?.requestInterceptor) {
        config = config.interceptors.requestInterceptor(config)
      }

      if (config.showLoading === false) {
        this.showLoading = config.showLoading
      }
      this.instance
        .request<any, T>(config)
        .then((res) => {
          // 单个请求处理数据
          if (config.interceptors?.responseinterceptor) {
            res = config.interceptors.responseinterceptor(res)
          }
          this.showLoading = DEFAULT_LOADING

          resolve(res)
        })
        .catch((err) => {
          this.showLoading = DEFAULT_LOADING
          reject(err)
          return err
        })
    })
  }

  get<T = any>(config: RequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'get'
    })
  }

  post<T>(config: RequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'post'
    })
  }

  delete<T>(config: RequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'delete'
    })
  }

  patch<T>(config: RequestConfig<T>): Promise<T> {
    return this.request<T>({
      ...config,
      method: 'patch'
    })
  }
}

export default GoecmRequest
