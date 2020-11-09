// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Cookies from 'js-cookie'
import App from './App'
import router from './router'
import store from './store'
import './permission' // permission control

import Vant from 'vant';
import Toast from 'vant';
import 'vant/lib/index.css';

Vue.use(Vant);
Vue.use(Toast);

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
