<template>
  <div
    class="histogram"
    v-loading="$apollo.queries.histogram.loading"
    element-loading-background="unset"
  >
    <div ref="chart" class="chart"></div>
  </div>
</template>
<script>
import histogramQuery from '../gql/query.histogram.gql';
import echarts from 'echarts';

export default {
  name: 'histogram',
  props: ['productID', 'deviceID', 'detectItem'],
  apollo: {
    histogram: {
      query: histogramQuery,
      variables() {
        return {
          id: this.productID,
          deviceID: this.deviceID,
          detectItemID: this.detectItem.id
        };
      }
    }
  },
  data() {
    return {
      histogram: undefined,
      colors: [
        '#03a9f4',
        '#8fc860',
        '#f9a230',
        '#f06d6b',
        '#000',
        '#aca0f2'
      ],
      options: {
        title: {
          top: 10,
          left: 20,
          textStyle: {
            color: '#fff'
          }
        },
        tooltip: {
          trigger: 'axis'
        },
        toolbox: {
          show: true,
          feature: {
            dataView: { show: true, readOnly: true },
            magicType: { show: true, type: ['line', 'bar'] },
            restore: { show: true },
            saveAsImage: { show: true }
          }
        },
        xAxis: [
          {
            type: 'category',
            name: '值区间',
            nameLocation: 'center',
            nameGap: 30,
            nameTextStyle: { color: '#fff' },
            axisLabel: { color: '#a5bbef' },
            axisLine: {
              symbol: ['none', 'arrow'],
              symbolSize: [5, 10],
              lineStyle: { color: '#909399' }
            }
          }
        ],
        yAxis: [
          {
            type: 'value',
            name: '产品数量',
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
          }
        ],
        series: [
          {
            name: '产品数量',
            type: 'bar',
            smooth: true,
            markPoint: {
              data: [
                { type: 'max', name: '最大值' },
                { type: 'min', name: '最小值' }
              ]
            },
            markLine: {
              data: [{ type: 'average', name: '平均值' }]
            }
          }
        ]
      }
    };
  },
  watch: {
    histogram(newVal) {
      if (newVal) {
        this.options.color = this.colors[this.detectItem.id % 6];
        this.options.xAxis[0].data = newVal.xAxisData;
        this.options.series[0].data = newVal.seriesData;
        this.options.title.text = `${this.detectItem.sign} - 直方图`;
        this.chart.setOption(this.options);
      }
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.chart);
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';
.device-statistics .histogram {
  padding: 1rem 0;

  .chart {
    border: 1px solid $--color-border__0;
    box-shadow: $--shadow__global-card;
    width: 100%;
    height: 300px;
  }
}
</style>
