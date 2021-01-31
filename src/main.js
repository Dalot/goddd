// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import axios from 'axios';
import Vue from 'vue';
import App from './App';
import router from './router';

Vue.config.productionTip = false;

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
});

// Add a request interceptor
axios.interceptors.request.use((config) => {
  const newConfig = config;
  const token = localStorage.getItem('token');
  newConfig.headers.Authorization = token;
  console.log('axios new config: ', newConfig);
  return newConfig;
});
