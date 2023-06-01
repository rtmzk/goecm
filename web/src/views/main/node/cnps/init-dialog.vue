<template>
  <div class="step">
    <el-steps
      :active="active"
      finish-status="finish"
      process-status="finish"
      simple
    >
      <el-step title="节点管理"> </el-step>
      <el-step title="环境检测"> </el-step>
      <el-step title="中间件部署"></el-step>
      <el-step title="数据库部署"></el-step>
      <el-step title="应用部署"></el-step>
      <el-step title="信息确认"></el-step>
    </el-steps>
  </div>

  <div class="step-body">
    <div v-if="active == 0">
      <node-manage></node-manage>
    </div>
    <div v-if="active == 1">
      <env-check></env-check>
    </div>
    <div v-if="active == 2">
      <middleware-deploy></middleware-deploy>
    </div>

    <div v-if="active == 3">
      <database-deploy></database-deploy>
    </div>

    <div v-if="active == 4">
      <app-deploy></app-deploy>
    </div>

    <div v-if="active == 5">
      <deploy-detail></deploy-detail>
    </div>
  </div>

  <div class="footer">
    <div class="prevstep">
      <el-button @click="handlePrevAction" v-if="active > 0">上一步</el-button>
    </div>

    <div class="buttons">
      <el-button @click="handleEnvcAction" v-if="active == 1" type="primary"
        >环境检测</el-button
      >
      <el-button
        @click="handleNextAction"
        v-if="active != 5"
        :disabled="disable"
        type="primary"
        >下一步</el-button
      >
      <el-button @click="handleConfirmInit" type="primary" v-if="active == 5"
        >完成</el-button
      >
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import NodeManage from './node-manage.vue'
import MiddlewareDeploy from './middleware-deploy.vue'
import DatabaseDeploy from './database-deploy.vue'
import EnvCheck from './envc.vue'
import { useStore } from 'vuex'
import AppDeploy from './app-deploy.vue'
import DeployDetail from './deploy-detail.vue'
import { Agents } from '@/service/login/types'
import { DeploySpec } from '@/store/deploy/types'
import {
  doEnvCheck,
  doDeployRequest,
  getDeployProcessRequest
} from '@/service/deploy/deploy'

defineProps({
  initing: Boolean
})
const emit = defineEmits<{
  (e: 'update:changeIniting', value: boolean): void
}>()
const store = useStore()
const agents = computed<Agents[]>(() => store.getters['login/getAgents'])

const active = ref(0)
const disable = ref(false)

const handlePrevAction = () => {
  active.value = active.value - 1
  disable.value = false
}
const handleNextAction = () => {
  if (active.value === 0) {
    setDefaultHost()
  }
  active.value = active.value + 1
  if (active.value == 1) {
    disable.value = true
  }
}

const handleEnvcAction = () => {
  let failedCount = 0
  doEnvCheck(agents.value).then((res) => {
    if (res) {
      store.commit('check/changeEnvCheckRules', res)
      res.forEach((e) => {
        if (e.status === 'FAILED') {
          failedCount = failedCount + 1
        }
      })
      if (failedCount === 0) {
        disable.value = false
      }
    }
  })
}
const handleConfirmInit = () => {
  const deploySpec = computed<DeploySpec>(
    () => store.getters['deploy/getDeploySpec']
  )
  if (deploySpec.value.app.storage.storage_type === '') {
    deploySpec.value.app.storage.storage_type = '0'
  }
  // console.log(deploySpec.value)
  // console.log('完成')
  doDeployRequest(deploySpec.value).then((res: any) => {
    if (res.task) {
      store.commit('progress/changeTaskACK', true)
    }
  })
  emit('update:changeIniting', true)
}
const setDefaultHost = () => {
  if (active.value === 0) {
    const agent = computed<Agents[]>(() => store.getters['login/getAgents'])
    const deploySpec = computed<DeploySpec>(
      () => store.getters['deploy/getDeploySpec']
    )
    deploySpec.value.app.hosts = ''
    deploySpec.value.redis.redis_hosts = ''
    deploySpec.value.elasticsearch.elasticsearch_hosts = ''
    deploySpec.value.rabbitmq.rabbitmq_hosts = ''
    deploySpec.value.database.db_hosts = ''

    agent.value.forEach((item) => {
      if (item.is_app) {
        deploySpec.value.app.hosts =
          deploySpec.value.app.hosts + ',' + item.address
      }
      if (item.is_middle) {
        deploySpec.value.elasticsearch.elasticsearch_hosts =
          deploySpec.value.elasticsearch.elasticsearch_hosts +
          ',' +
          item.address

        deploySpec.value.redis.redis_hosts =
          deploySpec.value.redis.redis_hosts + ',' + item.address

        deploySpec.value.rabbitmq.rabbitmq_hosts =
          deploySpec.value.rabbitmq.rabbitmq_hosts + ',' + item.address

        deploySpec.value.database.db_hosts =
          deploySpec.value.database.db_hosts + ',' + item.address
      }
    })
    deploySpec.value.app.hosts = deploySpec.value.app.hosts.substring(1)
    deploySpec.value.elasticsearch.elasticsearch_hosts =
      deploySpec.value.elasticsearch.elasticsearch_hosts.substring(1)
    deploySpec.value.redis.redis_hosts =
      deploySpec.value.redis.redis_hosts.substring(1)
    deploySpec.value.rabbitmq.rabbitmq_hosts =
      deploySpec.value.rabbitmq.rabbitmq_hosts.substring(1)
    deploySpec.value.database.db_hosts =
      deploySpec.value.database.db_hosts.substring(1)
  }
}
</script>

<style lang="less">
.step {
  margin-top: 40px;

  .el-step__icon.is-text {
    background-color: RGB(173, 176, 183);
  }

  .el-step__head.is-finish {
    .el-step__icon.is-text {
      background-color: RGB(64, 158, 255) !important;
    }
  }

  .el-steps--simple {
    background-color: transparent;
  }
}

.footer {
  display: flex;
  margin: 0px auto;
  width: 100%;
  justify-content: right;

  .el-button {
    margin-top: 20px;
  }

  .prevstep {
    padding-right: 10px;
  }
}
</style>
