<template>
  <div class="global-card params-realtime-chart">
    <div ref="realtimeChart" style="width: 100%; height: 500px"></div>
  </div>
</template>
<script>
import echarts from 'echarts';
import paramsQuery from './gql/query.params.gql';
import { mapState } from 'vuex';
import { timeFormatter } from 'js/utils';
import 'echarts/lib/chart/line';

export default {
  name: 'device-details-params-value-chart',
  props: ['device'],
  apollo: {
    params: {
      query: paramsQuery,
      variables() {
        return { deviceUUID: this.device.uuid };
      }
    }
  },
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
        text: '设备参数值实时波形图',
        textStyle: {
          color: '#dcdfe6',
          fontSize: 20,
          lineHeight: 30
        },
        left: 'center'
      },
      tooltip: {},
      legend: {
        data: []
      },
      xAxis: {
        type: 'time',
        splitLine: {
          show: false
        },
        axisLine: {
          lineStyle: {
            color: '#dcdf6e'
          }
        }
      },
      yAxis: {
        type: 'value',
        boundaryGap: [0, '100%'],
        splitLine: {
          show: false
        },
        axisLine: {
          lineStyle: {
            color: '#dcdf6e'
          }
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
      if (data.length === 100) data.shift();
      data.push({
        name: time.toString(),
        value: [timeFormatter(time, '%y/%m/%d %timestring'), payload.count]
      });
      this.chart.setOption({ series: [{ data: data }] });
    };
  }
};
</script>
