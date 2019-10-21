<template>
  <div
    class="realtime"
    v-loading="$apollo.queries.statistics.loading"
    element-loading-background="unset"
  >
    <div ref="chart" style="width: 100%; height: 250px"></div>
  </div>
</template>
<script>
import realtimeQuery from '../gql/query.realtime-statistics.gql';
import { timeFormatter } from 'js/utils';
import echarts from 'echarts';

export default {
  name: 'realtime',
  props: ['deviceID', 'product'],
  apollo: {
    statistics: {
      query: realtimeQuery,
      variables() {
        return {
          deviceID: this.deviceID,
          productID: this.product.id,
          limit: 100,
          afterTime: this.afterTime
        };
      }
    }
  },
  data() {
    return {
      afterTime: undefined,
      statistics: [],
      chart: null,
      options: {
        title: {},
        legend: {},
        xAxis: {
          type: 'category',
          boundaryGap: false,
          axisLine: {
            lineStyle: { color: '#c0c4cc' }
          }
        },
        yAxis: {
          type: 'value',
          axisLine: {
            lineStyle: { color: '#c0c4cc' }
          }
        },
        series: []
      },
      items: {
        datas: {},
        times: {}
      }
    };
  },
  watch: {
    product(p) {
      if (p) {
        this.options.title.text = `${p.name}生产数据`;
        this.options.legend.data = p.detectItems.map(i => i.sign);
        this.options.series = p.detectItems.map(i => {
          return {
            name: i.sign,
            type: 'line',
            data: []
          };
        });
      }
    },
    statistics(newVal) {
      if (newVal.length) {
        this.formatStatistics(newVal);
        this.refreshChart();
      }
    }
  },
  methods: {
    refreshChart() {
      var optionUpdate = {
        xAxis: { data: this.items.times[this.product.detectItems[0].sign] },
        series: []
      };
      this.product.detectItems.forEach(i => {
        optionUpdate.series.push({
          name: i.sign,
          type: 'line',
          data: this.items.datas[i.sign]
        });
      });
      this.chart.setOption(optionUpdate);
    },
    initChart() {
      this.chart = echarts.init(this.$refs.chart);
      this.chart.setOption(this.options);
    },
    formatStatistics(items) {
      var _this = this;
      items.forEach(item => {
        var seriesData = _this.items.datas[item.sign] || [];
        var seriesTime = _this.items.times[item.sign] || [];
        seriesData.push(item.value.toFixed(3)); // 存入数据
        seriesTime.push(timeFormatter(item.createdAt, '%timestring')); // 存入时间
        _this.items.datas[item.sign] = seriesData;
        _this.items.times[item.sign] = seriesTime;
      });
    },
    changeTime() {
      this.afterTime = this.statistics[0].createdAt;
    }
  },
  mounted() {
    this.initChart();
    // setInterval(() => this.changeTime(), 1000);
  }
};
</script>
