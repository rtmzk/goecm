import type { RouteRecordRaw } from 'vue-router'

export function mapMenuToRoutes(userMenu: any[]): RouteRecordRaw[] {
  const routes: RouteRecordRaw[] = []
  // 加载所有的routes.
  const allRoutes: RouteRecordRaw[] = []

  const routeFiles = import.meta.globEager('/src/router/main/*.ts')
  for (const path in routeFiles) {
    allRoutes.push(routeFiles[path].default)
  }

  const _recurseGetRoute = (menus: any[]) => {
    for (const menu of menus) {
      if (menu.type === 2 || menu.type === 7) {
        const route = allRoutes.find((route) => route.path === menu.url)
        if (route) routes.push(route)
      } else {
        _recurseGetRoute(menu.children)
      }
    }
  }
  _recurseGetRoute(userMenu)
  return routes
}
