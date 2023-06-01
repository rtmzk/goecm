export interface CommonMetadata {
  id: number
  createdAt: Date
  updatedAt: Date
}

export interface CommonResponse {
  success?: string
  code?: number
  message?: string
  ack?: string
}
