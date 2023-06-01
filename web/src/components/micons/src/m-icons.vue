<template>
  <svg aria-hidden="true" :style="{ width: size, height: size }">
    <use class="msvg" :class="fillClass" :xlink:href="symbolId" :fill="color" />
  </svg>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
const props = defineProps({
  name: {
    type: String,
    require: true
  },
  color: {
    type: String
  },
  fillClass: {
    type: String
  },
  size: {
    type: Number || undefined,
    default: 16
  }
})

const symbolId = computed(() => `#icon-${props.name}`)
</script>

<style lang="less" scope></style>
