<template></template>
<script>
import { timeFormatter } from 'js/utils';
import valuesQuery from './gql/query.values.gql';
import valuesSub from './gql/sub.values.gql';

export default {
  name: 'param-value-sub',
  props: ['param', 'seriesData'],
  apollo: {
    values: {
      query: valuesQuery,
      variables() {
        var time = new Date();
        time.setSeconds(time.getSeconds() - 600);
        return {
          paramID: this.param.id,
          limit: 1000,
          after: time.toISOString()
        };
      },
      subscribeToMore: {
        document: valuesSub,
        variables() {
          return {
            t: `dpv:${this.param.id}`
          };
        },
        updateQuery: (preData, { subscriptionData }) => {
          if (!preData) {
            return { values: [subscriptionData.data.values] };
          }
          preData.values = [subscriptionData.data.values];
          return preData;
        }
      }
    }
  },
  data() {
    return {
      values: []
    }
  },
  watch: {
    values(newVal) {
      var newSeriesDate = this.formatValues(newVal);
      var seriesData = []
      if (this.seriesData) seriesData = this.seriesData;
      for (var i = newSeriesDate.length; i > 0; i--) {
        if (seriesData.length > 1000) seriesData.shift();
        seriesData.push(newSeriesDate[i - 1]);
      }
      // this.$emit('update:seriesData', seriesData);
    }
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
    }
  }
};
</script>
