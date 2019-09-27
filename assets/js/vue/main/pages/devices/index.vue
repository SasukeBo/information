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
            <el-input placeholder="类型/名称模糊搜索" prefix-icon="el-icon-search" v-model="pattern"></el-input>
          </div>

          <div class="global-card ship-btns">
            <div class="ship-btns__label">
              <el-radio-group v-model="isRegister">
                <el-radio-button :label="true">我的设备</el-radio-button>
                <el-radio-button :label="false">所有设备</el-radio-button>
              </el-radio-group>
            </div>
          </div>
        </div>
      </transition>
    </div>

    <div class="page-right">
      <div class="right-title">设备列表</div>

      <div
        class="device-card-list"
        v-loading="$apollo.queries.deviceList.loading"
        element-loading-background="unset"
      >
        <device-card v-for="device in deviceList.devices" :item="device" :key="device.uuid"></device-card>
        <el-pagination
          :current-page="currentPage"
          @current-change="handleCurrentChange"
          :page-size="limit"
          hide-on-single-page
          layout="total, prev, pager, next, jumper"
          :total="deviceList.total"
        ></el-pagination>

        <div
          class="list__empty"
          v-if="!$apollo.queries.deviceList.loading && !deviceList.total"
        >没有设备</div>
      </div>
    </div>
  </div>
</template>
<script>
import DeviceCard from './_device-card';
import devicesQuery from './gql/query.devices.gql';
import { RadioGroup, RadioButton, Pagination } from 'element-ui';

export default {
  name: 'devices',
  components: {
    ElPagination: Pagination,
    ElRadioGroup: RadioGroup,
    ElRadioButton: RadioButton,
    DeviceCard
  },
  apollo: {
    deviceList: {
      query: devicesQuery,
      variables() {
        return {
          pattern: this.pattern,
          status: this.status,
          isRegister: this.isRegister,
          limit: this.limit,
          offset: this.offset
        };
      },
      fetchPolicy: 'network-only'
    }
  },
  data() {
    return {
      pattern: '',
      status: null,
      isRegister: false,
      limit: 9,
      currentPage: 1,
      isExpand: false,
      deviceList: {
        total: 0,
        devices: []
      }
    };
  },
  computed: {
    offset() {
      return (this.currentPage - 1) * this.limit;
    }
  },
  methods: {
    filterStatus(status) {
      if (this.status === status) this.status = null;
      else this.status = status;
    },
    handleCurrentChange(index) {
      this.currentPage = index;
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/main/devices/index.scss';
</style>
