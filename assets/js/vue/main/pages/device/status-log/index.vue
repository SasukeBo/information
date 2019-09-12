<template>
  <div class="status-log">
    <div class="global-card">
      <div class="time-range-picker">
        <el-date-picker
          v-model="timeRange"
          type="datetimerange"
          :clearable="false"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          popper-class="status-log-time-range"
        ></el-date-picker>
      </div>
      <div class="logs">
        <in-scatter :logs="logs" v-if="logs.length"></in-scatter>
        <in-table :logs="logs" v-if="logs.length"></in-table>
      </div>
    </div>
  </div>
</template>
<script>
import InScatter from './_scatter';
import InTable from './_table';
import logsQuery from './gql/query.status-logs.gql';
import { DatePicker } from 'element-ui';

export default {
  name: 'status-log',
  props: ['uuid'],
  components: {
    InScatter,
    InTable,
    ElDatePicker: DatePicker
  },
  apollo: {
    logs: {
      query: logsQuery,
      variables() {
        return {
          deviceUUID: this.uuid,
          afterTime: this.timeRange[0]
            ? this.timeRange[0].toISOString()
            : undefined,
          beforeTime: this.timeRange[1]
            ? this.timeRange[1].toISOString()
            : undefined
        };
      },
      fetchPolicy: 'network-only'
    }
  },
  data() {
    var todayBegin = new Date();
    todayBegin.setHours(0, 0, 0, 0);
    return {
      logs: [],
      timeRange: [todayBegin, new Date()]
    };
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.status-log {
  padding: 1rem;

  .logs {
    display: flex;
    flex-flow: wrap;
  }

  .time-range-picker {
    margin: 1rem;
  }

}

.status-log-time-range {
  .el-button.el-picker-panel__link-btn.el-button--text.el-button--mini {
    display: none;
  }
}
</style>
