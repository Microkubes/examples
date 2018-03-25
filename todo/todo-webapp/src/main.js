// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuex from 'vuex'
import App from './App'
import router from './router'
import store from "./state/index"
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'

Vue.config.productionTip = false

Vue.use(ElementUI)

import Auth from '@/auth/index'
Vue.use(Auth)

router.beforeEach((to, from, next) => {
  const user = store.getters.user
  if (!user.isLoggedIn) {
    var auth = localStorage.getItem('user')
    if (auth) {
      store.commit('setUserState', auth)
    }
  }
  next()
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
