import '@babel/polyfill'
import Vue from 'vue'
// import axios from 'axios'
import './plugins/vuetify'
import App from './App.vue'
import store from './store'
import router from './routes'

Vue.config.productionTip = false

// const API_URL = process.env.NODE_ENV == 'production' ? '' : 'http://localhost:3000'
router.beforeEach(Vue.prototype.$auth.authRedirectGuard())

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount('#app')
