<template>
  <div style="text-align: center">
    <p>数据实时渲染</p>
    <div style="font-size: 20px; font-weight: bold">{{datalist}}</div>
  </div>
</template>
<script>
export default {
  name: 'home',
  data() {
    return {
      datalist: ''
    };
  },
  mounted() {
    var _this = this;
    var ws = new WebSocket(`ws://${document.location.hostname}/websocket`);

    ws.onopen = function() {
      ws.send('fakeData');
    };
    ws.onmessage = function({ data }) {
      _this.datalist = data;
    };
  }
};
</script>
