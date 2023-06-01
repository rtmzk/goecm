<template>
  <div class="node-form">
    <div class="form-title">
      <div class="label-title">节点IP</div>

      <div class="value-title">节点角色</div>
    </div>
    <div class="form-body">
      <el-form label-width="60%" label-position="left">
        <div v-for="item in agents" :key="item.metadata.id" class="form-iter">
          <el-form-item :label="item.address">
            <div class="checkboxs">
              <el-checkbox v-model="item.is_app" :label="true" size="large"
                >应用节点</el-checkbox
              >
              <el-checkbox v-model="item.is_middle" :label="true" size="large"
                >中间件节点</el-checkbox
              >
              <el-icons
                v-if="agents.length > 1"
                @click="handleAgentRemove(item.address)"
                class="item-del-icon"
                name="delete"
                :size="16"
              ></el-icons>
            </div>
          </el-form-item>
          <hr class="dividing-line" />
        </div>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { Agents } from '@/service/login/types'
import { computed } from 'vue'
import { useStore } from 'vuex'
import ElIcons from '@/components/icons'

const store = useStore()

const agents = computed<Agents[]>(() => store.getters['login/getAgents'])
const handleAgentRemove = (ip: string) => {
  agents.value.forEach((item, index) => {
    if (item.address == ip) {
      agents.value.splice(index, 1)
      return
    }
  })
}
</script>

<style lang="less">
.node-form {
  width: 100%;
  height: 100%;
  margin-top: 36px;

  .form-title {
    width: 70%;
    padding: 0px 8%;
    display: flex;
    margin-bottom: 10px;
    font-size: 1.1rem;
    .label-title {
      text-align: left;
      width: 60%;
    }

    .value-title {
      text-align: left;
      width: 15%;
    }
  }

  .form-body {
    width: 70%;
    padding: 0px 8%;

    .dividing-line {
      border: none;
      height: 2px;
      background-color: #ddd;
      margin-top: 0;
    }

    .item-del-icon {
      padding-left: 10px;
      cursor: pointer;
    }
  }
}
</style>
