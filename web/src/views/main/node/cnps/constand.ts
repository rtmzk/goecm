import { DeploySpec } from './../../../../store/deploy/types'
import { useStore } from 'vuex'
import { computed } from 'vue'

export const storageTypes = [
  {
    label: '本地存储/NAS存储',
    value: '0'
  },
  {
    label: '华为对象存储(OBS)',
    value: '7'
  },
  {
    label: '阿里云对象存储(OSS)',
    value: '4'
  },
  {
    label: '腾讯对象存储(COS)',
    value: '8'
  },
  {
    label: 'Ceph存储',
    value: '2'
  },
  {
    label: '兼容性S3存储',
    value: '1000'
  }
]

export const schemeType = [
  {
    label: 'http',
    value: 'http'
  },
  {
    label: 'https',
    value: 'https'
  }
]
