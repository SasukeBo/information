<template>
  <div class="param-values global-card">
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

      <div class="indicators">
        <span
          class="indicator"
          @mouseenter="$refs.values.activeIndex = 0"
          :class="[currentIndex === 0 ? 'is-active' : '']"
        ></span>
        <span
          class="indicator"
          @mouseenter="$refs.values.activeIndex = 1"
          :class="[currentIndex === 1 ? 'is-active' : '']"
        ></span>
      </div>
    </div>

    <div class="body">
      <el-carousel
        ref="values"
        height="500px"
        :initial-index="currentIndex"
        :autoplay="false"
        indicator-position="none"
        arrow="never"
        @change="changeIndex"
      >
        <el-carousel-item key="line">
          <span>折线图</span>
        </el-carousel-item>
        <el-carousel-item key="table">
          <values-table v-if="uuid" :uuid="uuid" :afterTime="afterTime" :beforeTime="beforeTime"></values-table>
        </el-carousel-item>
      </el-carousel>
    </div>
  </div>
</template>
<script>
import ValuesTable from './_values-table';

export default {
  /**
   * 走马灯item数据请求分开
   * 结合实际性能
   */
  name: 'param-values',
  props: ['uuid'],
  components: { ValuesTable },
  data() {
    return {
      currentIndex: 1,
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
  methods: {
    changeIndex(index) {
      this.currentIndex = index;
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-details .param-values {
  .header {
    padding: 0 1rem;
    margin-bottom: 1rem;

    .el-range-separator {
      flex: auto;
    }

    .indicators {
      float: right;
      padding: 6px;
      cursor: default;

      .indicator {
        width: 10px;
        height: 10px;
        display: inline-block;
        cursor: pointer;
        border-radius: 50%;
        margin: 0 0.2rem;
        background: $--color-theme__light;
        transition: all 0.3s ease;
      }

      .indicator.is-active,
      .indicator:hover {
        background: $--color-theme__white;
      }
    }
  }

  .body {
    padding: 0 1rem;
  }
}
</style>
