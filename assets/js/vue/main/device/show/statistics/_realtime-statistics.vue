<template>
  <div class="realtime">
    <div
      ref="chart"
      class="chart"
      v-loading="$apollo.queries.chartData.loading"
      element-loading-background="unset"
    ></div>
  </div>
</template>
<script>
import chartInitQuery from '../gql/query.get-detectitem-chart-data.gql';
import { timeFormatter } from 'js/utils';
import echarts from 'echarts';
import gql from 'graphql-tag';

export default {
  name: 'realtime',
  props: ['deviceID', 'product'],
  apollo: {
    chartData: {
      query: chartInitQuery,
      variables() {
        return {
          deviceID: this.deviceID,
          productID: this.product.id,
          limit: 50
        };
      }
    },
    $subscribe: {
      productIns: {
        query: gql`
          subscription productIns($deviceID: Int!, $id: Int!) {
            productIns: productInsAdd(id: $id, deviceID: $deviceID) {
              createdAt
              detectItemValues {
                detectItem {
                  sign
                }
                value
              }
            }
          }
        `,
        variables() {
          return {
            deviceID: this.deviceID
          };
        },
        result({ data: { productIns } }) {
          this.options.xAxis.data.shift();
          this.options.xAxis.data.push(this.formatTime(productIns.createdAt));
          productIns.detectItemValues.forEach(v =>
            this.options.series.forEach(s => {
              if (s.name === v.detectItem.sign) {
                s.data.shift();
                s.data.push(v.value.toFixed(3));
              }
            })
          );
          this.chart.setOption(this.options);
        }
      }
    }
  },
  data() {
    return {
      limit: 50,
      chart: null,
      chartData: {},
      options: {
        title: {
          top: 10,
          left: 20,
          textStyle: {
            color: '#fff'
          }
        },
        color: [
          '#03A9F4', // blue
          '#8FC860', // green
          '#F9A230', // orange
          '#F06D6B', // red
          '#ACA0F2', // purple
          '#C0C4CC', // white
          '#000' // black
        ],
        grid: { show: false },
        legend: {
          top: 20,
          type: 'scroll',
          right: 20,
          orient: 'vertical',
          textStyle: {
            color: '#fff'
          }
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            animation: false
          }
        },
        axisPointer: {
          label: {
            backgroundColor: '#03a9f4'
          }
        },
        xAxis: {
          type: 'category',
          name: '生产时间',
          nameLocation: 'center',
          nameGap: 30,
          nameTextStyle: { color: '#fff' },
          boundaryGap: false,
          axisLabel: { color: '#a5bbef' },
          axisLine: {
            symbol: ['none', 'arrow'],
            symbolSize: [5, 10],
            lineStyle: { color: '#909399' }
          },
          splitLine: {
            show: true,
            lineStyle: { color: '#666' }
          },
          data: []
        },
        yAxis: {
          type: 'value',
          name: '检测值',
          nameGap: 50,
          nameLocation: 'center',
          nameTextStyle: { color: '#fff' },
          axisLabel: { color: '#a5bbef' },
          axisLine: {
            symbol: ['none', 'arrow'],
            symbolSize: [5, 10],
            lineStyle: { color: '#909399' }
          },
          splitLine: {
            show: true,
            lineStyle: { color: '#666' }
          }
        },
        series: []
      }
    };
  },
  watch: {
    chartData(newVal) {
      this.options.title.text = `${this.product.name}生产数据`;
      this.options.legend.data = this.product.detectItems.map(i => i.sign);
      this.options.xAxis.data = newVal.timestamps.reverse().map(t => this.formatTime(t));
      var series = newVal.items.map(item => {
        return {
          type: 'line',
          name: item.sign,
	  symbol: 'none',
          data: item.values.reverse().map(v => v.value.toFixed(3))
        };
      });
      this.options.series = series;
      this.chart.setOption(this.options);
    }
  },
  methods: {
    formatTime(timeString) {
      var time = new Date(timeString);
      return `${time.toTimeString().slice(0, 8)}+${time.getMilliseconds()}ms`;
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.chart);
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';
.device-statistics .realtime {
  padding: 1rem 0;

  .chart {
    border: 1px solid $--color-border__0;
    box-shadow: $--shadow__global-card;
    width: 100%;
    height: 400px;
  }
}
</style>
