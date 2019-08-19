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
      :class="['menu-item', $route.name == 'user-device' ? 'is-active': '']"
      @click.prevent="directTo('device-list')"
    >
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
    </a>

    <a href="javascript:;" :class="['menu-item', $route.name == 'devices' ? 'is-active': '']">
      <i class="iconfont icon-shebei"></i>
      <span>我的设备</span>
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
