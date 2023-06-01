export default {
  path: '/main/containers',
  name: 'containers',
  parent: 'main',
  component: () => import('@/views/main/containers/containers.vue'),
  children: [
    {
      path: '/main/containers/detail',
      component: () =>
        import('@/views/main/extraview/cnps/extraDetailView.vue'),
      name: 'detail',
      parent: 'extraContainers',
      meta: {
        title: '容器详情'
      }
    },
    {
      path: '/main/containers/console',
      component: () =>
        import('@/views/main/extraview/cnps/extraConsoleView.vue'),
      name: 'console',
      parent: 'extraContainers',
      meta: {
        title: '控制台'
      }
    }
  ],
  meta: {
    title: '容器管理'
  }
}
