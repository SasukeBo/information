<template>
  <div class="menu-panel">
    <a
      href="/home"
      :class="['menu-item',  $route.name == 'home' ? 'is-active': '']"
      @click.prevent="directTo('home')"
    >
      <i class="iconfont icon-xuanzhongshangcheng"></i>
      <span>首页</span>
    </a>

    <a
      href="/device/list"
      :class="['menu-item', $route.name == 'device-list' ? 'is-active': '']"
      @click.prevent="directTo('device-list')"
    >
      <i class="iconfont icon-shebei"></i>
      <span>设备</span>
    </a>

    <a
      href="/product/list"
      :class="['menu-item', $route.name == 'product-list' ? 'is-active': '']"
      @click.prevent="directTo('product-list')"
    >
      <i class="iconfont icon-chanpin"></i>
      <span>产品</span>
    </a>

    <a
      href="/settings"
      v-if="false"
      :class="['menu-item', $route.name == 'system-conf' ? 'is-active': '']"
      @click.prevent="directTo('system-conf')"
    >
      <i class="el-icon-s-tools"></i>
      <span>系统设置</span>
    </a>
  </div>
</template>
<script>
export default {
  name: 'manu-panel',
  methods: {
    directTo(pathName) {
      this.$router.push({ name: pathName });
      this.$emit('update:visible', false);
    }
  }
};
</script>
