import 'css/app.css';
import Vue from 'vue';
import apolloProvider from './apollo-provider';

import Vuex from 'vuex';
import VueRouter from 'vue-router';
Vue.use(VueRouter);
Vue.use(Vuex);

import {
  Form,
  FormItem,
  Button,
  Input,
  Checkbox,
  Message,
  Row,
  Col,
  Tag,
  Drawer
} from 'element-ui'

Vue.use(Form);
Vue.use(FormItem);
Vue.use(Button);
Vue.use(Input);
Vue.use(Checkbox);
Vue.use(Row);
Vue.use(Col);
Vue.use(Tag);
Vue.use(Drawer);
Vue.prototype.$message = Message;

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
