<template>
  <div class="global-card params-realtime-chart">
    <div ref="realtimeChart" style="width: 100%; height: 300px"></div>
  </div>
</template>
<script>
import echarts from 'echarts';
import { timeFormatter } from 'js/utils';
import 'echarts/lib/chart/line';
import valuesQuery from './gql/query.values.gql';
import valuesSub from './gql/sub.values.gql';

export default {
  name: 'device-details-params-value-chart',
  props: ['param'],
  apollo: {
    values: {
      query: valuesQuery,
      variables() {
        var time = new Date()
        time.setSeconds(time.getSeconds() + 100)
        return {
          paramID: this.param.id,
          limit: 100,
          before: time.toISOString()
        };
      },
      subscribeToMore: {
        document: valuesSub,
        variables() {
          return {
            topic: `device_param_value:${this.param.id}`
          };
        },
        updateQuery: (preData, { subscriptionData }) => {
          preData.values = [subscriptionData.data.values];
          return preData;
        }
      }
    }
  },
  data() {
    return {
      chart: null,
      values: [],
      seriesData: []
    };
  },
  watch: {
    values(newVal) {
      var newSeriesDate = this.formatValues(newVal);

      for (var i = newSeriesDate.length; i > 0; i--) {
        if (this.seriesData.length > 100) {
          this.seriesData.shift();
        }
        this.seriesData.push(newSeriesDate[i - 1]);
      }
      this.chart.setOption({ series: [{ data: this.seriesData }] });
    }
  },
  mounted() {
    var data = [];
    var option = {};
    this.seriesData = this.formatValues(this.values);

    this.chart = echarts.init(this.$refs.realtimeChart);
    var titleText = `参数 ${this.param.name} 实时数据`;
    option = {
      title: {
        text: titleText,
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
          data: this.seriesData
        }
      ]
    };
    this.chart.setOption(option);
  },
  methods: {
    formatValues(values) {
      return values.map(v => {
        var time = new Date(v.createdAt);
        return {
          name: timeFormatter(time, '%y年%m月%d日 %timestring'),
          value: [timeFormatter(time, '%y/%m/%d %timestring'), v.value]
        };
      });
    },
    initWS() {
      var ws = new WebSocket(`ws://${document.location.host}/websocket`);
      ws.onopen = function() {
        var data = { type: 'connection_init' };
        ws.send(JSON.stringify(data));
      };
      return ws;
    },
    push(ws, value) {
      var data = {
        type: 'data',
        payload: {
          variables: { topic: `device_param_value:${this.param.id}` },
          value: `${value}`,
          paramID: `${this.param.id}`
        }
      };

      ws.send(JSON.stringify(data));
    },
    start() {
      var ws = this.initWS();
      var interval = setInterval(() => {
        var i = Math.floor(Math.random() * 300 + 600);
        this.push(ws, i);
      }, 1000);
      return interval;
    }
  }
};
</script>
