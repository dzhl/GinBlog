import Vue from 'vue'
import App from './App.vue'
import router from './router'
import './plugin/antui'
import axios from 'axios'
import './assets/css/style.css'
axios.defaults.baseURL = 'http://localhost:3000/api/v1'
Vue.prototype.$http = axios
Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
