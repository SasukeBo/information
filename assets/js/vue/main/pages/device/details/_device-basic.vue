<template>
  <div class="details__basic-info">
    <div class="title">
      <i class="el-icon-document"></i>
      <span>基本信息</span>
    </div>

    <div v-loading="$apollo.queries.deviceBasic.loading" element-loading-background="unset">
      <div class="item">
        <span>名称</span>
        <span>{{ deviceBasic.name }}</span>
      </div>

      <div class="item">
        <span>设备类型</span>
        <span>{{ deviceBasic.type }}</span>
      </div>

      <div class="item">
        <span>Token</span>
        <span>{{ deviceBasic.token }}</span>
      </div>

      <div class="item">
        <span>地址</span>
        <span>{{ deviceBasic.address }}</span>
      </div>

      <div class="item">
        <span>描述</span>
        <span>{{ deviceBasic.description }}</span>
      </div>

    </div>
  </div>
</template>
<script>
import deviceBasicQuery from './gql/query.device-basic.gql';

export default {
  name: 'details-basic-info',
  props: ['uuid'],
  apollo: {
    deviceBasic: {
      query: deviceBasicQuery,
      variables() {
        return { uuid: this.uuid };
      }
    }
  },
  data() {
    return {
      deviceBasic: {}
    };
  }
};
</script>
