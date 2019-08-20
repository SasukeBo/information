<template>
  <div class="device-show">
    <div class="device-show__header">
      <div class="device-show__title">
        <span class="device-show__title-name">{{ device.name }}</span>
        <span class="device-show__title-type">{{ device.type }}</span>
      </div>

      <div class="device-menu">
        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-details'}"
          @click="$router.push({name: 'device-details'})"
        >
          <i class="el-icon-s-marketing"></i>
          <span>详情</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-charges'}"
          @click="$router.push({name: 'device-charges'})"
        >
          <i class="el-icon-s-custom"></i>
          <span>负责人</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-params'}"
          @click="$router.push({name: 'device-params'})"
        >
          <i class="el-icon-s-operation"></i>
          <span>参数</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-params-values'}"
          @click="$router.push({name: 'device-params-values'})"
        >
          <i class="el-icon-s-data"></i>
          <span>值记录</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-status-log'}"
          @click="$router.push({name: 'device-status-log'})"
        >
          <i class="el-icon-s-management"></i>
          <span>状态变更</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-config'}"
          @click="$router.push({name: 'device-config'})"
        >
          <i class="el-icon-s-tools"></i>
          <span>设置</span>
        </div>
      </div>
    </div>

    <router-view></router-view>
  </div>
</template>
<script>
import { apollo } from './graphql';
export default {
  name: 'device-show',
  props: ['uuid'],
  apollo,

  data() {
    return {
      device: {}
    };
  }
};
</script>
<style lang="scss">
@import 'css/main/device/index.scss';
</style>
