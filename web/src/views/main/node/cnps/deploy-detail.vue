<template>
  <div class="detail-top">
    <div class="middleware">
      <div class="title">
        <div class="circle"></div>
        <div class="title-text">
          <span>中间件</span>
        </div>
      </div>
      <div class="detail-body">
        <div class="itemRow">
          <div class="itemLabel">中间件部署模式：</div>
          <div
            v-if="deploySpec.middleware_mode == 'standalone'"
            class="itemValue"
          >
            单机
          </div>
          <div class="itemValue" v-else>集群</div>
        </div>
        <div class="itemRow">
          <div class="itemLabel parent">Elasticsearch</div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">IP</div>
          <div class="itemValue childValue">
            {{ deploySpec.elasticsearch.elasticsearch_hosts }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">数据路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.elasticsearch.elasticsearch_data_path }}
          </div>
        </div>
        <div v-if="deploySpec.middleware_mode === 'standalone'" class="itemRow">
          <div class="itemLabel childLabel">备份路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.elasticsearch.elasticsearch_backup_path }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel">RabbitMQ</div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">IP</div>
          <div class="itemValue childValue">
            {{ deploySpec.rabbitmq.rabbitmq_hosts }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">数据路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.rabbitmq.rabbitmq_data_path }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel parent">Redis</div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">IP</div>
          <div class="itemValue childValue">
            {{ deploySpec.redis.redis_hosts }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">数据路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.redis.redis_data_path }}
          </div>
        </div>
      </div>
    </div>
    <div class="database">
      <div class="title">
        <div class="circle"></div>
        <div class="title-text">
          <span>数据库</span>
        </div>
      </div>
      <div class="detail-body">
        <div class="itemRow">
          <div class="itemLabel">数据库部署模式：</div>
          <div v-if="deploySpec.database.is_external" class="itemValue">
            外置
          </div>
          <div v-else class="itemValue">内置</div>
        </div>
        <div v-if="!deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">IP</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_hosts }}
          </div>
        </div>
        <div v-if="!deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">数据路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_data_path }}
          </div>
        </div>
        <div v-if="!deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">备份路径</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_backup_path }}
          </div>
        </div>
        <div v-if="deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">数据库类型</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_type }}
          </div>
        </div>
        <div v-if="deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">连接地址</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_hosts }}
          </div>
        </div>
        <div v-if="deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">连接用户</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_user }}
          </div>
        </div>
        <!-- <div v-if="deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">连接密码</div>
          <div class="itemValue childValue">
            {{ secretText }}
          </div>
          <div class="showPass">
            <el-icons name=""></el-icons>
          </div>
        </div> -->
        <div v-if="deploySpec.database.is_external" class="itemRow">
          <div class="itemLabel childLabel">连接端口</div>
          <div class="itemValue childValue">
            {{ deploySpec.database.db_port }}
          </div>
        </div>
      </div>
    </div>
    <div class="app">
      <div class="title">
        <div class="circle"></div>
        <div class="title-text">
          <span>应用服务</span>
        </div>
      </div>
      <div class="detail-body">
        <div class="itemRow">
          <div class="itemLabel">存储类型：</div>
          <div
            v-if="
              deploySpec.app.storage.storage_type == 'local' ||
              deploySpec.app.storage.storage_type == ''
            "
            class="itemValue"
          >
            本地存储
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 'nas'"
            class="itemValue"
          >
            NAS存储
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 'oss'"
            class="itemValue"
          >
            阿里云对象存储(OSS)
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 'cos'"
            class="itemValue"
          >
            腾讯对象存储(COS)
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 'obs'"
            class="itemValue"
          >
            华为对象存储(OBS)
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 's3'"
            class="itemValue"
          >
            兼容性s3存储
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type == 'ceph'"
            class="itemValue"
          >
            ceph存储
          </div>
        </div>
        <div
          v-if="
            deploySpec.app.storage.storage_type === 's3' ||
            deploySpec.app.storage.storage_type === 'oss' ||
            deploySpec.app.storage.storage_type === 'ceph' ||
            deploySpec.app.storage.storage_type === 'obs' ||
            deploySpec.app.storage.storage_type === 'cos'
          "
          class="s3Storage"
        >
          <div class="itemRow">
            <div class="itemLabel childLabel">存储地址</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_url }}
            </div>
          </div>
          <div class="itemRow">
            <div class="itemLabel childLabel">存储桶名</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_bucket }}
            </div>
          </div>
          <div class="itemRow">
            <div class="itemLabel childLabel">登录ID</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_ak }}
            </div>
          </div>
          <div class="itemRow">
            <div class="itemLabel childLabel">登录秘钥</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_sk }}
            </div>
          </div>
          <div
            v-if="deploySpec.app.storage.storage_type === 'ceph'"
            class="itemRow"
          >
            <div class="itemLabel childLabel">健康检查地址</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_url }}
            </div>
          </div>
        </div>
        <div
          v-if="
            deploySpec.app.storage.storage_type === '' ||
            deploySpec.app.storage.storage_type === 'local' ||
            deploySpec.app.storage.storage_type === 'nas'
          "
          class="localStorage"
        >
          <div class="itemRow">
            <div class="itemLabel childLabel">存储路径</div>
            <div class="itemValue childValue">
              {{ deploySpec.app.storage.storage_path }}
            </div>
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel">访问方式</div>
          <div class="itemValue">
            {{ deploySpec.app.scheme }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel">访问端口</div>
          <div class="itemValue">
            {{ deploySpec.app.access_port }}
          </div>
        </div>

        <div class="itemRow">
          <div class="itemLabel">网络</div>
        </div>

        <div class="itemRow">
          <div class="itemLabel childLabel">docker0网络</div>
          <div class="itemValue childValue">
            {{ deploySpec.app.docker0_network }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">gwbridge网络</div>
          <div class="itemValue childValue">
            {{ deploySpec.app.docker_gwbridge_network }}
          </div>
        </div>
        <div class="itemRow">
          <div class="itemLabel childLabel">macrowing网络</div>
          <div class="itemValue childValue">
            {{ deploySpec.app.docker_overlay_macrowing_network }}
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from 'vuex'
import { DeploySpec } from '@/store/deploy/types'
const store = useStore()
const deploySpec = computed<DeploySpec>(
  () => store.getters['deploy/getDeploySpec']
)

const secretText = '******'
console.log(deploySpec.value)
</script>

<style lang="less">
.detail-top {
  margin-top: 20px;
  width: 70%;
  padding: 0 8%;

  .title {
    display: flex;
    align-items: center;
    font-size: 1rem;
    margin: 20px 0;
  }
  .circle {
    height: 16px;
    width: 16px;
    border-radius: 50%;
    background-color: #63bb5c;
    margin-right: 10px;
  }

  .itemRow {
    display: flex;
    margin: 10px 0;
    width: 100%;

    .childLabel {
      padding: 0 0 0 60px !important;
      justify-content: left;
      width: 15%;
    }

    .childValue {
      width: 60% !important;
    }

    .itemLabel {
      width: 15%;
      padding: 0 20px 0 40px;
      text-align: left;
    }

    .itemValue {
      text-align: left;
      width: 60%;
    }
  }
}
</style>
