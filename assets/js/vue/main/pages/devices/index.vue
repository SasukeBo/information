<template>
  <div class="user-device">
    <div class="page-left toggle-options">
      <div class="left-title">
        <span>筛选</span>
        <i class="el-icon-caret-bottom" @click="isExpand = !isExpand"></i>
      </div>

      <transition name="expand">
        <div v-show="isExpand" class="expand-panel">
          <div class="global-card status-btns">
            <div class="status-btn status_prod" @click="filterStatus('prod')">
              <div class="status-btn__icon" :class="[status === 'prod' ? 'is-selected' : '']">
                <i class="iconfont icon-production"></i>
              </div>

              <div class="status-btn__label">在生产</div>
            </div>

            <div class="status-btn status_stop" @click="filterStatus('stop')">
              <div class="status-btn__icon" :class="[status === 'stop' ? 'is-selected' : '']">
                <i class="iconfont icon-production"></i>
              </div>

              <div class="status-btn__label">停机中</div>
            </div>

            <div class="status-btn status_offline" @click="filterStatus('offline')">
              <div class="status-btn__icon" :class="[status === 'offline' ? 'is-selected' : '']">
                <i class="iconfont icon-off-line"></i>
              </div>

              <div class="status-btn__label">已离线</div>
            </div>
          </div>

          <div class="global-card name-search">
            <el-input placeholder="搜索设备" prefix-icon="el-icon-search" v-model="search"></el-input>
          </div>

          <div class="global-card ship-btns">
            <div class="ship-btns__label">关系类型：</div>
            <el-checkbox-group v-model="checkboxGroup">
              <el-checkbox-button label="register">创建的</el-checkbox-button>
              <el-checkbox-button label="charger">负责的</el-checkbox-button>
            </el-checkbox-group>
          </div>
        </div>
      </transition>
    </div>

    <div class="page-right">
      <div></div>
      <div class="right-title">设备列表</div>

      <div class="device-card-list">
        <device-card
          v-for="device in devices.filter(d => !status || d.status === status)"
          :item="device"
          v-bind:device.sync="device"
          :key="device.uuid"
        ></device-card>
      </div>
    </div>
  </div>
</template>
<script>
import DeviceCard from './_device-card';
import devicesQuery from './gql/query.devices.gql';

export default {
  name: 'devices',
  components: { DeviceCard },
  apollo: {
    devices: {
      query: devicesQuery,
      variables() {
        return {
          ownership: this.checkboxGroup,
          namePattern: this.search
        };
      }
    }
  },
  data() {
    return {
      search: '',
      checkboxGroup: ['register'],
      isExpand: false,
      devices: [],
      status: ''
    };
  },
  methods: {
    filterStatus(status) {
      if (this.status === status) this.status = '';
      else this.status = status;
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/devices/index.scss';
</style>
