<template>
  <el-icon :size="size" :color="color">
    <component :is="name"></component>
  </el-icon>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import * as Icons from '@element-plus/icons-vue'

export default defineComponent({
  components: Icons,
  name: 'ElIcons',
  props: {
    name: {
      type: String,
      required: true
    },
    size: {
      type: Number || null,
      default: null
    },
    color: {
      type: String,
      default: ''
    }
  }
})
</script>

<style scoped lang="less"></style>
