import Vue from 'vue'
import App from './App.vue'

import { BootstrapVue, BootstrapVueIcons,IconsPlugin } from 'bootstrap-vue'

// Install BootstrapVue
Vue.use(BootstrapVue)
// Optionally install the BootstrapVue icon components plugin
Vue.use(IconsPlugin)

Vue.use(BootstrapVueIcons)

import 'bootstrap-vue/dist/bootstrap-vue-icons.min.css'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.config.productionTip = false

import VModal from 'vue-js-modal'

Vue.use(VModal)

new Vue({
  render: h => h(App),
}).$mount('#app')