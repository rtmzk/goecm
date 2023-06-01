import { createApp } from 'vue'
import router from '@/router/index'
import 'normalize.css'
import '@/assets/css/index.less'
import 'virtual:svg-icons-register'

import App from './App.vue'
import store, { setupStore } from './store'

const app = createApp(App)

app.use(store)
setupStore()
app.use(router)
app.mount('#app')
