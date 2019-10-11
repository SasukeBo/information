<template>
  <div class="device-status-card">
    <div class="status" :class="['status_' + statusTag]">
      <div class="status-icon">
        <i class="el-icon-s-tools"></i>
        <i class="el-icon-setting"></i>
      </div>

      <div class="status-text">{{ status }}</div>
    </div>

    <div class="summary">
      <div class="summary-list">
        <div class="summary-item">
          <span>IP地址</span>
          <span>{{ device.remoteIP }}</span>
        </div>

        <div class="summary-item">
          <span>累计运行</span>
          <span>{{ duration }}</span>
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

        <div class="summary-item">
          <span>稼动率</span>
          <span>-</span>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'device-status-card',
  props: ['device', 'duration', 'status', 'statusTag']
};
</script>
