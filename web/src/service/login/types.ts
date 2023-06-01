import { CommonMetadata } from './../common'
export interface Account {
  username: string
  password: string
}

export interface AccountToken {
  expire?: string
  token?: string
  message?: string
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

export interface MenuObject {
  menu: Menu[]
}

export interface UserInfo {
  metadata: CommonMetadata
  username: string
  password: string
  nickname: string
  headerImg: string
  authorityId: string
}

export interface Agents {
  metadata: CommonMetadata
  address: string
  port: number
  is_app?: boolean
  is_middle?: boolean
}
