<template>
  <div class="node-header">
    <div class="init-btn">
      <el-button type="primary" :disabled="initDisabled" @click="openInitPannel"
        >初始化系统</el-button
      >
    </div>
    <div class="refresh">
      <el-button @click="handleRefreshClick">刷新</el-button>
    </div>

    <div class="progressBar" v-if="initDisabled">
      <el-progress style="width: 300px" :percentage="percent" :status="status">
        {{ message }}
      </el-progress>
    </div>
  </div>

  <div class="agent-list">
    <el-table :data="tableData" style="width: 100%">
      <el-table-column prop="address" label="地址" />
      <el-table-column prop="port" label="端口" />
      <el-table-column prop="metadata.createdAt" label="添加时间" />
      <el-table-column fixed="right" label="操作" width="200">
        <template #default="scope" v-if="tableData.length > 0">
          <el-button type="text" @click="handleAgentView">查看</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

  <div class="init-dialog">
    <el-dialog
      v-model="openInitDialog"
      width="70%"
      top="25vh"
      destroy-on-close
      :before-close="handleClose"
    >
      <template #title>
        <span>系统初始化向导</span>
      </template>
      <init-dialog
        :initing="false"
        @update:changeIniting="changeIniting"
      ></init-dialog>
    </el-dialog>
  </div>
</template>

<script setup lang="ts">
// import { useStore } from 'vuex'
import { ref, computed, onUpdated } from 'vue'
import { getAgents } from '@/service/login/login'
import { useStore } from 'vuex'
import { ElMessageBox } from 'element-plus'
import { DeployProgress } from '@/store/deploy/types'
import { getDeployProcessRequest } from '@/service/deploy/deploy'
import InitDialog from './cnps/init-dialog.vue'
import 'element-plus/theme-chalk/el-message-box.css'
import 'element-plus/theme-chalk/el-progress.css'

const store = useStore()
const tableData = ref([])
const openInitDialog = ref(false)
const percent = ref(0)
const message = ref('')
const status = ref('')
const deployProcess = computed<DeployProgress>(
  () => store.getters['progress/getProgress']
)
const initDisabled = ref(false)

const agents = () => {
  getAgents().then((res: any) => {
    if (res) {
      tableData.value = res
      store.commit('login/changeAgentInfo', res)
    }
  })
}
const changeIniting = (e: any) => {
  openInitDialog.value = !e
  initDisabled.value = e
  const interval = setInterval(() => {
    getDeployProcessRequest().then((res) => {
      if (res) {
        percent.value = Number(res.percent?.replace('%', ''))
        message.value = res.message as string
        console.log(percent.value)
        if (res.percent === '100%') {
          status.value = 'success'
          window.clearInterval(interval)
        }
      }
    })
  }, 3000)
}

const handleRefreshClick = () => {
  agents()
}

const handleAgentView = () => {
  console.log('查看agent信息')
}

const openInitPannel = () => {
  openInitDialog.value = true
}

const handleClose = (done: any) => {
  ElMessageBox.confirm('确定要退出吗', {
    confirmButtonText: '确认',
    cancelButtonText: '取消'
  }).then(() => {
    done()
  })
}
onUpdated(() => {
  if (deployProcess.value.taskack) {
  }
})

agents()
</script>

<style scoped lang="less">
.node-header {
  display: flex;
  align-items: center;

  .el-button {
    margin-right: 10px;
  }
}

.agent-list {
  margin-top: 30px !important;
}

.init-dialog {
  .el-dialog {
    border-radius: 9px !important;
  }
  .el-dialog__header {
    position: absolute;
    left: 5px;
    font-size: 15px;
  }
}
</style>
