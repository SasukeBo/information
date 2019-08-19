<template>
  <div class="com-header">
    <div class="com-header__left">
      <div class="menu-button">
        <i
          :class="['el-icon-s-unfold', menuOpen ? 'transform-rotate' : '']"
          @click="$emit('toggle-menu')"
        ></i>
      </div>
      <div class="title">
        <div>普创智控</div>
        <div>protron</div>
      </div>
    </div>
    <div class="com-header__center">
      <el-input placeholder="搜索" prefix-icon="el-icon-search"></el-input>
    </div>

    <div class="com-header__right">
      <img src="~images/avatar.jpg" class="avatar" />
    </div>
  </div>
</template>
<script>
import Vue from 'vue';

export default {
  name: 'com-header',
  props: ['menuOpen']
};
</script>
