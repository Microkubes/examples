import axios from 'axios'

axios.interceptors.request.use(function(config) {
	// config.headers = {};
	// console.log(config.headers, 'aaaa');
  var token = window.localStorage.getItem('token');
  token = 'eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDIzODQ0OTAsImlhdCI6MTU0MjI5ODA5MCwiaXNzIjoiSm9ybXVuZ2FuZHIgSldUIEF1dGhvcml0eSIsImp0aSI6IjE2NmE1ZjU0LWVkZmMtNGEwZS05ODFmLTk5N2Y2ZDk2MjIyZiIsIm5iZiI6MCwib3JnYW5pemF0aW9ucyI6IiIsInJvbGVzIjoidXNlciIsInNjb3BlcyI6InJlYWQiLCJzdWIiOiI1YmVkODdjY2RlZDA3MDAwMDFmZmNiNjEiLCJ1c2VySWQiOiI1YmVkODdjY2RlZDA3MDAwMDFmZmNiNjEiLCJ1c2VybmFtZSI6InRlc3RAZXhhbXBsZS5jb20ifQ.Ynz1Z9GhoSBaWXH2H6PHKEWRQOGt1of4aMYQ6lfN7FyNZKOOqUja4QvWoX91pBdMXXYhwE56sRQFE6GBSvFyfB8gomnWIbXe865vyygrzSrbEppCYdWkxadgM-mvo4O-7u7RX57rq6LwR2-8WzldttmvVmaa4N3GS4E3iNNZPoOPhfubI97Rc4P4WymgHBbndkCEImOnYknTUHUievdDInOzF17fYDfaq0l2OD3_mZafHlDeY_zNit3NvCZ2pbdTYV_QLQgAQThN_2QGN2uRYosgjPchI-ZQXm4ObzVeJg1Gv87FP12ha99j9eeOih0uPrLJyI5zQCaXfW2y3cGT0w';
  config.headers.Authorization =  'Bearer ' + token;
  
  console.log(config);
  return config;
});

export default axios;