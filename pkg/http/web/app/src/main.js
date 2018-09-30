import '@babel/polyfill'
import Vue from 'vue'
import VueRouter from 'vue-router'
// import axios from 'axios'
import './plugins/vuetify'
import App from './App.vue'
import store from './store'


Vue.config.productionTip = false
Vue.use(VueRouter)


// const API_URL = process.env.NODE_ENV == 'production' ? '' : 'http://localhost:3000'

new Vue({
  store,
  render: h => h(App)
}).$mount('#app')
