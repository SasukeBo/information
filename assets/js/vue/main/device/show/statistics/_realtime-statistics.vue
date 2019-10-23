<template>
  <div class="realtime">
    <div ref="chart" class="chart"></div>
  </div>
</template>
<script>
import realtimeQuery from '../gql/query.realtime-statistics.gql';
import { timeFormatter } from 'js/utils';
import echarts from 'echarts';

export default {
  name: 'realtime',
  props: ['deviceID', 'product'],
  data() {
    return {
      limit: 50,
      chart: null,
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
          type: 'time',
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
          }
        },
        yAxis: {
          type: 'value',
          name: '检测值',
          nameGap: 30,
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
      },
      updater: undefined,
      items: {}
    };
  },
  methods: {
    initChart() {
      this.chart = echarts.init(this.$refs.chart);
      this.options.title.text = `${this.product.name}生产数据`;
      this.options.legend.data = this.product.detectItems.map(i => i.sign);
      var data = this.makeFakeData();
      this.product.detectItems.forEach(di => {
        this.items[di.sign] = {
          name: di.sign,
          type: 'line',
          data: data.slice()
        };
      });
    },
    renderChart(options) {
      this.chart.setOption(options);
    },
    fetchData() {
      var now = new Date();
      now.setSeconds(now.getSeconds() - 2);

      this.$apollo
        .query({
          query: realtimeQuery,
          variables: {
            deviceID: this.deviceID,
            productID: this.product.id,
            floatPrecision: 3,
            afterTime: now.toISOString(),
            limit: 1
          },
          fetchPolicy: 'network-only'
        })
        .then(({ data }) => {
          this.updateItems(data.itemsUpdate);
          this.updateOptions();
        })
        .catch(e => {
          console.log(e.message);
        });
    },
    updateOptions() {
      var series = [];
      Object.keys(this.items).forEach(k => {
        series.push(this.items[k]);
      });
      this.renderChart({ series });
    },
    updateItems(newItems) {
      newItems.forEach(i => {
        var old = this.items[i.sign];
        var data = this.formatData(i);
        // 去除重复点
        if (old.data.length && old.data[old.data.length - 1].name === data.name)
          return;
        if (old.data.length >= this.limit) {
          old.data.shift();
        }
        old.data.push(data);
      });
    },
    formatData(item) {
      return {
        name: timeFormatter(item.time, '%timestring'),
        value: [timeFormatter(item.time, '%y/%m/%d %timestring'), item.value]
      };
    },
    makeFakeData() {
      var now = new Date().toISOString();
      var fakeData = [];
      var value = 0;
      for (var i = this.limit; i > 0; i--) {
        var time = new Date(now);
        time.setSeconds(time.getSeconds() - i);
        fakeData.push(this.formatData({ time, value }));
      }
      return fakeData;
    },
    stopFetch() {
      clearInterval(this.updater);
    }
  },
  mounted() {
    this.initChart();
    this.renderChart(this.options);
    this.fetchData();
    // this.updater = setInterval(() => this.fetchData(), 1000);
  },
  beforeDestroy() {
    clearInterval(this.updater);
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
    height: 300px;
  }
}
</style>
