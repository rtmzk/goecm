export default {
  path: '/main/services',
  name: 'services',
  component: () => import('@/views/main/services/services.vue'),
  parent: 'main',
  children: [],
  meta: {
    title: '服务管理'
  }
}
