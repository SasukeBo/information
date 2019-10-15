<template>
  <div class="details-histogram global-card">
    <div class="header">
      <el-date-picker
        v-model="timeDuration"
        size="mini"
        type="daterange"
        align="left"
        unlink-panels
        range-separator="至"
        start-placeholder="开始日期"
        end-placeholder="结束日期"
        :picker-options="pickerOptions"
      ></el-date-picker>
    </div>

    <div ref="histogramChart" style="height: 350px; width: 100%"></div>
  </div>
</template>
<script>
import queryHistogram from './gql/query.histogram.gql';
import queryParams from './gql/query.params.gql';
import echarts from 'echarts';
import 'echarts/lib/chart/line';

export default {
  name: 'details-histogram',
  props: ['uuid'],
  apollo: {
    histogram: {
      query: queryHistogram,
      variables() {
        return {
          paramID: this.paramID,
          beforeTime: this.beforeTime,
          afterTIme: this.afterTime
        };
      }
    },
    params: {
      query: queryParams,
      variables() {
        return {
          deviceUUID: this.uuid
        };
      }
    }
  },
  data() {
    return {
      chart: null,
      option: {},
      params: [],
      paramID: 1,
      histogram: {
        category: [],
        serie: []
      },
      timeDuration: [
        (() => {
          var time = new Date();
          time.setUTCHours(0, 0, 0, 0);
          return time;
        })(),
        new Date()
      ],
      pickerOptions: {
        shortcuts: [
          {
            text: '最近一周',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 7);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近一个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 30);
              picker.$emit('pick', [start, end]);
            }
          },
          {
            text: '最近三个月',
            onClick(picker) {
              const end = new Date();
              const start = new Date();
              start.setTime(start.getTime() - 3600 * 1000 * 24 * 90);
              picker.$emit('pick', [start, end]);
            }
          }
        ]
      }
    };
  },
  computed: {
    afterTime() {
      if (this.timeDuration.length > 0)
        return this.timeDuration[0].toISOString();
    },
    beforeTime() {
      if (this.timeDuration.length > 1)
        return this.timeDuration[1].toISOString();
    }
  },
  watch: {
    histogram(newVal) {
      this.chart.setOption({
        xAxis: { data: this.histogram.category },
        series: [{ data: newVal.serie }]
      });
    }
  },
  mounted() {
    this.chart = echarts.init(this.$refs.histogramChart);
    this.option = {
      title: {
        text: '参数直方图',
        textStyle: {
          color: '#dcdfe6',
          fontSize: 20,
          lineHeight: 30
        },
        left: 'center'
      },
      tooltip: {
        trigger: 'axis'
      },
      toolbox: {
        show: true,
        feature: {
          dataView: { show: true, readOnly: false },
          magicType: { show: true, type: ['line', 'bar'] },
          restore: { show: true },
          saveAsImage: { show: true }
        }
      },
      calculable: true,
      xAxis: {
        type: 'category',
        data: this.histogram.category,
        axisLine: {
          lineStyle: {
            color: '#dcdf6e'
          }
        }
      },
      yAxis: {
        type: 'value',
        axisLine: {
          lineStyle: {
            color: '#dcdf6e'
          }
        }
      },
      series: [
        {
          name: '频数',
          type: 'bar',
          data: this.histogram.serie,
          markPoint: {
            data: [
              { type: 'max', name: '最大值' },
              { type: 'min', name: '最小值' }
            ]
          },
          markLine: { data: [{ type: 'average', name: '平均值' }] }
        }
      ]
    };
    this.chart.setOption(this.option);
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.details-histogram {
  .header {
    padding: 0 1rem;
    margin-bottom: 0.5rem;
  }
}
</style>
