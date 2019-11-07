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
      <menu-panel :visible.sync="menuOpen"></menu-panel>
    </el-drawer>

    <div class="float-time">{{time}}</div>
  </div>
</template>
<script>
import IHeader from './components/header';
import MenuPanel from './components/menu-panel';
import { timeFormatter } from 'js/utils';

export default {
  name: 'vue-main',
  components: { IHeader, MenuPanel },
  data() {
    return {
      menuOpen: false,
      time: ''
    };
  },
  created() {
    var time = new Date();
    this.time = timeFormatter(time.toISOString());

    setInterval(() => {
      var time = new Date();
      this.time = timeFormatter(time.toISOString());
    }, 1000);
  }
};
</script>
<style lang="scss">
@import 'css/main/index.scss';

.main-pages .float-time {
  position: absolute;
  top: 60px;
  right: 15px;
  z-index: 100;
  color: $--color-theme__light;
  text-shadow: 0 1px 1px $--color-theme__black;
}
</style>
