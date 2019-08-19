<template>
  <div class="device-card">
    <div class="global-card">
      <div class="data-item device-name">{{ device.name }}</div>

      <div class="data-item device-status">
        <div class="label">设备状态</div>
        <i class="iconfont icon-production" :class="['device-' + device.status]"></i>
        <div class="status">生产中</div>
      </div>

      <div class="data-item">
        <span class="label">注册人</span>
        {{device.user.userExtend.name}}
      </div>

      <div class="data-item">
        <span class="label">设备类型</span>
        {{device.type}}
      </div>

      <div class="data-item">
        <span class="label">注册日期</span>
        {{ device.createdAt }}
      </div>

      <hr />

      <div class="data-item see-more">
        <span class="label">查看详情</span>
        <i class="el-icon-arrow-right"></i>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'device-card',
  props: ['device']
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.user-device .page-right .device-card {
  margin: 0 0.5rem 1rem;
  min-width: 300px;
  flex: 1;

  .data-item {
    padding: 0 1rem;
    margin-bottom: 0.5rem;
    color: $--color-font__gray;

    &:last-child {
      margin-bottom: 0;
    }
  }

  .data-item .label {
    width: 3.5rem;
    display: inline-block;
    font-size: 0.75rem;
    color: $--color-font__light;
  }

  .data-item.device-status {
    text-align: center;
    padding: 1rem 0;

    .iconfont {
      display: block;
      font-size: 2rem;
    }

    .device-producting {
      animation: rotating 1.5s linear infinite;
      color: $--color-theme__success;
    }

    .status {
      color: $--color-font__gray;
    }
  }

  hr {
    margin: 0 1rem;
  }

  .see-more {
    font-size: 0.875rem;
    padding-top: 0.5rem;

    &:before {
      content: ' ';
      display: inline-block;
      height: 1rem;
    }

    .label {
      float: left;
    }

    .el-icon-arrow-right {
      float: right;
    }
  }
}
</style>
