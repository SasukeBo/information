<template>
  <div
    class="param-values global-card"
    v-loading="$apollo.queries.params.loading"
    element-loading-background="unset"
  >
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
        <span
          class="indicator"
          @mouseenter="$refs.values.activeIndex = 2"
          :class="[currentIndex === 2 ? 'is-active' : '']"
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
        <el-carousel-item key="histogram">
          <span>直方图</span>
        </el-carousel-item>
        <el-carousel-item key="line">
          <span>折线图</span>
        </el-carousel-item>
        <el-carousel-item key="table">
          <values-table v-if="params && params.length" :params="params"></values-table>
        </el-carousel-item>
      </el-carousel>
    </div>
  </div>
</template>
<script>
import { DatePicker, Carousel, CarouselItem } from 'element-ui';
import paramWithValues from './gql/query.param-values.gql';
import ValuesTable from './_values-table';

export default {
  /**
   * 走马灯item数据请求分开
   * 结合实际性能
   */
  name: 'param-values',
  props: ['uuid'],
  components: {
    ValuesTable,
    ElDatePicker: DatePicker,
    ElCarousel: Carousel,
    ElCarouselItem: CarouselItem
  },
  apollo: {
    params: {
      query: paramWithValues,
      variables() {
        return {
          deviceUUID: this.uuid,
          beforeTime: this.beforeTime,
          afterTime: this.afterTime
        };
      }
    }
  },
  data() {
    return {
      params: [],
      currentIndex: 2,
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
