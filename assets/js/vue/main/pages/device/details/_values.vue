<template>
  <div class="param-values global-card">
    <div class="header">
      <el-date-picker
        v-model="value2"
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
          <span>表格</span>
        </el-carousel-item>
      </el-carousel>
    </div>
  </div>
</template>
<script>
import { DatePicker, Carousel, CarouselItem } from 'element-ui';

export default {
  name: 'param-values',
  props: ['uuid'],
  components: {
    ElDatePicker: DatePicker,
    ElCarousel: Carousel,
    ElCarouselItem: CarouselItem
  },
  data() {
    return {
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
      },
      value2: [],
      currentIndex: 0
    };
  },
  methods: {
    changeIndex(index) {
      this.currentIndex = index;
    }
  }
};
</script>
