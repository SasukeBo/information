<template>
  <div class="status-log__scatter">
    <div ref="chart" class="scatter-chart"></div>
  </div>
</template>
<script>
import echarts from 'echarts';

export default {
  name: 'status-log-scatter',
  props: ['logs'],
  data() {
    return {
      chart: undefined,
      option: {
        title: {
          text: '设备状态持续时间',
          textStyle: {
            color: '#dcdfe6',
            fontSize: 20,
            lineHeight: 30
          },
          left: 'center'
        },
        color: ['#c23531', '#ca8622', '#749f83'],
        legend: {
          top: '0',
          zlevel: 1,
          left: '0',
          orient: 'vertical',
          icon: 'roundRect',
          inactiveColor: '#909399',
          textStyle: {
            color: '#dcdfe6'
          }
        },
        tooltip: {
          trigger: 'axis',
          backgroundColor: 'rgba(245, 245, 245, 0.8)',
          borderWidth: 1,
          borderColor: '#ccc',
          padding: 10,
          textStyle: {
            color: '#000'
          }
        },
        xAxis: {
          type: 'category',
          axisLine: {
            lineStyle: {
              color: '#dcdf6e'
            }
          }
        },
        yAxis: {
          name: '单位/秒',
          gridIndex: 0,
          axisLine: {
            lineStyle: {
              color: '#dcdf6e'
            }
          }
        },
        grid: { top: '55%' },
        series: [
          { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          { type: 'line', smooth: true, seriesLayoutBy: 'row' },
          { type: 'line', smooth: true, seriesLayoutBy: 'row' }
        ]
      }
    };
  },
  computed: {
    dataset() {
      var source = [['duration'], ['离线'], ['停机'], ['生产']];
      var statusIndex = { offline: 1, stop: 2, prod: 3 };

      for (var i = 0; i < this.logs.length; i++) {
        var log = this.logs[i];
        var addIndex = statusIndex[log.status];
        if (log.createdAt === '0001-01-01T00:00:00Z')
          source[0].push(new Date().toLocaleString());
        else source[0].push(new Date(log.createdAt).toLocaleString());

        for (var j = 1; j < 4; j++) {
          var base = typeof source[j][i] === 'number' ? source[j][i] : 0;
          if (j === addIndex) source[j][i + 1] = base + log.duration;
          else source[j][i + 1] = base;
        }
      }
      return { source };
    }
  },
  watch: {
    dataset(value) {
      var firstCol = value.source[0][1];
      this.chart.setOption({ dataset: value });
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.chart);
    this.option.dataset = this.dataset;
    this.option.series.push({
      type: 'pie',
      id: 'pie',
      radius: '30%',
      center: ['50%', '30%'],
      label: {
        formatter: params => {
          var seconds = params.data[params.data.length - 1];
          var h = Math.floor(seconds / 3600);
          seconds = seconds - h * 3600;
          var m = Math.floor(seconds / 60);
          seconds = seconds - m * 60;

          return `${params.name}: ${h}时${m}分${seconds}秒 ${params.percent}%`;
        }
      },
      encode: {
        itemName: 'duration',
        value: this.dataset.source[0][this.dataset.source[0].length - 1],
        tooltip: this.dataset.source[0][this.dataset.source[0].length - 1]
      }
    });
    this.chart.on('updateAxisPointer', event => {
      var xAxisInfo = event.axesInfo[0];
      if (xAxisInfo) {
        var index = xAxisInfo.value + 2;
        this.chart.setOption({
          series: {
            id: 'pie',
            label: {
              formatter: params => {
                var seconds = params.data[index];
                var h = Math.floor(seconds / 3600);
                seconds = seconds - h * 3600;
                var m = Math.floor(seconds / 60);
                seconds = seconds - m * 60;

                return `${params.name}: ${h}时${m}分${seconds}秒 ${params.percent}%`;
              }
            },
            encode: {
              value: this.dataset.source[0][index],
              tooltip: this.dataset.source[0][index]
            }
          }
        });
      }
    });
    this.chart.setOption(this.option);
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.status-log__scatter {
  padding: 1rem;
  height: 500px;
  flex: 1;
  min-width: 500px;

  .scatter-chart {
    width: 100%;
    height: 100%;
  }
}
</style>
