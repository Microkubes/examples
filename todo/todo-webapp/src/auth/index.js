import Vue from 'vue'
import router from '@/router'
import store from '@/state/index'
import axios from 'axios'

const LOGIN_URL = process.env.LOGIN_URL

export default {
  install(Vue, options) {
    axios.interceptors.request.use(function(config) {
      const token = store.getters.user.auth_token

      if (token && !config.headers.hasOwnProperty('Authorization')) [
        config.headers.Authorization = token
      ]
      return config
    })
    Vue.prototype.$auth = Vue.auth = this
  },

  login (creds, redirect) {
    const payload = {'email': creds.email, 'password': creds.password}

    var auth = this
    return axios.post(LOGIN_URL, payload)
      .then(function(response) {
        auth._storeToken(response.data)

        if (redirect)
          router.push({name: redirect})

        return null
      })
      .catch(function(errorResponse) {
        return errorResponse
      })
  },

  _storeToken(token) {
    store.commit("setAuthToken", token)
  }
}
