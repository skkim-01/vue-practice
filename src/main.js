import { createApp } from 'vue'
import App from './App.vue'
import router from '@/router'
import appLayout from '@/layouts/appLayout'
import { vuetify3 } from '@/plugins/vuetify3';
import { setComponents } from './plugins/components';
// import vuetify from '@/plugins/vuetify'

// Vuetify
// import '@mdi/font/css/materialdesignicons.css'
// import 'vuetify/styles'
// import { createVuetify } from 'vuetify'
// import { aliases, mdi } from 'vuetify/iconsets/mdi'
// import * as components from 'vuetify/components'
// import * as directives from 'vuetify/directives'

// easy data table
// import 'vue3-easy-data-table/dist/style.css';
// import Vue3EasyDataTable from 'vue3-easy-data-table';

// const vuetify = createVuetify({
//   ssr: true,
//   components,
//   directives,
//   // https://vuetifyjs.com/en/features/theme/
//   theme: {
//     defaultTheme: 'dark',
//   },
//   icons: {
//     defaultSet: 'mdi',
//     aliases,
//     sets: {
//       mdi,
//     }
//   }
// })

// createApp(App)
//   .use(router)
//   .component('appLayout', appLayout)
//   .component('EasyDataTable', Vue3EasyDataTable)
//   .use(vuetify)
//   .mount('#app')

// createApp(App)
//   .use(router)
//   .component('appLayout', appLayout)
//   .setComponents(app)
//   .use(vuetify3)
//   .mount("#app");

var app = createApp(App)
app.use(router)
app.component('appLayout', appLayout)
setComponents(app)
app.use(vuetify3)
app.mount("#app");