<template>
  <div class="register-card">
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
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';

export default {
  name: 'register-card',
  props: ['device'],
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    }
  }
};
</script>
