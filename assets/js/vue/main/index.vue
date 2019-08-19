<template>
  <div class="main-pages">
    <i-header @toggle-menu="menuOpen = !menuOpen" :menuOpen="menuOpen"></i-header>

    <router-view class="page-body"></router-view>

    <el-drawer
      size="300"
      :show-close="false"
      :modal="false"
      :visible.sync="menuOpen"
      custom-class="menu-drawer"
      direction="ltr"
      append-to-body
    >
      <menu-panel></menu-panel>
    </el-drawer>
  </div>
</template>
<script>
import IHeader from './components/header';
import MenuPanel from './components/menu-panel';

export default {
  name: 'vue-main',
  components: { IHeader, MenuPanel },
  data() {
    return {
      menuOpen: false,

      datalist: ''
    };
  },
  mounted() {
    var _this = this;
    var ws = new WebSocket(`ws://${document.location.hostname}/websocket`);

    ws.onopen = function() {
      ws.send('fakeData');
    };
    ws.onmessage = function({ data }) {
      _this.datalist = data;
    };
  }
};
</script>
<style lang="scss">
@import 'css/main/index.scss';
</style>
