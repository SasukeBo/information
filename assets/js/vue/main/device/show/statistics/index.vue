<template>
  <div class="device-statistics">
    <div class="block">
      <span class="field-title">月数据统计</span>
      <month-data></month-data>
    </div>
  </div>
</template>
<script>
import MonthData from './_month-data';

export default {
  name: 'device-statistics',
  components: { MonthData },
  props: ['id']
};
</script>
<style lang="scss">
@import 'css/vars.scss';
</style>
