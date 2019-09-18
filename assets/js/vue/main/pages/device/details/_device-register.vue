<template>
  <div class="details__basic-info">
    <div class="title">
      <i class="el-icon-s-order"></i>
      <span>创建信息</span>
    </div>

    <div v-loading="$apollo.queries.device.loading" element-loading-background="unset">
      <div class="item">
        <span>创建人</span>
        <span
          v-if="device.register && device.register.userExtend && device.register.userExtend.name"
        >{{ device.register.userExtend.name }}</span>
        <span v-else>-</span>
      </div>

      <div class="item">
        <span>手机号</span>
        <span v-if="device.register && device.register.phone">{{ device.register.phone }}</span>
        <span v-else>-</span>
      </div>

      <div class="item">
        <span>邮箱</span>
        <span
          v-if="device.register && device.register.userExtend && device.register.userExtend.email"
        >{{ device.register.userExtend.email }}</span>
        <span v-else>-</span>
      </div>

      <div class="item">
        <span>创建日期</span>
        <span v-if="device.createdAt">{{ timeFormatter(device.createdAt) }}</span>
        <span v-else>-</span>
      </div>
    </div>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';
import deviceRegister from './gql/query.device-register.gql';

export default {
  name: 'details-register-info',
  props: ['uuid'],
  apollo: {
    device: {
      query: deviceRegister,
      variables() {
        return { uuid: this.uuid };
      }
    }
  },
  data() {
    return {
      device: {}
    };
  },
  methods: {
    timeFormatter(val) {
      return timeFormatter(val);
    }
  }
};
</script>
