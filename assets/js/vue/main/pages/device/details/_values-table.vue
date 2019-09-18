<template>
  <div class="values-table">
    <div v-for="(row, i) in tableData" :key="'row_' + i">
      <span v-for="(col, j) in row" :key="'col_' + i + j">{{col}}</span>
    </div>
  </div>
</template>
<script>
export default {
  /*
    TODO: 列表显示采用分页
  */
  name: 'values-table',
  props: ['params'],
  data() {
    return {
      tableData: []
    };
  },
  created() {
    var col = this.params.length;
    var row = this.params[0].values.length;
    var tableData = [];
    var tableHeader = [];
    for (var j = 0; j < col; j++) {
      tableHeader.push(this.params[j].name);
    }
    tableData.push(tableHeader);
    for (var i = 0; i < row; i++) {
      var rowData = [];
      for (var j = 0; j < col; j++) {
        rowData.push(this.params[j].values[i].value);
      }
      tableData.push(rowData);
    }
    this.tableData = tableData;
  }
};
</script>
