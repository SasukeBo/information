<template>
  <div class="base-card">
    <div class="card-title">
      <i class="el-icon-document"></i>
      <span>基本信息</span>
    </div>

    <div class="data-item">
      <span>名称</span>
      <span>{{ device.name }}</span>
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
</template>
<script>
export default {
  name: 'base-card',
  props: ['device']
};
</script>
