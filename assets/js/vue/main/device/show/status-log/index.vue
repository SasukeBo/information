<template>
  <div class="device-status-log">
    <div class="block">
      <div class="field-title">停机原因统计</div>
      <type-chart :id="id"></type-chart>
    </div>

    <div class="block">
      <div class="field-title">日志列表</div>
    </div>
  </div>
</template>
<script>
import typeChart from './stop-type-count-chart';

export default {
  name: 'device-status-log',
  props: ['id'],
  components: { typeChart }
};
</script>
<style lang="scss">
.device-status-log {
  .block {
    position: relative;
    margin-bottom: 2rem;
  }

  .field-title {
    font-size: 24px;
    line-height: 24px;
    margin-bottom: 25px;
    font-weight: 400;
  }
}
</style>