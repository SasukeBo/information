<template>
  <div class="device-details" v-if="!$apollo.queries.device.loading">
    <div class="device-details__left">
      <div class="global-card base-card">
        <div class="card-title">
          <i class="el-icon-document"></i>
          <span>基本信息</span>
        </div>

        <div class="data-item">
          <span>ID</span>
          <span>{{ device.uuid }}</span>
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
          <span>{{ device.createdAt }}</span>
        </div>

        <div class="data-item">
          <span>最近修改</span>
          <span>{{ device.updatedAt }}</span>
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
import { apollo } from './graphql';
import ParamsValueChart from './_params-value-chart.vue';

export default {
  name: 'device-details',
  props: ['uuid'],
  apollo,
  components: { ParamsValueChart },
  data() {
    return {
      device: {}
    };
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-details {
  display: flex;
  padding: 1rem 0.5rem;
}

.device-details__left {
  flex: 1;
  margin: 0.5rem;

  .card-title {
    padding: 0 0.8rem;
    font-size: 1rem;
    padding-bottom: 0.4rem;
    border-bottom: 1px solid rgba(0, 0, 0, 0.3);
    margin-bottom: 0.6rem;

    i {
      font-size: 1.5rem;
      padding-right: 0.5rem;
      vertical-align: middle;
    }

    span {
      display: inline-block;
      vertical-align: middle;
    }
  }

  .data-item {
    padding: 0 1rem;
    margin-bottom: 0.3rem;
    font-size: 0.85rem;
    color: $--color-font__gray;
    display: flex;
  }

  .data-item span:first-child {
    color: $--color-font__light;
    min-width: 100px;
  }
}

.device-details__right {
  flex: 2;
  margin: 0.5rem;

  .params-realtime-chart {
    padding: 1rem;
  }

  .device-status-card {
    display: flex;
    padding: 1rem;

    .summary {
      flex: auto;
      display: flex;
      flex-flow: wrap;
    }

    .summary-list {
      flex: 1;
    }

    .summary-item {
      padding: 0.5rem 0;

      span {
        font-size: 0.86rem;
        color: $--color-font__gray;
      }

      span:first-child {
        display: inline-block;
        color: $--color-font__light;
        width: 100px;
      }
    }

    .status {
      text-align: center;
      padding-right: 2rem;
      color: $--color-font__gray;
      font-size: 0.85rem;

      .status-text {
        padding-top: 0.5rem;
      }
    }

    .status-icon {
      width: 100px;
      height: 100px;
      line-height: 100px;
      position: relative;

      i.el-icon-s-tools {
        color: $--color-theme__warning;
        font-size: 3rem;
        position: absolute;
        bottom: 0;
        right: 0;
        animation: rotating 1.5s linear infinite;
      }

      i.el-icon-setting {
        font-size: 4.5rem;
        color: $--color-theme__success;
        animation: rotating reverse 1.5s linear infinite;
      }
    }
  }
}
</style>
