import 'css/app.css';
import 'css/theme/index.css';
import ElementUI from 'element-ui';
import Vue from 'vue';
import Vuex from 'vuex';
import VueRouter from 'vue-router';
import apolloProvider from './apollo-provider';

Vue.use(VueRouter);
Vue.use(ElementUI);
Vue.use(Vuex);

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
