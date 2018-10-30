import axios from 'axios'

axios.interceptors.request.use(function(config) {
  config.headers.Authorization = `Bearer token`;
  console.log(config);
  return config;
});

export default axios;