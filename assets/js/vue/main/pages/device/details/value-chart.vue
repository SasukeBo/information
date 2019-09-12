<template>
  <div class="global-card params-realtime-chart">
    <div ref="realtimeChart" style="width: 100%; height: 500px"></div>
    <div v-if="series.length">
      <value-sub
        v-for="(p, i) in params"
        :key="p.id"
        :param="p"
        v-bind:seriesData.sync="series[i].data"
      ></value-sub>
    </div>
  </div>
</template>
<script>
import echarts from 'echarts';
import { timeFormatter } from 'js/utils';
import 'echarts/lib/chart/line';
import ValueSub from './_param-value-sub';

export default {
  name: 'value-chart',
  props: ['params'],
  components: { ValueSub },
  data() {
    return {
      chart: null,
      chartStop: false,
      legendData: [],
      series: [],
      option: {},
      refreshInterval: null
    };
  },
  mounted() {
    this.chart = echarts.init(this.$refs.realtimeChart);
    this.series = this.params.map(p => {
      return {
        name: p.name,
        type: 'line',
        // symbol: 'emptyCircle',
        // symbolSize: 4,
        showSymbol: false,
        data: [],
        hoverAnimation: false
      };
    });

    this.legendData = this.params.map(p => {
      return {
        name: p.name,
        icon: 'roundRect'
      };
    });

    this.option = {
      title: {
        text: '设备参数实时数据',
        textStyle: {
          color: '#dcdfe6',
          fontSize: 20,
          lineHeight: 30
        },
        left: 'center'
      },
      legend: {
        top: '5%',
        zlevel: 1,
        left: '85%',
        orient: 'vertical',
        inactiveColor: '#909399',
        textStyle: {
          color: '#dcdfe6'
        },
        data: this.legendData,
      },
      grid: {
        left: '10%',
        right: '15%'
      },
      tooltip: {
        trigger: 'axis',
        axisPointer: {
          type: 'cross',
          animation: false
        },
        backgroundColor: 'rgba(245, 245, 245, 0.8)',
        borderWidth: 1,
        borderColor: '#ccc',
        padding: 10,
        textStyle: {
          color: '#000'
        }
      },
      axisPointer: {
        link: { xAxisIndex: 'all' },
        label: {
          backgroundColor: '#777'
        }
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
      series: this.series
    };
    this.chart.setOption(this.option);

    this.refreshInterval = setInterval(() => {
      this.chart.setOption({ series: this.series });
    }, 1000);
  }
};
</script>
