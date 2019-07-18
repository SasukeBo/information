import 'css/app.css';
import 'css/theme/index.css';
import ElementUI from 'element-ui';
import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);
Vue.use(ElementUI);

const info = document.querySelector('#vue-entry');

if (info) {
  require.ensure([], () => {
    const entry = require('js/vue').default;
    const router = require('js/router').default;
    // const store = require('js/vuex').default;

    new Vue({
      el: info,
      // apolloProvider,
      render: h => h(entry),
      router,
      // store
    })
  });
}
