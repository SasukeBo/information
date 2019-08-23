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
      myChart: null
    };
  },
  computed: {
    ...mapState({
      deviceChannel: state => state.socket.deviceChannel
    })
  },
  mounted() {
    var el = this.$refs['realtimeChart'];
    this.myChart = echarts.init(el);
    // var data = [];
    // this.params.forEach(e => data.push(e.name))

    this.deviceChannel.Join(`device_${this.device.id}`);
    this.deviceChannel.onData = ({ payload }) => {
      var data = [{name: 'hello', value: ['2019/08/23 15:23', 100]}]
      var params = [];
      var time = new Date(payload._TIME_STAMP_);
      data.push({
        name: time.toLocaleString(),
        value: [time.toLocaleString(), payload.count]
      });
      var option = {
        title: {
          text: '设备参数值实时波形图'
        },
        tooltip: {},
        legend: {
          data: params
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
            showSymbol: false,
            // hoverAnimation: false,
            data: data
          }
        ]
      };
      console.log(option.series[0]);
      this.myChart.setOption(option);
    };
  }
};
</script>
