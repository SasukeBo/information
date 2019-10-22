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
      chart: null,
      options: {
        title: {
          top: 20,
          left: 20,
          textStyle: {
            color: '#c0c4cc'
          }
        },
        legend: {
          top: 20,
          textStyle: {
            color: '#03a9f4'
          }
        },
        tooltip: {
          axisPointer: {
            animation: false
          }
        },
        xAxis: {
          type: 'time',
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
      this.product.detectItems.forEach(di => {
        this.items[di.sign] = { name: di.sign, type: 'line', data: [] };
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
        if (old.data.length >= 50) {
          old.data.shift();
          old.data.push(this.formatData(i));
        } else {
          old.data.push(this.formatData(i));
        }
      });
    },
    formatData(item) {
      return {
        name: timeFormatter(item.time, '%timestring'),
        value: [timeFormatter(item.time, '%y/%m/%d %timestring'), item.value]
      };
    }
  },
  mounted() {
    this.initChart();
    this.renderChart(this.options);
    this.updater = setInterval(() => this.fetchData(), 1000);
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
