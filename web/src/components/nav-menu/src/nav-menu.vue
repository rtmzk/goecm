<template>
  <div class="nav-menu">
    <el-menu
      class="el-memu-vertical"
      background-color="#fff"
      :collapse-transition="false"
      active-text-color="#1a8afb"
      text-color="#666666"
      style="--el-menu-hover-bg-color: #f5f9ff"
    >
      <template v-for="item in menus" :key="item.id">
        <el-menu-item
          v-if="item.type == 2"
          :index="item.id"
          @click="handleMenuItemClick(item)"
        >
          <m-svg-icon v-if="item.icon" :name="item.icon"></m-svg-icon>
          <p style="padding-left: 15px">{{ item.zh_name }}</p>
        </el-menu-item>
      </template>
    </el-menu>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useStore } from '@/store'
import { useRouter } from 'vue-router'
import MSvgIcon from '@/components/micons'

const store = useStore()
const router = useRouter()
const menus = computed(() => store.state.login.menuInfo)
const handleMenuItemClick = (item: any) => {
  router.push({
    path: item.url
  })
}
</script>

<style lang="less" scoped>
.el-menu {
  border: none;
  .el-menu-item.is-active {
    background-color: #f6faff !important;
  }

  .el-menu-item:hover {
    color: #1a8afb !important;
  }
}
</style>
