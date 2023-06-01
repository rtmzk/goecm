<template>
  <div class="nav-header">
    <div @click="redirectToIndex" class="top-header-logo">
      <img src="@/assets/image/top-head-logo.svg" alt="head-logo" />
    </div>
    <div class="content">
      <user-info></user-info>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router'
import UserInfo from './user-info.vue'
const route = useRouter()

const redirectToIndex = () => {
  route.push({
    path: '/main/index'
  })
}
</script>

<style lang="less" scoped>
.nav-header {
  display: flex;
  width: 100%;
  align-items: center;
  height: 60px;

  .top-header-logo {
    position: relative;
    width: 180px;
  }

  .top-header-logo:before {
    content: '';
    position: absolute;
    right: 0;
    height: 50%;
    border-right: 1px solid #d0d4d9;
    top: 25%;
  }
  .content {
    display: flex;
    justify-content: space-between;
    align-items: center;
    flex: 1;
    padding: 0 20px;
  }
}
</style>
