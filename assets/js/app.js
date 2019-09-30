import 'css/app.scss';
import './utils/use-element';

import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import apolloProvider from './utils/apollo-provider';
import NProgress from 'js/utils/nprogress';

Vue.use(VueRouter);
Vue.use(Vuex);
window.NProgress = NProgress;


const info = document.querySelector('#vue-entry');

if (info) {
  require.ensure([], () => {
    const entry = require('js/vue').default;
    const store = require('js/vuex').default;
    const router = require('js/router').default;

    new Vue({
      el: info,
      store,
      router,
      apolloProvider,
      render: h => h(entry)
    })
  });
}

if (process.env.NODE_ENV !== 'production') {
  var script = document.createElement('script');
  script.setAttribute('type', 'text/javascript');
  script.setAttribute('src', 'http://localhost:35729/livereload.js');
  document.body.appendChild(script);
}
