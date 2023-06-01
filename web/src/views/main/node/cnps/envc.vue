<template>
  <div class="check-rule-content">
    <div
      v-for="(ruleItem, ruleIdx) in checkrules"
      :key="ruleIdx"
      class="check-rule-item"
    >
      <el-row :gutter="0">
        <el-col :span="4">
          <div class="check-rule-item-name">
            <span>{{ ruleItem.name }}：</span>
          </div>
        </el-col>
        <el-col :span="10">
          <div class="check-rule-item-desc">
            {{ ruleItem.description }}
          </div>
        </el-col>
        <el-col :span="10">
          <div class="check-rule-item-status">
            <el-icon
              class="check-rule-item-status-icon__success"
              :size="20"
              v-if="ruleItem.status == 'OK'"
            >
              <el-icons name="CircleCheckFilled" color="#67c23a"></el-icons>
            </el-icon>
            <el-popover
              v-if="ruleItem.status == 'FAILED'"
              placement="top"
              :width="600"
              trigger="click"
            >
              <p>{{ ruleItem.message }}</p>
              <template #reference>
                <el-icon class="check-rule-item-status-icon" :size="20">
                  <el-icons name="CircleCloseFilled" color="#f56c6c"></el-icons>
                </el-icon>
              </template>
            </el-popover>
          </div>
        </el-col>
      </el-row>
      <hr class="dividing-line" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { useStore } from 'vuex'
import ElIcons from '@/components/icons'

const store = useStore()
const checkrules = computed(() => store.getters['check/getEnvCheckRule'])
</script>

<style lang="less" scoped>
.el-form-item__content {
  width: 100%;
}
.check-rule-content {
  width: 80%;
  text-align: center;
  padding-top: 20px;
  margin: auto;
}
.check-rule-item-name {
  // text-align: right;
  padding-top: 15px;
  padding-right: 15px;
}
.check-rule-item-desc {
  padding-top: 15px;
  text-align: left;
}
.check-rule-item-status {
  padding-top: 15px;
  text-align: right;

  .check-rule-item-status-icon:hover {
    cursor: pointer;
  }
}

.dividing-line {
  border: none;
  height: 1px;
  background-color: #ddd;
  margin-top: 8px;
}

span {
  float: left;
}
</style>
