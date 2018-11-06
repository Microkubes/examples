<template>
  <div id="app">
    <router-link to="/" class="link">Home</router-link> |
    <router-link to="/list" class="link">List</router-link> |
    <router-link to="/login" class="link">Login</router-link>
    <br/>
    <img src="./assets/logo.png" class="logo">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

.logo {
  height: 70px;
}

.link {
  font-size: 40px;
  text-decoration: none;
  margin: 10px;
}
</style>
