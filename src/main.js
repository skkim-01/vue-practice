import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import appLayout from '@/layouts/appLayout'
import { setComponents } from './plugins/components';

var app = createApp(App)
app.use(router)
app.component('appLayout', appLayout)
setComponents(app)
app.mount("#app");