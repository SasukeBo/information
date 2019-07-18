<template>
  <div class="main-passport">
    <div class="title">
      <div class="title-main">设备资讯化</div>
      <div class="title-sub">普创智控</div>
    </div>

    <transition name="fade" mode="out-in">
      <router-view></router-view>
    </transition>
  </div>
</template>
<style lang="scss">
@import 'css/passport/index.scss';
</style>
