<template>
  <div class="global-card params-realtime-chart">
    <div ref="realtimeChart" style="width: 100%; height: 500px"></div>
  </div>
</template>
<script>
import echarts from 'echarts';
import { chartApollo } from './graphql';
import { mapState } from 'vuex';
import 'echarts/lib/chart/line';

export default {
  name: 'device-details-params-value-chart',
  props: ['device'],
  apollo: chartApollo,
  data() {
    return {
      params: [],
      chart: null
    };
  },
  computed: {
    ...mapState({
      deviceChannel: state => state.socket.deviceChannel
    })
  },
  mounted() {
    var data = [];
    var option = {};

    this.chart = echarts.init(this.$refs.realtimeChart);
    option = {
      title: {
        text: '设备参数值实时波形图'
      },
      tooltip: {},
      legend: {
        data: []
      },
      xAxis: {
        type: 'time',
        splitLine: {
          show: false
        }
      },
      yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
          show: false
        }
      },
      series: [
        {
          name: '模拟数据',
          type: 'line',
          showSymbol: true,
          data: data
        }
      ]
    };
    this.chart.setOption(option);

    if (!this.deviceChannel.topics.indexOf(`device_${this.device.id}`) > -1)
      this.deviceChannel.Join(`device_${this.device.id}`);

    this.deviceChannel.onData = ({ payload }) => {
      var time = new Date(payload._TIME_STAMP_);
      if (data.length > 99) data.shift();
      var h = time.getHours() < 10 ? `0${time.getHours()}` : `${time.getHours()}`;
      var m = time.getMinutes() < 10 ? `0${time.getMinutes()}` : `${time.getMinutes()}`;
      var s = time.getSeconds() < 10 ? `0${time.getSeconds()}` : `${time.getSeconds()}`;
      data.push({
        name: time.toString(),
        value: [
          `${time.toLocaleDateString()} ${h}:${m}:${s}`,
          payload.count
        ]
      });
      this.chart.setOption({ series: [{ data: data }] });
    };
  }
};
</script>
