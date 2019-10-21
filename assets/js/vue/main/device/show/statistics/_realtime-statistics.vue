<template>
  <div
    class="realtime"
    v-loading="$apollo.queries.statistics.loading"
    element-loading-background="unset"
  >
    <div v-for="item in statistics" :key="'data_' + item.id">
      <span>{{item.sign}}</span>
      <span>{{item.value}}</span>
      <span>{{item.createdAt}}</span>
    </div>
  </div>
</template>
<script>
import realtimeQuery from '../gql/query.realtime-statistics.gql';

export default {
  name: 'realtime',
  props: ['deviceID', 'productID'],
  apollo: {
    statistics: {
      query: realtimeQuery,
      variables() {
        return {
          deviceID: this.deviceID,
          productID: this.productID,
          limit: 100,
          afterTime: this.afterTime
        };
      }
    }
  },
  data() {
    return {
      afterTime: undefined,
      statistics: []
    };
  }
};
</script>
