<template>
  <div
    class="device-card"
    @click="$router.push({name: 'device-show', params: {uuid: device.uuid}})"
  >
    <div class="global-card">
      <div class="data-item device-name">{{ device.name }}</div>

      <div class="data-item device-status">
        <div class="label-center">设备状态</div>
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
