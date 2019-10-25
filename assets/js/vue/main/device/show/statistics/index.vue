<template>
  <div class="device-statistics">
    <div class="block">
      <div class="field-title">最近一个月数据统计</div>
      <month-data :id="id"></month-data>
    </div>

    <div class="block" v-if="product">
      <div class="field-title">实时数据</div>
      <real-time :deviceID="id" :product="product"></real-time>
    </div>

    <div class="block" v-if="product">
      <div class="field-title">检测项直方图</div>
      <div class="select-item">
        <el-select v-model="selected" placeholder="选择检测项" size="small">
          <el-option
            v-for="(item, index) in product.detectItems"
            :key="item.sign"
            :label="item.sign"
            :value="index"
          ></el-option>
        </el-select>
      </div>
      <div class="description">统计产品检测项数据，生成直方图，区间段由检测项值最大最小值区间等分40后获得</div>
      <histogram :deviceID="id" :productID="product.id" :detectItem="product.detectItems[selected]"></histogram>
    </div>

    <div class="block">
      <div class="field-title">设备状态统计</div>
      <device-status-chart :deviceID="id"></device-status-chart>
    </div>
  </div>
</template>
<script>
import MonthData from './_month-data';
import RealTime from './_realtime-statistics';
import Histogram from './_histogram';
import DeviceStatusChart from './_device-status';

import productQuery from '../gql/query.device-products.gql';

export default {
  name: 'device-statistics',
  components: {
    Histogram,
    MonthData,
    RealTime,
    DeviceStatusChart
  },
  props: ['id'],
  apollo: {
    product: {
      query: productQuery,
      variables() {
        return { deviceID: this.id };
      },
      update(data) {
        return data.device.product;
      }
    }
  },
  data() {
    return {
      product: undefined,
      selected: 0
    };
  },
  watch: {
    product(newVal) {
      if (newVal && newVal.detectItems && newVal.detectItems.length)
        this.selectedItem = newVal.detectItems[0];
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';
.device-statistics {
  .block {
    position: relative;
    margin-bottom: 2rem;
  }

  .select-item {
    position: absolute;
    right: 0;
    top: 0;
  }

  .description {
    color: $--color-font__gray-deep;
    border-left: 5px solid $--color-theme__gray;
    padding-left: 1rem;
  }

  .block .field-title {
    font-size: 24px;
    line-height: 24px;
    margin-bottom: 25px;
    font-weight: 400;
  }
}
</style>
