<template></template>
<script>
import devicesQuery from './gql/query.devices.gql';

export default {
  name: 'device-list',
  components: {},
  apollo: {
    deviceList: {
      query: devicesQuery,
      variables() {
        return {
          pattern: this.pattern,
          status: this.status,
          isRegister: this.isRegister,
          limit: this.limit,
          offset: this.offset
        };
      },
      fetchPolicy: 'network-only'
    }
  },
  data() {
    return {
      pattern: '',
      status: null,
      isRegister: false,
      limit: 9,
      currentPage: 1,
      isExpand: false,
      deviceList: {
        total: 0,
        devices: []
      }
    };
  },
  methods: {},
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
</style>
