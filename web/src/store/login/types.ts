import { Agents } from '@/service/login/types'
export interface ILoginState {
  token: string
  username: string
  menuInfo: Menu[]
  agents: Agents[]
}
export interface UserInfo {
  metadata: UserMeta
  uuid: string
  username: string
  password: string
  nickname: string
  headerImg: string
  authorityId: string
}

export interface UserMeta {
  id: number
  createAt: string
  updateAt: string
}

export interface Child {
  id: string
  name: string
  parent_id: string
  zh_name: string
  type: number
  url: string
  icon: string
  children?: any
}

export interface Menu {
  id: string
  name: string
  type: number
  parent_id: string
  zh_name: string
  url: string
  icon: string
  children: Child[]
}

export interface MenuObject {
  menu: Menu[]
}

export interface UserInfo {
  metadata: UserMeta
  uuid: string
  username: string
  password: string
  nickname: string
  headerImg: string
  authorityId: string
}

export interface UserMeta {
  id: number
  name: string
  instanceID: string
  createAt: string
  updateAt: string
}
