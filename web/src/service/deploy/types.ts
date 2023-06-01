export interface DeploySpec {
  middleware_mode: string
  redis: RedisDeploySpec
  rabbitmq: RabbitmqDeploySpec
  database: DatabaseDeploySpec
  elasticsearch: ElasticsearchDeploySepc
  app: ECMDeploySpec
}

export interface ECMDeploySpec {
  storage: ECMStorageSpec
  scheme: string
  access_port: number
  hosts: string
  docker0_network: string
  docker_gwbridge_network: string
  docker_overlay_macrowing_network: string
}

export interface ECMStorageSpec {
  storage_type: string
  storage_path: string
  storage_url: string
  storage_health_check: string
  storage_ak: string
  storage_sk: string
  storage_bucket: string
}

export interface DatabaseDeploySpec {
  is_external: boolean
  db_type: string
  db_hosts: string
  db_port: string
  db_user: string
  db_pass: string
  db_data_path: string
  db_backup_path: string
}

export interface RedisDeploySpec {
  redis_hosts: string
  redis_data_path: string
}

export interface ElasticsearchDeploySepc {
  elasticsearch_hosts: string
  elasticsearch_data_path: string
  elasticsearch_backup_path: string
}

export interface RabbitmqDeploySpec {
  rabbitmq_hosts: string
  rabbitmq_data_path: string
}

export interface rules {
  rules: rule[]
  code?: number
  message?: string
}

export interface Status {
  status: number
}

export interface rule {
  name: string
  description: string
  func: string
  status: string
  message: string
}

export interface DeployProgress {
  taskack?: boolean
  percent?: string
  message?: string
}
