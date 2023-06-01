<template>
  <div class="image-content">
    <div class="image-content-header">
      <div class="nodeSelector">
        <el-select
          v-model="currentNode"
          placeholder="全部"
          @change="nodeSelectChange"
        >
          <el-option
            v-for="item in agents"
            :key="item.metadata.id"
            :label="item.address"
            :value="item.address"
          />
        </el-select>
      </div>
      <div class="operations">
        <el-button plain :disabled="!imageSelected" @click="deleteImage"
          >删除

          <template #icon>
            <m-svg-icon :size="12.5" name="delete"></m-svg-icon>
          </template>
        </el-button>
        <el-button plain @click="importDialogVisable = true"
          >导入
          <template #icon>
            <m-svg-icon
              class="msvgicon"
              :size="12.5"
              name="import"
            ></m-svg-icon>
          </template>
        </el-button>
        <el-button
          plain
          :disabled="!canExport"
          :loading="exporting"
          @click="exportImage"
          >{{ exporting ? '正在导出' : '导出' }}
          <template #icon>
            <m-svg-icon name="export" :size="12.5"></m-svg-icon>
          </template>
        </el-button>
        <el-input class="searchBar" v-model="searchText" @input="searchInLine">
          <template #suffix>
            <m-svg-icon
              name="search"
              style="cursor: pointer"
              @click="searchInLine"
            ></m-svg-icon>
          </template>
        </el-input>
      </div>
      <div class="pulling">
        <el-button
          class="pulling-btn"
          type="primary"
          :loading="pulling"
          @click="pullingDialogVisable = true"
          >{{ pulling ? '正在拉取镜像' : '拉取镜像' }}</el-button
        >
      </div>

      <div class="dialogs">
        <el-dialog v-model="importDialogVisable" width="40%" title="导入镜像">
          <image-import :active="active" ref="importRef" />

          <template #footer>
            <span class="image-pulling-footer">
              <el-button @click="cancelImport">取消</el-button>
              <el-button
                type="primary"
                v-if="(active == 1 && agents.length > 1) || active == 2"
                @click="active = active - 1"
                >上一步</el-button
              >
              <el-button v-if="active == 1" type="primary" @click="submitUpload"
                >上传</el-button
              >
              <el-button
                v-else-if="active == 0"
                type="primary"
                @click="importImageInSon"
                >下一步</el-button
              >
              <el-button v-else type="primary" @click="importDone"
                >完成</el-button
              >
            </span>
          </template>
        </el-dialog>

        <el-dialog v-model="pullingDialogVisable" width="25%" title="拉取镜像">
          <image-pull ref="pullingRef" />
          <template #footer>
            <span class="image-pulling-footer">
              <el-button @click="pullingDialogVisable = false">取消</el-button>
              <el-button type="primary" @click="pullImageActionInSon"
                >拉取镜像</el-button
              >
            </span>
          </template>
        </el-dialog>
      </div>
    </div>
    <div class="image-content-body">
      <el-table
        :data="
          !searching
            ? imageTableData.slice(
                (currentPage - 1) * currentPageSize,
                currentPage * currentPageSize
              )
            : searchedList.slice(
                (currentPage - 1) * currentPageSize,
                currentPage * currentPageSize
              )
        "
        style="width: 100%"
        @selection-change="handleImageSelectionChange"
        empty-text="暂无镜像"
        :row-key="(row: any) => row.Id+row.host"
      >
        <el-table-column
          :reserve-selection="true"
          type="selection"
          width="55"
        />
        <el-table-column
          :show-overflow-tooltip="true"
          prop="RepoTags[0]"
          label="名称"
        ></el-table-column>
        <el-table-column label="状态" width="90">
          <template #default="scope">
            <div>
              <el-tag
                v-if="FormattByImageUsed(scope.row)"
                type="success"
                effect="dark"
                >使用中</el-tag
              >
              <el-tag v-else type="warning" effect="dark">未使用</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="Id"
          label="id"
        ></el-table-column>
        <el-table-column
          :show-overflow-tooltip="true"
          prop="Size"
          :formatter="(row: any) => SizeFormatterFromBytes(row.Size)"
          label="大小"
          width="120"
        ></el-table-column>
        <el-table-column
          width="290"
          :show-overflow-tooltip="true"
          :formatter="(row: any) => TimeFormatter(row.Created)"
          prop="Created"
          label="创建时间"
        ></el-table-column>
        <el-table-column
          width="200"
          :show-overflow-tooltip="true"
          prop="host"
          label="节点"
        ></el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          v-model:currentPage="currentPage"
          v-model:page-size="currentPageSize"
          :page-sizes="[10, 20, 50]"
          :background="true"
          layout="sizes, prev, pager, next"
          :hide-on-single-page="true"
          :total="!searching ? imageTableData.length : searchedList.length"
          @size-change="(val: number) => currentPageSize = val"
          @current-change="(val: number) => currentPage = val"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import ImagePull from './cnps/image-pull.vue'
import ImageImport from './cnps/image-import.vue'
import MSvgIcon from '@/components/micons'
import { Agents } from '@/service/login/types'
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { getImageList } from '@/service/image/image'
import { getContainerList } from '@/service/container/container'
import { SizeFormatterFromBytes } from '@/utils/size'
import { TimeFormatter } from '@/utils/time'
import { imageOperation, Images } from '@/service/image/types'
import { deleteImageRequest, exportImageRequest } from '@/service/image/image'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'

const imageSelected = ref(false)
const canExport = ref(false)
const store = useStore()
const currentNode = ref()
const imageMultipleSelection = ref<Array<any>>()

const pullingDialogVisable = ref(false)
const importDialogVisable = ref(false)
const currentPage = ref(1)
const currentPageSize = ref(10)
const pullingRef = ref()
const importRef = ref()
const active = ref(0)
const imageTableData = ref<Array<any>>([])
const containerListRef = ref<Array<any>>([])
const agents = computed<Agents[]>(() => store.getters['login/getAgents'])
const searchText = ref('')
const searching = ref(false)
const searchedList = ref<Array<any>>([])
const operatingImage = ref<imageOperation[]>([])
const exporting = ref(false)
const pulling = ref(false)

// const resetActive = () => {
//   if (agents.value.length <= 1) {
//     active.value = 1
//   } else {
//     active.value = 0
//   }
// }

const resetActive = () => {
  active.value = 0
}

const cancelImport = () => {
  operatingImage.value = []
  importDialogVisable.value = false
  resetActive()
  location.reload()
}

const importDone = () => {
  importDialogVisable.value = false
  resetActive()
  location.reload()
}

const nodeSelectChange = (val: string) => {
  currentNode.value = val
  if (searchText.value === '') {
    searchedList.value = []
    imageTableData.value.forEach((e: any) => {
      if (e.host.includes(val)) {
        searchedList.value.push(e)
      }
    })

    searching.value = true
  } else {
    searchInLine()
  }
}

const searchInLine = () => {
  if (searchText.value === '') {
    searching.value = false
    return
  }
  searchedList.value = []

  imageTableData.value.forEach((element: any) => {
    if (element.RepoTags) {
      if (
        element.RepoTags![0].includes(searchText.value) ||
        element.Id.includes(searchText.value) ||
        element.host.includes(searchText.value)
      ) {
        if (currentNode.value === '') {
          searchedList.value.push(element)
        } else if (element.host === currentNode.value) {
          searchedList.value.push(element)
        }
      }
    }
  })

  searching.value = true
}

const deleteImage = () => {
  operatingImage.value = []
  let temp: imageOperation = {
    host: '',
    imageId: ['']
  }
  let singleHostDeleteImages: Map<string, string[]> = new Map<
    string,
    string[]
  >()
  imageMultipleSelection.value!.forEach((element: Images) => {
    if (singleHostDeleteImages.get(element.host)) {
      singleHostDeleteImages.get(element.host)?.push(element.Id)
    } else {
      singleHostDeleteImages.set(element.host, [element.Id])
    }
  })
  singleHostDeleteImages.forEach((v, k) => {
    temp.host = k
    temp.imageId = v
    operatingImage.value.push(temp)
  })
  deleteImageRequest(operatingImage.value).then((res) => {
    if (res.ack) {
      ElMessage({
        type: 'success',
        message: '镜像删除成功'
      })
      setTimeout(() => {
        location.reload()
      }, 1000)
    } else {
      ElMessage({
        type: 'error',
        message: '镜像删除失败'
      })
    }
  })
}

const exportImage = () => {
  operatingImage.value = []
  exporting.value = true
  let host = imageMultipleSelection.value![0].host
  let data: imageOperation = {
    host: '',
    imageId: []
  }
  imageMultipleSelection.value?.forEach((element) => {
    data.host = host
    if (element.RepoTags![0]) {
      data.imageId.push(element.RepoTags[0])
    } else {
      data.imageId.push(element.Id)
    }
  })

  exportImageRequest(data).then((res: any) => {
    if (res === {}) {
      ElMessage({
        type: 'error',
        message: '镜像下载失败'
      })
      return
    }
    const content: any = res
    const timestamp = new Date().getTime()
    const fileName = `images${timestamp}.tar`
    const blob = new Blob([content])
    const url = window.URL.createObjectURL(blob)

    let dom = document.createElement('a')
    dom.style.display = 'none'
    dom.href = url
    dom.setAttribute('download', fileName)
    dom.click()
    document.body.appendChild(dom)

    document.body.removeChild(dom)
    window.URL.revokeObjectURL(url)
    exporting.value = false

    ElMessage({
      type: 'success',
      message: `已导出镜像包: ${fileName}`
    })
  })
}

const handleImageSelectionChange = (val: any) => {
  imageMultipleSelection.value = val
  if (imageMultipleSelection.value!.length > 0) {
    imageSelected.value = true
    if (imageMultipleSelection.value!.length === 1) {
      canExport.value = true
    } else {
      canExport.value = false
      let idx0Host = imageMultipleSelection.value![0].host
      let result = imageMultipleSelection.value?.filter(
        (element) => element.host != idx0Host
      )

      if (result!.length == 0) {
        canExport.value = true
      }
    }
  } else {
    imageSelected.value = false
    canExport.value = false
  }
}

const importImageInSon = () => {
  active.value += 1
}

const submitUpload = () => {
  importRef.value.submitUpload()
  active.value += 1
}

const pullImageActionInSon = () => {
  pulling.value = true
  pullingDialogVisable.value = false
  pullingRef.value.pullImageAction()
}

const FormattByImageUsed = (row: any): boolean => {
  for (var i = 0; i < containerListRef.value.length; i++) {
    if (
      containerListRef.value[i].ImageID === row.Id &&
      containerListRef.value[i].State === 'running'
    ) {
      return true
    }
  }
  return false
}

resetActive()
onMounted(() => {
  getImageList().then((ImageRes: any) => {
    if (ImageRes.length > 0) {
      imageTableData.value = ImageRes
    }
  })
  getContainerList().then((containerRes: any) => {
    if (containerRes.length > 0) {
      containerListRef.value = containerRes
    }
  })
})
</script>

<style lang="less">
.image-content {
  width: 100%;
  height: 100%;
  .image-content-header {
    display: flex;
    text-align: center;
    align-items: center;
    .nodeSelector {
      margin-right: 12px;
    }
    .searchBar {
      width: 300px;
      margin: 0 12px;

      .search-icon {
        cursor: pointer;
      }
    }
    .el-input__suffix {
      align-items: center;
    }

    .pulling {
      color: white;
      padding-left: 15px;
      justify-content: space-between;
      flex: 1;
      text-align: right;
    }
  }

  .image-content-body {
    margin-top: 30px;
  }

  .el-dialog {
    border-radius: 8px;
    .el-dialog__header {
      text-align: left;
      padding: 20px 20px 10px 15px;
      .el-dialog__title {
        font-size: 15px;
        font-weight: bold;
      }
    }
  }
  .pagination {
    position: absolute;
    padding-top: 20px;
    right: 20px;
  }
}
</style>
