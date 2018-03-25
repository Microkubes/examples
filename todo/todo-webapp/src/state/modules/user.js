const state = {
  user: {
    email: '',
    isLoggedIn: false,
    auth_token: '',
    todos: []
  }
}

const getters = {
  user: state => state.user
}

const mutations = {
  setAuthToken (state, token) {
    state.user.auth_token = token
    state.user.isLoggedIn = true
    localStorage.setItem("user", JSON.stringify(state.user))
  },
  setUserTodos (state, todos) {
    state.user.todos = todos
    localStorage.setItem("user", JSON.stringify(state.user))
  },
  setUserState (state, user) {
    state.user = JSON.parse(user)
  }
}

const actions = {
  fetchUserTodos ({ commit, state }) {
    axios.get('http://localhost:8080/todo')
      .then(function (response) {
        commit('setUserTodos', response.data)
      })
      .catch(function(error) {
        console.log(error)
        console.log(error.response.data)
      })
  }
}

export default {
  state,
  getters,
  mutations,
  actions
}
