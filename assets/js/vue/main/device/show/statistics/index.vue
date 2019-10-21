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
  </div>
</template>
<script>
import MonthData from './_month-data';
import RealTime from './_realtime-statistics';

import productQuery from '../gql/query.device-products.gql';

export default {
  name: 'device-statistics',
  components: { MonthData, RealTime },
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
      product: undefined
    };
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';
.device-statistics {
  .block {
    margin-bottom: 2rem;
  }

  .block .field-title {
    font-size: 24px;
    line-height: 24px;
    margin-bottom: 25px;
    font-weight: 400;
  }
}
</style>
