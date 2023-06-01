<template>
  <div class="main">
    <el-container class="main-content">
      <el-aside>
        <div class="head-logo">
          <img
            src="@/assets/image/top-head-logo.svg"
            alt="head-logo"
            class="top-head-logo"
          />
        </div>
        <nav-menu></nav-menu
      ></el-aside>
      <el-container class="page-view-content">
        <el-header class="page-header">
          <nav-user-info></nav-user-info>
        </el-header>

        <el-main>
          <div class="page-content">
            <div class="breadcrumb">
              <el-breadcrumb separator="/">
                <el-breadcrumb-item
                  v-for="(item, index) in breadrumbData"
                  :key="item.path"
                >
                  <span v-if="index === breadrumbData.length - 1">
                    {{ item.meta.title }}
                  </span>
                  <a v-else class="redirect" @click.prevent="onLinkClick(item)">
                    {{ item.meta.title }}
                  </a>
                </el-breadcrumb-item>
              </el-breadcrumb>
            </div>
            <router-view :key="$route.path"></router-view>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import type { RouteLocationMatched } from 'vue-router'
import NavUserInfo from '@/components/nav-header/src/user-info.vue'
import NavMenu from '@/components/nav-menu'

const isCollapse = ref(false)
const handleFoldChange = (isFold: boolean) => {
  console.log(isFold)
  isCollapse.value = isFold
}

const route = useRoute()
const router = useRouter()
const breadrumbData = ref<RouteLocationMatched[]>([])
const getBreadcrumbData = () => {
  breadrumbData.value = route.matched.filter(
    (item) => item.meta && item.meta.title
  )
}

watch(
  route,
  () => {
    getBreadcrumbData()
  },
  {
    immediate: true
  }
)
const onLinkClick = (item: RouteLocationMatched) => {
  console.log(item)
  router.push(item.path)
}
</script>

<style lang="less">
.main {
  position: flex;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: #f7f8f9;
}

.main-content,
.page {
  height: 100%;
}

.page-view-content {
  margin-left: 12px;

  .page-content {
    margin-top: 12px;
  }
}

.el-header,
.el-footer {
  display: flex;
  color: #333;
  text-align: center;
  align-items: center;
}

.el-header {
  height: 60px !important;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background: #fff;
  margin-bottom: 12px;
}

.el-aside {
  width: 200px;
  align-items: center;
  cursor: pointer;
  background-color: #fff;
  transition: width 0.3s linear;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

.el-main {
  position: relative;
  color: #333;
  text-align: center;
  background-color: white;
  padding: 20px;
  overflow-x: auto;
}

.breadcrumb {
  margin: 15px;
  margin-top: 5px;
}

.top-head-logo {
  margin: auto;
  padding-top: 5px;
  padding-bottom: 5px;
  background-color: #fff;
  align-items: center;
  widows: 180px;
}

.top-head-logo:before {
  content: '';
  position: absolute;
  right: 0;
  height: 50%;
  border-right: 1px solid #d0d4d9;
  top: 25%;
}
</style>
