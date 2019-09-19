<template>
  <div
    class="values-table"
    v-loading="$apollo.queries.params.loading"
    element-loading-background="unset"
  >
    <div class="table__inner">
      <div class="table__header">
        <span class="col-name" v-for="(colName, i) in tableHeader" :key="'colname_' + i">{{colName}}</span>
      </div>
      <div class="table__body">
        <div
          :class="['table-row', i % 2 === 0 ? 'odd-row' : 'even-row']"
          v-for="(row, i) in tableData"
          :key="'row_' + i"
        >
          <span class="col" v-for="(col, j) in row" :key="'col_' + i + j">{{col}}</span>
        </div>
      </div>
    </div>

    <el-pagination
      class="table__pagination"
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="currentPage"
      :page-sizes="[100, 200, 300, 400]"
      :page-size="limit"
      layout="total, sizes, prev, pager, next, jumper"
      :total="count"
    ></el-pagination>
  </div>
</template>
<script>
import paramWithValues from './gql/query.param-values.gql';
import paramValueCount from './gql/query.param-value-count.gql';
import { Pagination } from 'element-ui';
import { timeFormatter } from 'js/utils';

export default {
  /*
    TODO: 列表显示采用分页
  */
  name: 'values-table',
  props: ['uuid', 'afterTime', 'beforeTime'],
  components: { ElPagination: Pagination },
  apollo: {
    params: {
      query: paramWithValues,
      variables() {
        return {
          limit: this.limit,
          offset: this.offset,
          deviceUUID: this.uuid,
          beforeTime: this.beforeTime,
          afterTime: this.afterTime
        };
      }
    },
    count: {
      query: paramValueCount,
      variables() {
        return {
          deviceUUID: this.uuid,
          beforeTime: this.beforeTime,
          afterTime: this.afterTime
        };
      }
    }
  },
  data() {
    return {
      params: [],
      count: 0,
      limit: 100,
      currentPage: 1
    };
  },
  computed: {
    offset() {
      return (this.currentPage - 1) * this.limit;
    },
    tableHeader() {
      if (!this.params.length) return;
      var tableHeader = ['时间'];
      for (var i = 0; i < this.params.length; i++) {
        tableHeader.push(this.params[i].name);
      }
      return tableHeader;
    },
    tableData() {
      var col = this.params.length;
      if (!col) return [];
      var row = this.params[0].values.length;
      var tableData = [];
      for (var i = 0; i < row; i++) {
        var rowData = [];
        rowData.push(timeFormatter(this.params[0].values[i].createdAt));
        for (var j = 0; j < col; j++) {
          rowData.push(this.params[j].values[i].value);
        }
        tableData.push(rowData);
      }
      return tableData;
    }
  },
  methods: {
    handleSizeChange(val) {
      this.limit = val;
      this.currentPage = 1;
    },
    handleCurrentChange(val) {
      this.currentPage = val;
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-details .values-table {
  height: 100%;
  width: 100%;

  .table__inner {
    height: calc(100% - 50px);
    overflow-x: scroll;
  }

  .table__pagination {
    text-align: center;
    padding-top: 18px;
  }

  .table__inner .table__header {
    position: fixed;
    font-size: 1.2rem;
    width: 100%;
    display: flex;
    background: $--color-theme__bg;
    padding: 0.5rem 0;
    box-shadow: $--shadow__global-card;

    .col-name {
      flex: auto;
      min-width: 100px;
      text-align: center;
    }
  }

  .table__inner .table__body {
    margin-top: 43px;

    .table-row {
      padding: 0.2rem 0;
      display: flex;

      &.odd-row {
        background: $--color-theme__light;
      }

      &.even-row {
        background: $--color-theme__bg;
      }
    }

    .col {
      flex: 1;
      text-align: center;
      min-width: 100px;
    }
  }
}
</style>
