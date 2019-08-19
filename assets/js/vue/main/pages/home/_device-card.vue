<template>
  <div class="global-card device-card">
    <span class="global-card__title">设备概览</span>

    <div class="data-item">
      <span class="data-label">创建的设备</span>
      <div>
        <el-tag type="info">共{{ own.total }}台</el-tag>
        <el-tag type="success">{{ own.online }}台在线</el-tag>
        <el-tag>{{ own.offline }}台离线</el-tag>
      </div>

      <span class="data-label">负责的设备</span>
      <div>
        <el-tag type="info">共{{ charge.total }}台</el-tag>
        <el-tag type="success">{{ charge.online }}台在线</el-tag>
        <el-tag>{{ charge.offline }}台离线</el-tag>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'device-card',
  data() {
    return {
      own: {
        total: 10,
        online: 8,
        offline: 2
      },
      charge: {
        total: 23,
        online: 23,
        offline: 0
      }
    };
  }
};
</script>
