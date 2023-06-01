import defaultRequest from '..'
import GoecmRequest from '../request'
import { BASE_URL, TIME_OUT } from '../request/config'
import { Account, AccountToken, MenuObject, Agents } from './types'

const loginRequest = new GoecmRequest({
  baseURL: BASE_URL,
  timeout: TIME_OUT
})

enum LoginAPI {
  AccountLogin = '/login',
  AccountLogout = '/logout',
  CurrentUserInfo = '/users/current',
  MenuInfo = '/menu/list',
  GetAgents = '/agent/all'
}

export function accountLoginRequest(account: Account) {
  return loginRequest.post<AccountToken>({
    url: LoginAPI.AccountLogin,
    data: account
  })
}

export function getCurrentUserInfo() {
  return defaultRequest.get({
    url: LoginAPI.CurrentUserInfo
  })
}

export function getMenus() {
  return defaultRequest.get<MenuObject>({
    url: LoginAPI.MenuInfo
  })
}

export function getAgents() {
  return defaultRequest.get<Agents[]>({
    url: LoginAPI.GetAgents
  })
}
