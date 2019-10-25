<template>
  <div
    class="device-status-chart"
    v-loading="$apollo.queries.statistics.loading"
    element-loading-background="unset"
  >
    <div ref="chart" class="chart"></div>
  </div>
</template>
<script>
import statusDailyDurationQuery from '../gql/query.status-daily-duration.gql';
import echarts from 'echarts';

export default {
  name: 'device-status-chart',
  props: ['deviceID'],
  apollo: {
    statistics: {
      query: statusDailyDurationQuery,
      variables() {
        return {
          deviceID: this.deviceID,
          daysCount: 30
        };
      }
    }
  },
  data() {
    return {
      statistics: undefined,
      chart: undefined,
      options: {
        // color: ['#909399', '#409EFF', '#E6A23C'],
        title: {
          top: 10,
          left: 20,
          text: '最近一个月设备状态时间',
          textStyle: { color: '#fff' }
        },
        grid: {
          top: 80
        },
        tooltip: {
          trigger: 'axis',
          position: ['90%', 40],
          axisPointer: {
            type: 'shadow'
          }
        },
        legend: {
          top: 15,
          data: ['生产', '停机', '离线'],
          textStyle: { color: '#fff' }
        },
        toolbox: {
          show: true,
          feature: {
            dataView: { show: true, readOnly: true },
            saveAsImage: { show: true }
          },
          iconStyle: { borderColor: '#fff' }
        },
        xAxis: [
          {
            name: '日期',
            nameGap: 30,
            nameTextStyle: { color: '#fff' },
            type: 'category',
            data: [],
            axisLabel: { color: '#a5bbef', rotate: -45 },
            axisLine: { lineStyle: { color: '#c0c4cc' } }
          }
        ],
        yAxis: [
          {
            name: '小时',
            nameTextStyle: { color: '#fff' },
            type: 'value',
            axisLabel: { color: '#a5bbef' },
            axisLine: { lineStyle: { color: '#c0c4cc' } },
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
    statistics(newVal) {
      if (newVal) {
        this.options.xAxis[0].data = newVal.days;
        var stop = {
          name: '停机',
          type: 'bar',
          data: newVal.stop.map(i => parseFloat(i.toFixed(2)))
        };
        var prod = {
          name: '生产',
          type: 'bar',
          data: newVal.prod.map(i => parseFloat(i.toFixed(2)))
        };
        var offline = {
          name: '离线',
          type: 'bar',
          data: newVal.offline.map(i => parseFloat(i.toFixed(2)))
        };
        this.options.series = [stop, offline, prod];
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

.device-status-chart {
  padding: 1rem 0;

  .chart {
    border: 1px solid $--color-border__0;
    box-shadow: $--shadow__global-card;
    width: 100%;
    height: 300px;
  }
}
</style>
