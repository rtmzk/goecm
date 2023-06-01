<template>
  <div class="image-pulling-form">
    <el-form label-width="20%">
      <el-form-item label="镜像名">
        <el-input v-model="pull_items.imageName"></el-input>
      </el-form-item>
      <el-form-item v-if="agent.length > 1" label="选择节点">
        <el-table
          @selection-change="selectionChange"
          :data="agent"
          empty-text="无可用节点"
        >
          <el-table-column label="主机IP" prop="address"></el-table-column>
          <el-table-column type="selection" width="55" />
        </el-table>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import { Agents } from '@/service/login/types'
import { imagePull } from '@/service/image/image'
import { ImagePullItems } from '@/service/image/types'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'

const pull_items = ref<ImagePullItems>({
  imageName: '',
  node: ''
})
const store = useStore()
const agent = computed<Agents[]>(() => store.getters['login/getAgents'])
const pullImageAction = () => {
  imagePull(pull_items.value).then((res: any) => {
    if (res.message) {
      ElMessage({
        type: 'error',
        message: res.message
      })
      return
    }

    ElMessage({
      type: 'success',
      message: `已拉取镜像: ${pull_items.value.imageName}`
    })
    setTimeout(() => {
      window.location.reload()
    }, 1000)
  })
}
const multipleSelection = ref<Agents[]>()
const selectionChange = (row: Agents[]) => {
  multipleSelection.value = row
  multipleSelection.value!.forEach((element) => {
    pull_items.value.node = `${element.address}:${element.port}`
  })
}

defineExpose({
  pullImageAction
})
</script>

<style lang="less"></style>
