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
    >
      <div>this is a menu panel</div>
    </el-drawer>
  </div>
</template>
<script>
import IHeader from './components/header';

export default {
  name: 'vue-main',
  components: { IHeader },
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
@import 'css/vars.scss';

.main-pages {
  margin-top: 50px;
}
</style>
