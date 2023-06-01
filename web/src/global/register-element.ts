import { App } from 'vue'
import 'element-plus/theme-chalk/base.css'

import { ElForm, ElButton, ElInput, ElDialog } from 'element-plus'

const components = [ElButton, ElForm, ElInput, ElDialog]
export default function (app: App): void {
  for (const component of components) {
    app.use(component)
  }
}
