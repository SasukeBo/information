<template>
  <div class="device-show">
    <div class="device-show__header">
      <div class="device-show__title">
        <div class="device-title">
          <span class="device-show__device-name">{{ device.name }}</span>
          <span class="device-show__device-type">{{ device.type }}</span>
        </div>
        <hr style="margin-top: 0; border-color: #fff" />
        <div class="device-show__tab-box">
          <transition name="slide-fade" mode="out-in">
            <div class="device-show__tab-name" :key="$route.name">{{ tagTitle }}</div>
          </transition>
        </div>
      </div>

      <div class="device-menu">
        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-details'}"
          @click="$router.push({name: 'device-details'})"
        >
          <i class="el-icon-s-data"></i>
          <span>{{ tabNameMap['device-details'] }}</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-realtime'}"
          @click="$router.push({name: 'device-realtime'})"
        >
          <i class="el-icon-s-marketing"></i>
          <span>{{ tabNameMap['device-realtime'] }}</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-status-log'}"
          @click="$router.push({name: 'device-status-log'})"
        >
          <i class="el-icon-s-management"></i>
          <span>{{ tabNameMap['device-status-log'] }}</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-charges'}"
          @click="$router.push({name: 'device-charges'})"
        >
          <i class="el-icon-s-custom"></i>
          <span>{{ tabNameMap['device-charges'] }}</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-params'}"
          @click="$router.push({name: 'device-params'})"
        >
          <i class="el-icon-s-operation"></i>
          <span>{{ tabNameMap['device-params'] }}</span>
        </div>

        <div
          :class="{'device-menu__item': true, 'is-selected': $route.name === 'device-config'}"
          @click="$router.push({name: 'device-config'})"
        >
          <i class="el-icon-s-tools"></i>
          <span>{{ tabNameMap['device-config'] }}</span>
        </div>
      </div>
    </div>

    <router-view></router-view>
  </div>
</template>
<script>
import deviceQuery from './gql/query.device.gql';

export default {
  name: 'device-show',
  props: ['uuid'],
  apollo: {
    device: {
      query: deviceQuery,
      variables() {
        return { uuid: this.uuid };
      }
    }
  },
  data() {
    return {
      device: {},
      tabNameMap: {
        'device-details': '设备详情',
        'device-charges': '负责人',
        'device-params': '参数',
        'device-realtime': '实时监控',
        'device-status-log': '状态变更',
        'device-config': '管理设备'
      }
    };
  },
  computed: {
    tagTitle() {
      return this.tabNameMap[this.$route.name];
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/device/index.scss';
</style>
