<template>
  <div class="top">
    <div class="itemRow">
      <el-row>
        <el-col :span="11">
          <div class="database-mode-radio">
            <div class="label">
              <span>数据库部署模式</span>
            </div>
            <div class="radios">
              <el-radio v-model="deploySpec.database.is_external" :label="false"
                >内置容器</el-radio
              >
              <el-radio v-model="deploySpec.database.is_external" :label="true"
                >外置</el-radio
              >
            </div>
          </div>
        </el-col>
        <el-col :span="2">
          <div class="dividling-line"></div>
        </el-col>
        <el-col :span="11">
          <div class="database-deploy-card">
            <div
              class="database-deploy-spec externel"
              v-if="deploySpec.database.is_external"
            >
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">连接地址</div>
                  </el-col>
                  <el-col :span="10">
                    <el-input v-model="deploySpec.database.db_hosts"></el-input>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">账号</div>
                  </el-col>

                  <el-col :span="10">
                    <el-input v-model="deploySpec.database.db_user"></el-input>
                  </el-col>
                </el-row>
              </div>

              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">密码</div>
                  </el-col>
                  <el-col :span="10">
                    <el-input
                      type="password"
                      v-model="deploySpec.database.db_pass"
                    ></el-input>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">端口</div>
                  </el-col>
                  <el-col :span="10">
                    <el-input v-model="deploySpec.database.db_port"></el-input>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">
                      <div>数据库类型</div>
                      <el-tooltip
                        effect="dark"
                        content="MySQL服务版本要求8.0.17;MSSQL服务版本要求2017+"
                        placement="top"
                      >
                        <el-icons name="InfoFilled"></el-icons
                      ></el-tooltip>
                    </div>
                  </el-col>
                  <el-col :span="10">
                    <el-radio
                      v-model="deploySpec.database.db_type"
                      label="mysql"
                      >MySQL</el-radio
                    >
                    <el-radio
                      v-model="deploySpec.database.db_type"
                      label="mssql"
                    >
                      MSSQL
                    </el-radio>
                  </el-col>
                </el-row>
              </div>
            </div>
            <div class="database-deploy-spec inside" v-else>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">IP</div>
                  </el-col>
                  <el-col :span="10">
                    <el-input v-model="deploySpec.database.db_hosts"></el-input>
                  </el-col>
                </el-row>
              </div>

              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">数据存储路径</div>
                  </el-col>
                  <el-col :span="10">
                    <el-input
                      v-model="deploySpec.database.db_data_path"
                    ></el-input>
                  </el-col>
                </el-row>
              </div>
              <div class="itemRow">
                <el-row>
                  <el-col :span="6">
                    <div class="itemText">数据备份路径</div>
                  </el-col>

                  <el-col :span="10">
                    <el-input
                      v-model="deploySpec.database.db_backup_path"
                    ></el-input>
                  </el-col>
                </el-row>
              </div>
            </div>
          </div>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from 'vuex'
import { DeploySpec } from '@/store/deploy/types'
import ElIcons from '@/components/icons'

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

  .database-mode-radio {
    display: flex;
    align-items: center;
    width: 70%;
    padding: 0px 8%;
    .label {
      padding-right: 40px;
    }
  }
  .dividling-line {
    height: 100%;
    border-left: 2px solid #ddd;
  }

  .database-deploy-card {
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

      .el-tooltip__trigger {
        padding-left: 5px;
      }
    }
  }
}
</style>
