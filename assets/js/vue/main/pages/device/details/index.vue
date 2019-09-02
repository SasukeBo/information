<template>
  <div class="device-details" v-if="!$apollo.queries.device.loading">
    <div class="device-details__left">
      <div class="global-card base-card">
        <div class="card-title">
          <i class="el-icon-document"></i>
          <span>基本信息</span>
        </div>

        <div class="data-item">
          <span>名称</span>
          <span>{{ device.name }}</span>
        </div>

        <div class="data-item">
          <span>Mac地址</span>
          <span>{{ device.mac }}</span>
        </div>

        <div class="data-item">
          <span>设备类型</span>
          <span>{{ device.type }}</span>
        </div>

        <div class="data-item">
          <span>描述</span>
          <span>{{ device.description }}</span>
        </div>

        <div class="data-item">
          <span>Token</span>
          <span>{{ device.token }}</span>
        </div>
      </div>

      <div class="global-card register-card">
        <div class="card-title">
          <i class="el-icon-s-order"></i>
          <span>创建信息</span>
        </div>

        <div class="data-item">
          <span>创建人</span>
          <span>{{ device.register.userExtend.name }}</span>
        </div>

        <div class="data-item">
          <span>手机号</span>
          <span>{{ device.register.phone }}</span>
        </div>

        <div class="data-item">
          <span>邮箱</span>
          <span>{{ device.register.userExtend.email }}</span>
        </div>

        <div class="data-item">
          <span>创建日期</span>
          <span>{{ timeFormatter(device.createdAt) }}</span>
        </div>

        <div class="data-item">
          <span>最近修改</span>
          <span>{{ timeFormatter(device.updatedAt) }}</span>
        </div>
      </div>
    </div>

    <div class="device-details__right">
      <div class="global-card device-status-card">
        <div class="status">
          <div class="status-icon">
            <i class="el-icon-s-tools"></i>
            <i class="el-icon-setting"></i>
          </div>

          <div class="status-text">生产中</div>
        </div>

        <div class="summary">
          <div class="summary-list">
            <div class="summary-item">
              <span>IP地址</span>
              <span>127.0.0.1</span>
            </div>

            <div class="summary-item">
              <span>累计运行时长</span>
              <span>365天9小时23分</span>
            </div>

            <div class="summary-item">
              <span>稼动率</span>
              <span>-</span>
            </div>
          </div>

          <div class="summary-list">
            <div class="summary-item">
              <span>良率</span>
              <span>-</span>
            </div>

            <div class="summary-item">
              <span>产能</span>
              <span>-</span>
            </div>
          </div>
        </div>
      </div>

      <params-value-chart :device="device"></params-value-chart>
    </div>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';
import ParamsValueChart from './_params-value-chart.vue';
import deviceQuery from './gql/query.device-get.gql';

export default {
  name: 'device-details',
  props: ['uuid'],
  apollo: {
    device: {
      query: deviceQuery,
      variables() {
        return { uuid: this.uuid };
      }
    }
  },
  components: { ParamsValueChart },
  data() {
    return {
      device: {}
    };
  },
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    }
  }
};
</script>
