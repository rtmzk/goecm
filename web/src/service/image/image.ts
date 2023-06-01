import {
  ImagePullItems,
  ImagePullOrLoadStatus,
  ImagePullOrLoadRequestResponse,
  imageOperation
} from './types'
import GoecmRequest from '../request'
import { BASE_URL, TIME_OUT } from '../request/config'
// import image from '@/router/main/image'
import localCache from '@/utils/cache'
import { CommonResponse } from '../common'

const imagesRequest = new GoecmRequest({
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
enum ImagesAPI {
  ImagePull = '/agent/image/pull',
  ImagePullOrLoadStatus = '/agent/image/pullOrLoadStatus?task=',
  ImageUpload = '/agent/image/upload',
  ImageImport = '/agent/image/import',
  ImageList = '/agent/getReportItems',
  ImageDelete = '/agent/image/delete',
  ImageExport = '/agent/image/export',
  ImageImportProgress = '/agent/image/importProgress'
}

export function imagePull(item: ImagePullItems) {
  return imagesRequest.post<ImagePullOrLoadRequestResponse>({
    url: ImagesAPI.ImagePull,
    data: item,
    timeout: 0
  })
}

export function getPullOrLoadStatus(taskId: string) {
  return imagesRequest.get<ImagePullOrLoadStatus>({
    url: ImagesAPI.ImagePullOrLoadStatus + taskId
  })
}

export function getImageList(node: string = '') {
  return imagesRequest.get({
    url: ImagesAPI.ImageList + `?n=${node}&t=image`
  })
}

export function deleteImageRequest(data: imageOperation[]) {
  return imagesRequest.delete<CommonResponse>({
    url: ImagesAPI.ImageDelete,
    data: data
  })
}

export function exportImageRequest(data: imageOperation) {
  return imagesRequest.post({
    url: ImagesAPI.ImageExport,
    data: data,
    timeout: 0,
    responseType: 'blob'
  })
}

export function postFileUpload(data: FormData) {
  return imagesRequest.post({
    url: ImagesAPI.ImageUpload,
    data: data
  })
}

export function fileMergeRequest(data: FormData) {
  return imagesRequest.post({
    url: ImagesAPI.ImageImport,
    data: data
  })
}

export function getImageLoadProgress() {
  return imagesRequest.get({
    url: ImagesAPI.ImageImportProgress
  })
}

// export function imageUpload(data: ) {
//   return imagesRequest.post<>({
//     url: ImagesAPI.ImageUpload,
//     data: data
//   })
// }
