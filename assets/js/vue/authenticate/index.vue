<template>
  <div class="main-passport">
    <div class="title">
      <div class="title-main">普创智控设备云平台</div>
      <div class="title-sub">Protron</div>
    </div>

    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>
<script>
export default {
  name: 'authenticate'
};
</script>

<style lang="scss">
@import 'css/authenticate/index.scss';
</style>
