import {
  DatabaseDeploySpec,
  DeploySpec,
  ElasticsearchDeploySepc,
  RabbitmqDeploySpec,
  RedisDeploySpec,
  ECMDeploySpec,
  EnvCheckOption
} from './types'
import { Module } from 'vuex'
import { getEnvCheckRule } from '@/service/deploy/deploy'
import localCache from '@/utils/cache'

const defaultPathPrefix = '/home/edoc2/macrowing/edoc2v5/data/'

const deployModule: Module<DeploySpec, any> = {
  namespaced: true,
  state() {
    return {
      middleware_mode: 'standalone',
      redis: {
        redis_data_path: defaultPathPrefix + 'redis',
        redis_hosts: ''
      },
      rabbitmq: {
        rabbitmq_data_path: defaultPathPrefix + 'rabbitmq',
        rabbitmq_hosts: ''
      },
      elasticsearch: {
        elasticsearch_hosts: '',
        elasticsearch_data_path: defaultPathPrefix + 'es',
        elasticsearch_backup_path: defaultPathPrefix + 'es/backup'
      },
      database: {
        is_external: false,
        db_type: 'mysql',
        db_hosts: '',
        db_backup_path: defaultPathPrefix + 'dbbackup',
        db_data_path: defaultPathPrefix + 'mysql',
        db_user: '',
        db_pass: '',
        db_port: ''
      },
      app: {
        storage: {
          storage_type: '',
          storage_path: defaultPathPrefix + 'edoc2Docs',
          storage_url: '',
          storage_ak: '',
          storage_sk: '',
          storage_health_check: '',
          storage_bucket: ''
        },
        hosts: '',
        scheme: 'http',
        access_port: 80,
        docker0_network: '172.17.0.0/24',
        docker_gwbridge_network: '172.18.0.0/16',
        docker_overlay_macrowing_network: '10.1.0.0/24'
      }
    }
  },
  mutations: {
    changeDeploySpec(state, value: DeploySpec) {
      state = value
    },
    changeDeployMode(state, value: string) {
      state.middleware_mode = value
    },
    changeAppDeploySpec(state, value: ECMDeploySpec) {
      state.app = value
    },
    changeRedisDeploySpec(state, value: RedisDeploySpec) {
      state.redis = value
    },
    changeRabbitmqDeploySpec(state, value: RabbitmqDeploySpec) {
      state.rabbitmq = value
    },
    changeDatabaseDeploySpec(state, value: DatabaseDeploySpec) {
      state.database = value
    },
    changeElasticsearchDeploySpec(state, value: ElasticsearchDeploySepc) {
      state.elasticsearch = value
    }
  },
  actions: {},
  getters: {
    getDeploySpec(state) {
      return state
    },
    getRedisDeploySpec(state) {
      return state.redis
    },
    getDatabaseDeploySpec(state) {
      return state.database
    },
    getElasticsearchDeploySepc(state) {
      return state.elasticsearch
    },
    getRabbitmqDeploySpec(state) {
      return state.rabbitmq
    }
  }
}

export default deployModule
