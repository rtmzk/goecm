export default {
  path: '/main/index',
  name: 'node',
  component: () => import('@/views/main/node/node.vue'),
  parent: 'main',
  children: [],
  meta: {
    title: '节点管理'
  }
}
