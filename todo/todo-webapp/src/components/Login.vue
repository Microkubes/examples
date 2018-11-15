<template>
<el-main>
<el-row type="flex" class="row-bg" justify="center">
<el-col :span="8">
</el-col>
<el-col :span="8">
  <el-alert v-if="loginError" v-bind:title="loginErrorMessage" type="error"></el-alert>
  <el-form ref="form" :model="form" label-width="120px">
  <el-form-item label="Email">
    <el-input placeholder="Please input" v-model="email"></el-input>
  </el-form-item>
  <el-form-item label="Password">
    <el-input placeholder="Please input" type="password" v-model="password"></el-input>
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
import axios from '../helpers/axios'

export default {
  data() {
    return {
        email: '',
        password: ''
      }
  },
  methods: {
    //Logs in a user and stores its token 
    login: function() {
      var self = this;
      axios.post('http://127.0.0.1:8000/jwt/signin', {
          email: this.$data.email,
          password: this.$data.password
           }).then(function(response) {
              window.localStorage.setItem('token',response.data.token);
              // router.push('List');
           console.log(response);

        });
    }
  },
}

</script>
