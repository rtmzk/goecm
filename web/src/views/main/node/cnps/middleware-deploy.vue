<template>
  <div class="top">
    <div class="deploy-mode-radio">
      <div class="label">
        <span>中间件部署模式</span>
      </div>
      <div class="radios">
        <el-radio v-model="deploySpec.middleware_mode" label="standalone"
          >单机</el-radio
        >
        <el-radio v-model="deploySpec.middleware_mode" label="cluster"
          >集群</el-radio
        >
      </div>
    </div>

    <div class="deploy-spec-card">
      <el-row>
        <el-col :span="8">
          <el-card v-if="deploySpec.middleware_mode == 'standalone'">
            <template #header>
              <span>Elasticsearch</span>
            </template>
            <div class="itemLabel">
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">IP</div>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="deploySpec.elasticsearch.elasticsearch_hosts"
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">存储路径</div>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="
                          deploySpec.elasticsearch.elasticsearch_data_path
                        "
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">备份路径</div>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="
                          deploySpec.elasticsearch.elasticsearch_backup_path
                        "
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-card>
          <el-card v-if="deploySpec.middleware_mode == 'cluster'">
            <template #header>
              <span>Elasticsearch</span>
            </template>
            <div class="itemRow">
              <div class="itemLabel">
                <el-row>
                  <el-col :span="6">
                    <span class="itemText">IP</span>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="deploySpec.elasticsearch.elasticsearch_hosts"
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
            <div class="itemRow">
              <div class="itemLabel">
                <el-row>
                  <el-col :span="6">
                    <span class="itemText">存储路径</span>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="
                          deploySpec.elasticsearch.elasticsearch_data_path
                        "
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card>
            <template #header>
              <span>RabbitMQ</span>
            </template>
            <div class="itemRow">
              <el-row>
                <el-col :span="6">
                  <span class="itemText">IP</span>
                </el-col>

                <el-col :span="18">
                  <div class="itemInput">
                    <el-input
                      v-model="deploySpec.rabbitmq.rabbitmq_hosts"
                    ></el-input>
                  </div>
                </el-col>
              </el-row>
            </div>
            <div class="itemRow">
              <div class="itemLabel">
                <el-row>
                  <el-col :span="6">
                    <span class="itemText">存储路径</span>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="deploySpec.rabbitmq.rabbitmq_data_path"
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-card>
        </el-col>
        <el-col :span="8">
          <el-card>
            <template #header>
              <span>Redis</span>
            </template>
            <div class="itemRow">
              <div class="itemLabel">
                <el-row>
                  <el-col :span="6">
                    <span class="itemText">IP</span>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="deploySpec.redis.redis_hosts"
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
            <div class="itemRow">
              <div class="itemLabel">
                <el-row>
                  <el-col :span="6">
                    <span class="itemText">存储路径</span>
                  </el-col>
                  <el-col :span="18">
                    <div class="itemInput">
                      <el-input
                        v-model="deploySpec.redis.redis_data_path"
                      ></el-input>
                    </div>
                  </el-col>
                </el-row>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import { DeploySpec } from '@/store/deploy/types'

const store = useStore()
const deploySpec = computed<DeploySpec>(
  () => store.getters['deploy/getDeploySpec']
)
</script>
<style scoped lang="less">
.top {
  width: 100%;
  height: 100%;
  margin-top: 36px;

  .deploy-mode-radio {
    display: flex;
    align-items: center;
    width: 70%;
    padding: 0px 8%;
    .label {
      padding-right: 40px;
    }
  }

  .deploy-spec-card {
    padding: 4% 8% 0px;
    .el-card {
      margin: 0 3%;
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
}
</style>
