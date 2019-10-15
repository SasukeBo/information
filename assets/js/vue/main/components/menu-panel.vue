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
      href="devices"
      :class="['menu-item', $route.name == 'device-list' ? 'is-active': '']"
      @click.prevent="directTo('device-list')"
    >
      <i class="iconfont icon-shebei"></i>
      <span>设备</span>
    </a>

    <a
      href="products"
      :class="['menu-item', $route.name == 'product-list' ? 'is-active': '']"
      @click.prevent="directTo('product-list')"
    >
      <i class="iconfont icon-chanpin"></i>
      <span>产品</span>
    </a>

    <a href="https://github.com/SasukeBo/information/tree/develop" target="_blank" class="menu-item">
      <span>Github 仓库</span>
    </a>

    <a href="https://github.com/SasukeBo/information/issues/5" target="_blank" class="menu-item">
      <span>Github Issues</span>
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
