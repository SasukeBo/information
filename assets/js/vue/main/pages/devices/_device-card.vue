<template>
  <div class="device-card" @click="pushRouter">
    <div class="global-card">
      <div class="data-item device-name">{{ device.name }}</div>

      <div class="data-item device-status" :class="['device-' + device.status]">
        <div class="label-center">当前状态</div>
        <i class="iconfont icon-production"></i>
        <div>{{ status }}</div>
      </div>

      <div class="data-item">
        <span class="label">设备类型</span>
        {{device.type}}
      </div>

      <div class="data-item">
        <span class="label">设备地址</span>
        {{device.address}}
      </div>

      <div class="data-item">
        <span class="label">今日良率</span>
        {{'-'}}
      </div>

      <div class="data-item">
        <span class="label">今日稼动率</span>
        {{'-'}}
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
import { timeFormatter } from 'js/utils';

export default {
  name: 'device-card',
  props: ['item'],
  data() {
    return {
      device: this.item,
      statusMap: {
        prod: '生产中',
        stop: '停机',
        offline: '离线'
      }
    };
  },
  computed: {
    status() {
      return this.statusMap[this.device.status];
    }
  },
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    },
    pushRouter() {
      NProgress.start();
      this.$router.push({
        name: 'device-show',
        params: { uuid: this.device.uuid }
      });
    }
  }
};
</script>
