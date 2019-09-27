<template>
  <div class="status-log__table">
    <el-table
      :data="tableData"
      height="468"
      style="width: 100%"
      header-cell-class-name="custom-header-cell"
      cell-class-name="custom-cell"
    >
      <el-table-column type="index" width="50"></el-table-column>
      <el-table-column
        prop="status"
        label="设备状态"
        :filters="[{text: '生产', value: 'prod'}, {text: '停机', value: 'stop'}, {text: '离线', value: 'offline'}]"
        :filter-method="filterStatus"
      >
        <template slot-scope="scope">
          <span :class="'status_' + scope.row.status">{{ statusMap[scope.row.status] }}</span>
        </template>
      </el-table-column>
      <el-table-column prop="duration" label="持续时间">
        <template slot-scope="scope">{{ formatDuration(scope.row.duration) }}</template>
      </el-table-column>
      <el-table-column prop="createdAt" label="状态变更时间">
        <template slot-scope="scope">{{ formatTime(scope.row.createdAt)}}</template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import { Table, TableColumn } from 'element-ui';
export default {
  name: 'status-log-table',
  props: ['logs'],
  components: {
    ElTable: Table,
    ElTableColumn: TableColumn
  },
  data() {
    return {
      statusMap: {
        stop: '停机',
        prod: '生产',
        offline: '离线'
      }
    };
  },
  computed: {
    tableData() {
      var logs = new Array(...this.logs);
      return logs.reverse();
    }
  },
  methods: {
    formatDuration(duration) {
      var h = Math.floor(duration / 3600);
      var m = Math.floor((duration % 3600) / 60);
      var s = duration % 60;

      return `${h}小时${m}分${s}秒`;
    },
    formatTime(timeStr) {
      if (timeStr === '0001-01-01T00:00:00Z') return '至现在';
      var time = new Date(timeStr);
      return time.toLocaleString();
    },
    filterStatus(value, row) {
      return row.status === value;
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.status-log__table {
  flex: 1;
  min-width: 500px;
  padding: 0 1rem;
  height: 500px;

  .el-table {
    background: $--color-theme__bg;

    &:before {
      display: none;
    }

    th.is-leaf,
    td {
      border-bottom: 1px solid rgba(255, 255, 255, 0.1);
      box-shadow: 1px 0 2px $--color-theme__white;
    }
  }

  .custom-cell {
    color: $--color-font__gray;
    font-size: 0.875rem;
    padding: 0.5rem 0;
    background: $--color-theme__bg;
    border-color: $--color-theme__bg;
  }

  .custom-header-cell {
    font-weight: bold;
    color: $--color-font__silver;
    background: $--color-theme__bg;
    border-color: $--color-theme__bg;
  }

  .el-table__column-filter-trigger {
    background: $--color-theme__main;
    margin-left: 0.5rem;
    border-radius: 4px;

    i {
      color: #fff;
      padding: 0.1rem;
      font-size: 0.9rem;
      font-weight: bold;
    }
  }

  .status_prod {
    color: $--color-theme__success;
  }

  .status_stop {
    color: $--color-theme__warning;
  }

  .status_offline {
    color: $--color-theme__danger;
  }
}
</style>
