import { createApp } from 'vue'
import { createPinia } from 'pinia'
import 'vuetify/styles'
import createVuetify from '@/plugins/vuetify'

import App from './App.vue'
import router from './router'
import './index.css'

const app = createApp(App)

app.use(createVuetify)
app.use(createPinia())
app.use(router)

app.mount('#app')
