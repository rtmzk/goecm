<template>
  <div class="step-title">
    <el-steps
      :active="props.active"
      :headers="token"
      finish-status="finish"
      process-status="finish"
      simple
    >
      <!-- <el-step v-show="agents.length > 1" title="选择节点"></el-step> -->
      <el-step title="选择节点"></el-step>
      <el-step title="选择镜像文件"></el-step>
      <el-step title="导入镜像包"></el-step>
    </el-steps>
  </div>

  <!-- <div v-show="props.active == 0" class="select-node"> -->

  <div class="step-content">
    <div class="select-node" v-show="props.active == 0">
      <el-table :data="agents" @selection-change="selectionChange">
        <el-table-column label="节点ip" property="address"></el-table-column>
        <el-table-column type="selection" width="55" />
      </el-table>
    </div>
    <div class="select-file" v-show="props.active == 1">
      <el-upload
        ref="uploadRef"
        class="upload"
        :headers="token"
        action="http://192.168.20.174:7000/v1/docker/images/upload"
        :auto-upload="false"
        :limit="1"
        :http-request="upload"
        accept="application/x-tar"
      >
        <div class="upload-content">
          <m-svg-icon name="imgpackage" :size="104"></m-svg-icon>
          <div class="el-upload__text">点击选择文件<em>上传</em></div>
        </div>
      </el-upload>
      <div class="upload-tip">
        <div>您可以上传包含镜像的tar格式的压缩包文件</div>
      </div>
    </div>
  </div>

  <div v-show="props.active == 2" class="upload-process">
    <div class="process">
      <el-progress
        type="circle"
        :progressStatus="progressStatus"
        :width="200"
        :indeterminate="true"
        :duration="2"
        :color="colors"
        :percentage="endOfUpload ? loadPerentage : uploadPercentage"
      >
        <template #default="{ percentage }">
          <span class="percentage">{{ percentage }}%</span>
          <span class="statusText">{{
            endOfUpload ? loadMessage : '正在上传镜像'
          }}</span>
        </template>
      </el-progress>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { Agents } from '@/service/login/types'
import { ElMessage, UploadInstance } from 'element-plus'
import MSvgIcon from '@/components/micons/index'
import LocalCache from '@/utils/cache'
import type { UploadRequestOptions } from 'element-plus'
import {
  postFileUpload,
  fileMergeRequest,
  getImageLoadProgress
} from '@/service/image/image'

const store = useStore()
const active = ref(0)
const agents = computed<Agents[]>(() => store.getters['login/getAgents'])

const token = ref({
  Authorization: 'Bearer ' + LocalCache.getCache('token')
})

const loadMessage = ref('正在导入镜像')
const chunkSize = 5 * 1024 * 1024
const chunkFormData = ref<any>()
const fileHash = ref()
const endOfUpload = ref(false)
const progressStatus = ref('')
const importNode = ref<Array<String>>()

const uploadPercentage = computed(() => {
  if (!chunkFormData.value?.length) return 0
  let uploaded = chunkFormData.value.filter(
    (item: any) => item.percentage != 0
  ).length
  return Number(
    (((uploaded / chunkFormData.value.length) * 100) / 2).toFixed(2)
  )
})
const loadPerentage = ref(50)
const multipleSelection = ref<Agents[]>()
const selectionChange = (row: Agents[]) => {
  importNode.value = []
  multipleSelection.value = row
  multipleSelection.value!.forEach((element) => {
    importNode.value!.push(`${element.address}:${element.port}`)
  })
  console.log(multipleSelection.value)
  console.log(importNode.value)
}
const uploadRef = ref<UploadInstance>()

const getCurrentActive = () => {
  return active.value
}

const setActiveValue = (val: number) => {
  active.value = val
}

const submitUpload = () => {
  uploadRef.value!.submit()
}

const importImageAction = () => {
  console.log('import action')
}
const colors = [
  { color: '#6f7ad3', percentage: 20 },
  { color: '#1989fa', percentage: 40 },
  { color: '#11aeaf', percentage: 60 },
  { color: '#c6c758', percentage: 80 },
  { color: '#67c23a', percentage: 100 }
]

const props = defineProps({
  active: {
    type: Number,
    require: true
  }
})

const upload = (option: UploadRequestOptions) => {
  chunkFormData.value = ''
  endOfUpload.value = false
  let fileChunkList = new Array()
  let cur = 0
  while (cur < option.file.size) {
    fileChunkList.push(option.file.slice(cur, cur + chunkSize))
    cur += chunkSize
  }

  let chunkList = fileChunkList.map((file, index) => {
    return {
      file: file,
      chunkHash: fileHash.value + '-' + index,
      fileHash: fileHash.value
    }
  })

  chunkFormData.value = chunkList.map((chunk, index) => {
    let formData = new FormData()
    formData.append('file', chunk.file)
    formData.append('chunk', String(index + 1))
    formData.append('chunks', String(fileChunkList.length))
    formData.append('fileName', option.file.name)

    return {
      formData: formData,
      percentage: 0
    }
  })

  Promise.all(
    chunkFormData.value.map((data: any) => {
      return new Promise((resolve, reject) => {
        postFileUpload(data.formData)
          .then((res) => {
            resolve(res)
            data.percentage = data.percentage + 1
          })
          .catch((err) => {
            reject(err)
          })
      })
    })
  ).then(() => {
    endOfUpload.value = true
    let formData = new FormData()
    let node = ''
    importNode.value?.forEach((element) => {
      node = element + ','
    })
    node = node.substring(0, node.lastIndexOf(','))
    formData.append('chunks', String(chunkList.length))
    formData.append('fileName', option.file.name)
    formData.append('hosts', node)
    const getImageLoadProgressStatus = setInterval(() => {
      getImageLoadProgress().then((result) => {
        if (result.status === 'error') {
          window.clearInterval(getImageLoadProgressStatus)
          ElMessage({
            type: 'error',
            message: '镜像导入失败'
          })
          loadMessage.value = '镜像导入失败'
        } else if (result.percent !== undefined) {
          loadPerentage.value = Number(result.percent.toFixed(2))
          if (result.status === 'end') {
            ElMessage({
              type: 'success',
              message: '镜像导入成功'
            })
            loadMessage.value = '镜像导入成功'
            window.clearInterval(getImageLoadProgressStatus)
            setTimeout(() => {
              location.reload()
            }, 2000)
          }
        }
      })
    }, 1000)
    fileMergeRequest(formData).then((res: any) => {
      if (!res.task) {
        ElMessage({
          type: 'error',
          message: '镜像导入失败'
        })
      }
    })
  })
}

defineExpose({
  active,
  getCurrentActive,
  setActiveValue,
  importImageAction,
  submitUpload
})
</script>

<style lang="less">
.step-title {
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

.step-content {
  margin-top: 40px;
  .select-file {
    width: 100%;
    height: 100%;
    border: 1px dashed #dcdfe6;
    margin: auto;
    border-radius: 6px;
  }
  .upload {
    height: 100%;
    margin: auto;
    justify-content: center;
    align-items: center;
  }

  .upload-tip {
    position: absolute;
    padding-top: 10px;
    color: #dcdfe6;
  }
}

.percentage {
  display: block;
  margin-top: 10px;
  font-size: 24px;
}

.statusText {
  display: block;
  margin-top: 10px;
  font-size: 16px;
}
</style>
