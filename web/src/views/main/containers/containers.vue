<template>
  <div v-if="$route.path == '/main/containers/console'" class="child-view">
    <router-view></router-view>
  </div>
  <div v-if="$route.fullPath == '/main/containers'" class="container-content">
    <div class="container-content-header">
      <div class="nodeSelector">
        <el-select
          v-model="currentNode"
          placeholder="全部"
          @change="containerNodeSelectChange"
          clearable
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
        <el-button
          plain
          :disabled="!selected"
          :icon="Delete"
          @click="deleteContainerAction"
          >删除</el-button
        >
        <el-input class="searchBar" v-model="searchText" @input="searchInLine">
          <template #suffix>
            <el-icons
              @click="searchInLine"
              class="search-icon"
              name="search"
              :size="16"
            ></el-icons>
          </template>
        </el-input>
      </div>
    </div>
    <div class="container-body">
      <el-table
        :data="
          !searching
            ? containerList.slice(
                (currentPage - 1) * currentPageSize,
                currentPage * currentPageSize
              )
            : searchedList.slice(
                (currentPage - 1) * currentPageSize,
                currentPage * currentPageSize
              )
        "
        style="width: 100%"
        @selection-change="handleSelectionChange"
        :row-key="(row: any) => row.Names[0].split('/internal')"
        empty-text="暂无容器"
      >
        <el-table-column
          type="selection"
          width="55"
          :reserve-selection="true"
        />
        <el-table-column
          prop="Names[0]"
          :formatter="(row: any) => row.Names[0].split('/')"
          :show-overflow-tooltip="true"
          label="名称"
        ></el-table-column>
        <el-table-column label="状态" prop="status" width="120">
          <template #default="scope">
            <div>
              <el-tag
                v-if="scope.row.State == 'running'"
                type="success"
                effect="dark"
                >运行中</el-tag
              >
              <el-tag v-else type="danger" effect="dark">已停止</el-tag>
            </div>
          </template>
        </el-table-column>
        <el-table-column
          prop="Image"
          :formatter="(row: any) => row.Image.split('@')[0]"
          :show-overflow-tooltip="true"
          label="镜像"
        ></el-table-column>
        <el-table-column
          prop="Created"
          :formatter="(row: any) => TimeFormatter(row.Created)"
          label="创建时间"
          width="200"
        ></el-table-column>
        <el-table-column
          prop="Labels['com.docker.stack.namespace']"
          :formatter="(row: any) =>  
            StackFormatter(row.Labels['com.docker.stack.namespace'])
          "
          label="组"
          width="100"
        ></el-table-column>
        <el-table-column prop="host" label="节点" width="200"></el-table-column>
        <el-table-column label="操作" width="160">
          <template #default="scope">
            <el-button @click="detailView(scope.row)" type="text"
              >详情</el-button
            >
            <el-button
              v-if="scope.row.State == 'running'"
              @click="consoleView(scope.row)"
              type="text"
              >控制台</el-button
            >
          </template>
        </el-table-column>
      </el-table>
      <div class="pagination">
        <el-pagination
          v-model:currentPage="currentPage"
          v-model:page-size="currentPageSize"
          :page-sizes="[10, 20, 50]"
          :background="true"
          layout="sizes, prev, pager, next"
          :hide-on-single-page="true"
          :total="!searching ? containerList.length : searchedList.length"
          @size-change="(val: number) => currentPageSize = val"
          @current-change="(val: number) => currentPage = val"
        />
      </div>
    </div>

    <div class="dialogs">
      <el-dialog v-model="openConsoleDialog" width="30%" title="控制台">
        <div class="loginShell">
          <div class="label">
            <span>自定义shell</span>
          </div>
          <el-switch v-model="customLoginShell" />
        </div>
        <div v-if="!customLoginShell" class="loginShell">
          <div class="label">
            <span>登录shell</span>
          </div>
          <el-select v-model="inputOrSelectedShell">
            <el-option
              v-for="item in loginShell"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </div>

        <div class="loginShell" v-if="customLoginShell">
          <div class="label">
            <span>登录shell</span>
          </div>
          <el-input
            style="width: 50%"
            v-model="inputOrSelectedShell"
            placeholder="/bin/bash"
          ></el-input>
        </div>

        <div class="loginShell">
          <div class="label">
            <span>用户名</span>
          </div>
          <el-input
            style="width: 50%"
            v-model="loginUser"
            placeholder="root"
          ></el-input>
        </div>

        <template #footer>
          <el-button plain @click="openConsoleDialog = false">取消</el-button>
          <el-button type="primary" @click="toConsoleView(currentConsoleRow)"
            >连接</el-button
          >
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue'
import { useStore } from 'vuex'
import { Delete } from '@element-plus/icons-vue'
import { Agents } from '@/service/login/types'
import ElIcons from '@/components/icons'
import { useRouter } from 'vue-router'
import { getContainerList } from '@/service/container/container'
import { TimeFormatter } from '@/utils/time'
import { StackFormatter } from '@/utils/formater'
import { deleteContainer, Container } from '@/service/container/types'
import { deleteContainerRequest } from '@/service/container/container'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'

const route = useRouter()
const store = useStore()
const agents = computed<Agents[]>(() => store.getters['login/getAgents'])

const containerList = ref<Array<any>>([])
const searching = ref(false)
const searchedList = ref<Array<any>>([])

const loginShell = [
  {
    label: '/bin/bash',
    value: '/bin/bash'
  },
  {
    label: '/bin/sh',
    value: '/bin/sh'
  },
  {
    label: '/bin/ash',
    value: '/bin/ash'
  }
]

const inputOrSelectedShell = ref('')
const customLoginShell = ref(false)

const currentPage = ref(1)
const currentPageSize = ref(10)

const currentNode = ref('')
const selected = ref(false)
const openConsoleDialog = ref(false)
const loginUser = ref('')

const multipleSelection = ref<Array<any>>([])
const currentConsoleRow = ref<any>()
const searchText = ref('')
const deleteContainers = ref<deleteContainer[]>([])

const containerNodeSelectChange = (val: string) => {
  currentNode.value = val

  if (searchText.value === '') {
    searchedList.value = []

    containerList.value.forEach((e: any) => {
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

  containerList.value.forEach((element: any) => {
    let hasStack = false
    if (element.Labels?.['com.docker.stack.namespace']) {
      hasStack = true
    }
    if (
      element.Names[0].includes(searchText.value) ||
      element.Image.includes(searchText.value) ||
      element.host.includes(searchText.value) ||
      (hasStack &&
        element.Labels?.['com.docker.stack.namespace'].includes(
          searchText.value
        ))
    ) {
      if (currentNode.value === '') {
        searchedList.value.push(element)
      } else if (element.host === currentNode.value) {
        searchedList.value.push(element)
      }
    }
  })

  searching.value = true
}

const deleteContainerAction = () => {
  deleteContainers.value = []
  let temp: deleteContainer = {
    host: '',
    containerId: ['']
  }
  let singleHostDeleteContainers: Map<string, string[]> = new Map<
    string,
    string[]
  >()
  multipleSelection.value.forEach((element: Container) => {
    if (singleHostDeleteContainers.get(element.host)) {
      singleHostDeleteContainers.get(element.host)?.push(element.Id)
    } else {
      singleHostDeleteContainers.set(element.host, [element.Id])
    }
  })
  singleHostDeleteContainers.forEach((v, k) => {
    temp.host = k
    temp.containerId = v
    deleteContainers.value.push(temp)
  })
  deleteContainerRequest(deleteContainers.value).then((res) => {
    if (res.ack) {
      ElMessage({
        type: 'success',
        message: '容器删除成功'
      })

      setTimeout(() => {
        location.reload()
      }, 1000)
    } else {
      ElMessage({
        type: 'error',
        message: '容器删除失败'
      })
    }
  })
}

const handleSelectionChange = (val: Array<any>[]) => {
  multipleSelection.value = val
  if (multipleSelection.value.length > 0) {
    selected.value = true
  } else {
    selected.value = false
  }
}

const detailView = (row: any) => {
  route.push({
    path: '/main/containers/detail',
    query: {
      cid: row.Id,
      host: row.host
    }
  })
}

const toConsoleView = (row: any) => {
  route.push({
    path: '/main/containers/console',
    query: {
      cid: row.Id,
      host: row.host,
      port: row.port,
      command: inputOrSelectedShell.value
    }
  })
}

const consoleView = (row: any) => {
  currentConsoleRow.value = row
  currentConsoleRow.value.port = getCurrentRowHostPort(row)
  openConsoleDialog.value = true
}

const getCurrentRowHostPort = (row: any): number => {
  const host = row.host
  let port = 80
  for (let i = 0; i < agents.value.length; i++) {
    if (agents.value[i].address === host) {
      port = agents.value[i].port
      break
    }
  }
  return port
}

getContainerList().then((containerRes: any) => {
  if (containerRes.length > 0) {
    containerList.value = containerRes
  }
})
</script>

<style lang="less">
.container-content {
  width: 100%;
  height: 100%;
  .container-content-header {
    display: flex;
    flex: 1;
    text-align: center;
    align-items: center;
    padding-bottom: 30px;
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
  .loginShell {
    display: flex;
    text-align: center;
    align-items: center;
    margin: 10px 0;
    .label {
      width: 20%;
      padding-left: 20px;
    }
  }
  .pagination {
    position: absolute;
    padding-top: 20px;
    right: 20px;
  }
}
</style>
