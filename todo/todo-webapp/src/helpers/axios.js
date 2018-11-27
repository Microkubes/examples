import axios from 'axios'

axios.interceptors.request.use(function(config) {

	config.headers["Content-Type"]  = 'application/x-www-form-urlencoded';
  var token = window.localStorage.getItem('token');
  
  config.headers.Authorization =  token;
  
  console.log(config);
  return config;
});

export default axios;