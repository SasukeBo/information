<template>
  <div class="device-status-log">
    <div class="block" v-if="false">
      <div class="field-title">停机原因统计</div>
      <type-chart :id="id"></type-chart>
    </div>

    <div class="block">
      <div class="field-title">
        <span class="label">日志列表</span>
        <el-date-picker
          :editable="false"
          v-model="date"
          type="date"
          size="small"
          placeholder="选择日期"
        ></el-date-picker>
      </div>

      <div class="log-table">
        <el-table
          row-class-name="log-table-row"
          @expand-change="setExpanded"
          cell-class-name="log-table-cell"
          header-row-class-name="log-table-header"
          header-cell-class-name="log-table-header-cell"
          :highlight-current-row="false"
          ref="table"
          :data="logList.logs"
          style="width: 100%"
        >
          <el-table-column type="expand">
            <template slot-scope="props">
              <stop-reasons :id="props.row.id"></stop-reasons>
            </template>
          </el-table-column>

          <el-table-column label="日志ID" width="100px" prop="id"></el-table-column>

          <el-table-column label="停机开始时间">
            <template slot-scope="scope">
              <span>{{ format(scope.row.beginAt) }}</span>
            </template>
          </el-table-column>

          <el-table-column label="停机结束时间">
            <template slot-scope="scope">
              <span>{{ format(scope.row.finishAt) }}</span>
            </template>
          </el-table-column>

          <el-table-column>
            <template slot="header">
              <el-button type="info" @click="closeAllExpanded">收起所有</el-button>
            </template>

            <template slot-scope="scope">
              <el-button type="primary" @click="toggleExpand(scope.row)">点击查看原因</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <div class="log-pagination" v-if="logList.total">
        <el-pagination
          background
          layout="prev, pager, next"
          :page-size="limit"
          :total="logList.total"
          :current-page.sync="currentPageIndex"
        ></el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
import typeChart from './stop-type-count-chart';
import stopReasons from './stop-reason';

import logListQuery from '../gql/query.stop-log-list.gql';
import { timeFormatter } from 'js/utils';

export default {
  name: 'device-status-log',
  props: ['id'],
  components: { typeChart, stopReasons },
  apollo: {
    logList: {
      query: logListQuery,
      variables() {
        return {
          deviceID: this.id,
          offset: this.offset,
          beginTime: this.beginTime,
          endTime: this.endTime,
          limit: this.limit
        };
      }
    }
  },
  data() {
    return {
      date: new Date(),
      limit: 50,
      expanded: [],
      currentPageIndex: 1,
      logList: {
        total: 0,
        logs: []
      }
    };
  },
  computed: {
    offset() {
      return this.limit * (this.currentPageIndex - 1);
    },
    beginTime() {
      if (this.date === '') return '0001-01-01T00:00:00Z';
      var time = new Date(this.date);
      time.setMilliseconds(0);
      time.setSeconds(0);
      time.setMinutes(0);
      time.setHours(0);
      return time.toISOString();
    },
    endTime() {
      if (this.date === '') return '0001-01-01T00:00:00Z';
      var now = new Date();
      var time = new Date(this.date);
      if (now.toDateString() === time.toDateString()) {
        return now;
      }

      time.setMilliseconds(0);
      time.setSeconds(0);
      time.setMinutes(0);
      time.setHours(0);
      time.setDate(time.getDate() + 1);
      return time.toISOString();
    }
  },
  methods: {
    closeAllExpanded() {
      var expanded = this.expanded.slice(0, this.expanded.length);
      expanded.forEach(row => this.$refs.table.toggleRowExpansion(row, false));
    },
    format(time) {
      return timeFormatter(time);
    },
    toggleExpand(row) {
      this.$refs.table.toggleRowExpansion(row);
    },
    setExpanded(_, expanded) {
      this.expanded = expanded;
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

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

  .field-title .label {
    display: inline-block;
    vertical-align: middle;
    padding-right: 2rem;
  }

  .block .log-pagination {
    text-align: center;
  }

  .log-table {
    margin-bottom: 1rem;
  }

  .log-table .el-table {
    background: $--color-theme__bg;

    td {
      border-bottom: 1px solid $--color-theme__light;
    }

    .el-table__empty-block {
      height: unset;
    }
  }

  .log-table .el-table .log-table-row {
    background: $--color-theme__bg;

    &:hover td {
      background: $--color-theme__bg;
    }
  }

  .log-table .el-table .el-table__expanded-cell {
    background: lighten($--color-theme__bg, 5%);
    border-top: 1px solid $--color-theme__gray;
    border-bottom: 1px solid $--color-theme__gray;

    &:hover {
      background: $--color-theme__bg !important;
    }
  }

  .log-table .el-table .log-table-cell {
    text-align: center;
    border-bottom: none;
    color: $--color-font__silver;
  }

  .log-table .el-table .log-table-header {
    background: $--color-theme__bg;
  }

  .log-table .el-table .log-table-header-cell {
    text-align: center;
    background: $--color-theme__bg;
    color: $--color-font__light;
  }
}
</style>
