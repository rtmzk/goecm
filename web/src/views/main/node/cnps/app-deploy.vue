<template>
  <div class="app-top">
    <div class="app-storage-select">
      <div class="label">
        <span>存储类型</span>
      </div>

      <el-select
        class="storage-select"
        v-model="deploySpec.app.storage.storage_type"
        placeholder="--默认本地存储--"
      >
        <el-option
          v-for="item in storageTypes"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>

    <div class="app-storage-spec">
      <div
        v-if="
          deploySpec.app.storage.storage_type == '0' ||
          deploySpec.app.storage.storage_type == ''
        "
        class="localStorage"
      >
        <div class="itemRow">
          <el-row>
            <el-col :span="3">
              <div class="itemText">存储路径</div>
            </el-col>
            <el-col :span="9">
              <el-input
                v-model="deploySpec.app.storage.storage_path"
              ></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
      <div v-else class="s3-storage">
        <div class="itemRow">
          <el-row>
            <el-col :span="4">
              <div class="itemText">连接地址</div>
            </el-col>
            <el-col :span="8">
              <el-input v-model="deploySpec.app.storage.storage_url"></el-input>
            </el-col>
            <el-col :span="4">
              <div class="itemText">存储桶名</div>
            </el-col>
            <el-col :span="8">
              <el-input
                v-model="deploySpec.app.storage.storage_bucket"
              ></el-input>
            </el-col>
          </el-row>
        </div>
        <div class="itemRow">
          <el-row>
            <el-col :span="4">
              <div class="itemText">登录ID</div>
            </el-col>
            <el-col :span="8">
              <el-input v-model="deploySpec.app.storage.storage_ak"></el-input>
            </el-col>
            <el-col :span="4">
              <div class="itemText">登录秘钥</div>
            </el-col>
            <el-col :span="8">
              <el-input
                type="password"
                v-model="deploySpec.app.storage.storage_sk"
              ></el-input>
            </el-col>
          </el-row>
        </div>
        <div class="itemRow">
          <el-row v-if="deploySpec.app.storage.storage_type == '2'">
            <el-col :span="4">
              <div class="itemText">存储健康检查地址</div>
            </el-col>
            <el-col :span="8">
              <el-input
                v-model="deploySpec.app.storage.storage_health_check"
              ></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>

    <div class="divilding-line">
      <hr />
    </div>
    <div class="app-scheme">
      <div class="scheme">
        <div class="label">
          <span>访问方式</span>
        </div>
        <el-select
          class="scheme-select"
          v-model="deploySpec.app.scheme"
          placeholder="--默认http协议--"
        >
          <el-option
            v-for="item in schemeType"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </div>
      <div class="scheme-spec">
        <div class="itemRow">
          <el-row>
            <el-col :span="4">
              <div class="itemText">访问端口</div>
            </el-col>
            <el-col :span="8">
              <el-input
                placeholder="--为空时取默认值 http:80或https:443--"
                v-model.number="deploySpec.app.access_port"
              ></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
    <div class="app-networks">
      <div class="networks-title">
        <div class="label">
          <span>网络名称</span>
        </div>
      </div>
      <div class="app-networks-spec">
        <div class="itemRow">
          <el-row>
            <el-col :span="4"> <div class="itemText">docker0网络</div></el-col>
            <el-col :span="8">
              <el-input v-model="deploySpec.app.docker0_network"></el-input>
            </el-col>
          </el-row>
        </div>
        <div class="itemRow">
          <el-row>
            <el-col :span="4">
              <div class="itemText">gwbridge网络</div>
            </el-col>
            <el-col :span="8">
              <el-input
                v-model="deploySpec.app.docker_gwbridge_network"
              ></el-input>
            </el-col>
          </el-row>
        </div>
        <div class="itemRow">
          <el-row>
            <el-col :span="4">
              <div class="itemText">macrowing网络</div>
            </el-col>
            <el-col :span="8">
              <el-input
                v-model="deploySpec.app.docker_overlay_macrowing_network"
              ></el-input>
            </el-col>
          </el-row>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { storageTypes, schemeType } from './constand'
import { DeploySpec } from '@/store/deploy/types'
import { computed } from 'vue'
import { useStore } from 'vuex'

const store = useStore()
const deploySpec = computed<DeploySpec>(
  () => store.getters['deploy/getDeploySpec']
)
</script>

<style scoped lang="less">
.app-top {
  width: 100%;
  height: 100%;
  margin-top: 36px;
  .app-storage-select {
    display: flex;
    align-items: center;
    width: 70%;
    padding: 0px 8%;
    .label {
      padding-right: 40px;
      font-size: 1rem;
    }
  }

  .app-storage-spec {
    margin-top: 15px;
    width: 70%;
    padding: 0px 8%;
  }
  .divilding-line {
    width: 84%;
    margin-top: 20px;
    margin-bottom: 20px;
    padding-left: 8%;
    hr {
      border: none;
      height: 2px;
      background-color: #ddd;
    }
  }

  .app-scheme {
    .scheme {
      display: flex;
      align-items: center;
      width: 70%;
      padding: 0px 8%;
      .label {
        padding-right: 40px;
        font-size: 1rem;
      }
    }
    .scheme-spec {
      align-items: center;
      width: 70%;
      padding: 0px 8%;
    }
  }

  .app-networks {
    margin-top: 20px;
    width: 70%;
    padding: 0px 8%;
    align-items: center;

    .networks-title {
      display: flex;
    }

    .label {
      font-size: 1rem;
    }
  }

  .itemRow {
    padding-top: 15px;
    align-items: center;

    .itemText {
      display: flex;
      height: 100%;
      align-items: center;
      justify-content: center;
    }
  }
}
</style>
