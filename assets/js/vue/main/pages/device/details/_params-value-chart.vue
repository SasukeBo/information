<template>
  <div class="global-card params-realtime-chart">
    <div ref="realtimeChart" style="width: 100%; height: 500px"></div>
  </div>
</template>
<script>
import echarts from 'echarts';
import { chartApollo } from './graphql';
import 'echarts/lib/chart/line';

export default {
  name: 'device-details-params-value-chart',
  props: ['uuid'],
  apollo: chartApollo,
  data() {
    return {
      params: []
    }
  },
  mounted() {
    var el = this.$refs['realtimeChart'];
    var myChart = echarts.init(el);
    var params = [];
    // this.params.forEach(e => data.push(e.name))

    var option = {
      title: {
        text: '设备参数值实时波形图'
      },
      tooltip: {},
      legend: {
        data: params
      },
      xAxis: {
        data: ['衬衫', '羊毛衫', '雪纺衫', '裤子', '高跟鞋', '袜子']
      },
      yAxis: {},
      series: [
        {
          name: '销量',
          type: 'line',
          data: [5, 20, 36, 10, 10, 20]
        }
      ]
    };
    myChart.setOption(option);
  }
};
</script>
