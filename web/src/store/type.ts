import { ILoginState } from './login/types'

interface IRootWithModule {
  login: ILoginState
}

export type IStoreType = IRootWithModule
