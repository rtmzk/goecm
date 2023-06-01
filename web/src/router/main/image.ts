export default {
  path: '/main/images',
  name: 'images',
  component: () => import('@/views/main/images/images.vue'),
  parent: 'main',
  children: [],
  meta: {
    title: '镜像管理'
  }
}
