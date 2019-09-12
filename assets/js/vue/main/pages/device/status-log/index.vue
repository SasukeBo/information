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
          :picker-options="pickerOptions"
          align="left"
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
      timeRange: [todayBegin, new Date()],
      pickerOptions: {
        shortcuts: [
          {
            text: '今天',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setHours(0, 0, 0, 0);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近一周',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit('pick', [start, end]);
            }
          }
        ]
      }
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
    margin: 0.5rem 1rem 0;
  }
}

.status-log-time-range {
  .el-button.el-picker-panel__link-btn.el-button--text.el-button--mini {
    display: none;
  }
}
</style>
