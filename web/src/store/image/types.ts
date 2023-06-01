import { ImagePullItems, ImagePullOrLoadStatus } from '@/service/image/types'
export interface ImageSpecState {
  pull_items: ImagePullItems
  pull_or_load_status: ImagePullOrLoadStatus
}
