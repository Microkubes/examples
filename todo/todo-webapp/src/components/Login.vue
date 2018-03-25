<template>
<el-main>
<el-row type="flex" class="row-bg" justify="center">
<el-col :span="8">
</el-col>
<el-col :span="8">
  <el-alert v-if="loginError" v-bind:title="loginErrorMessage" type="error"></el-alert>
  <el-form ref="form" :model="form" label-width="120px">
  <el-form-item label="Email">
    <el-input placeholder="Please input" v-model="form.email"></el-input>
  </el-form-item>
  <el-form-item label="Password">
    <el-input placeholder="Please input" type="password" v-model="form.password"></el-input>
</el-form-item>
  <el-form-item>
  <el-button type="primary" @click="login">Login</el-button>
  </el-form-item>
  </el-form>
</el-col>
<el-col :span="8">
</el-col>
</el-row>
</el-main>
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      form: {
        email: '',
        password: ''
      },
      loginError: false,
      loginErrorMessage: ''
    }
  },
  methods: {
    login: function() {
      var vm = this
      this.$auth.login(this.$data.form, 'home').then((error) => {
        console.log(error)
        if (error) {
          vm.loginError = true
          vm.loginErrorMessage = error.response.data
        }
      })
    }
  }
}

</script>
