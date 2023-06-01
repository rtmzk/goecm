<template>
  <div class="user-head-image"></div>
  <el-dropdown trigger="click">
    <el-button type="text" class="el-dropdown-link">
      {{ name }}<i class="el-icon-arrow-down el-icon--right"></i>
    </el-button>
    <template #dropdown>
      <el-dropdown-menu>
        <!-- <el-dropdown-item>用户信息</el-dropdown-item>
        <el-dropdown-item>修改密码</el-dropdown-item> -->
        <el-dropdown-item @click="exitLogin">退出登录</el-dropdown-item>
      </el-dropdown-menu>
    </template>
  </el-dropdown>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '@/store'
import { useRouter } from 'vue-router'
import LocalCache from '@/utils/cache'
const store = useStore()
const route = useRouter()
const name = computed(() => store.state.login.username)
const exitLogin = () => {
  if (LocalCache.getCache('token')) {
    LocalCache.deleteCache('token')
  }
  route.push({
    path: '/login'
  })
}
</script>

<style lang="less" scoped>
.el-dropdown-link {
  color: black;
}

.el-button--text {
  color: RGB(64, 184, 255);
}
</style>
