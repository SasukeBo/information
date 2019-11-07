<template>
  <div class="month-data">
    <div class="data-row">
      <span>运行时间:</span>
      <span>{{ statistics.runningTime || '-' }}</span>
    </div>

    <div class="data-row">
      <span>稼动率:</span>
      <span>{{ (parseFloat(statistics.activation) * 100).toFixed(2) + '%' || '-' }}</span>
    </div>

    <div class="data-row">
      <span>产量:</span>
      <span>{{ statistics.yield + ' 个' || '-' }}</span>
    </div>

    <div class="data-row">
      <span>良率:</span>
      <span>{{ (parseFloat(statistics.yieldRate) * 100).toFixed(2) + '%' || '-' }}</span>
    </div>
  </div>
</template>
<script>
import statisticsQuery from '../gql/query.device-monthly-analysis.gql';

export default {
  name: 'monthly-analysis',
  props: ['id'],
  apollo: {
    statistics: {
      query: statisticsQuery,
      variables() {
        return { deviceID: this.id, format: '%D天 %H小时 %M分钟' };
      }
    }
  },
  data() {
    return {
      statistics: {}
    };
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-statistics {
  .month-data {
    color: $--color-font__gray;
  }

  .month-data .data-row {
    line-height: 26px;

    span:first-child {
      width: 100px;
      display: inline-block;
    }

    span:last-child {
      color: $--color-font__white;
    }
  }
}
</style>
