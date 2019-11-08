<template>
  <div
    class="stop-type-count-chart"
    element-loading-background="unset"
    v-loading="$apollo.queries.stopTypeCount.loading"
  >
    <div ref="chart" class="chart"></div>
  </div>
</template>
<script>
import stopTypeCountQuery from '../gql/query.stop-type-count.gql';
import echarts from 'echarts';

export default {
  name: 'stop-type-count-chart',
  props: ['id'],
  apollo: {
    stopTypeCount: {
      query: stopTypeCountQuery,
      variables() {
        var now = new Date();
        var amago = new Date();
        amago.setMonth(amago.getMonth() - 1);
        return {
          deviceID: this.id,
          beginTime: amago,
          endTime: now
        };
      }
    }
  },
  data() {
    return {
      stopTypeCount: null,
      chart: null,
      options: {
        legend: {
          top: 50,
          type: 'scroll',
          right: 20,
          orient: 'vertical',
          textStyle: {
            color: '#fff'
          }
        },
        tooltip: {
          trigger: 'axis',
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
            name: '时间',
            nameLocation: 'center',
            nameGap: 30,
            nameTextStyle: { color: '#fff' },
            axisLabel: { color: '#a5bbef' },
            axisLine: {
              lineStyle: { color: '#909399' }
            }
          }
        ],
        yAxis: [
          {
            type: 'value',
            name: '停机次数',
            nameTextStyle: { color: '#fff' },
            axisLabel: { color: '#a5bbef' },
            axisLine: {
              lineStyle: { color: '#909399' }
            },
            splitLine: {
              show: true,
              lineStyle: { color: '#666' }
            }
          }
        ],
        series: []
      }
    };
  },
  watch: {
    stopTypeCount(newVal) {
      if (newVal) {
        this.options.legend.data = newVal.types;
        this.options.xAxis[0].data = newVal.days;
        newVal.counts.forEach(i => {
          this.options.series.push({
            name: i.name,
            smooth: true,
            type: 'bar',
            data: i.numbers
          });
        });
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

.stop-type-count-chart {
  padding: 1rem 0;

  .chart {
    border: 1px solid $--color-border__0;
    box-shadow: $--shadow__global-card;
    width: 100%;
    height: 300px;
  }
}
</style>
