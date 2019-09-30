<template>
  <div class="device-list">
    <div class="header-fix">
      <div class="title">设备列表</div>
      <div class="btns">
        <el-button size="small" type="info" icon="el-icon-refresh"></el-button>
        <el-button size="small" type="primary" @click="$router.push({name: 'device-new'})">创建设备</el-button>
      </div>
    </div>

    <div class="device-filter">
      <el-input
        placeholder="设备类型/设备名称/地址/设备编号模糊搜索"
        v-model="pattern"
        size="small"
        class="filter__search"
      >
        <el-button slot="append" type="primary" icon="el-icon-search" @click="search = pattern"></el-button>
      </el-input>

      <div class="filter__register">
        <el-checkbox v-model="filter" label="只看我注册的" size="small"></el-checkbox>
      </div>

      <el-select v-model="status" placeholder="状态筛选" size="small" class="filter__status">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>

    <div class="device-list__body"></div>
  </div>
</template>
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
          search: this.search,
          status: this.status,
          filter: this.filter,
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
      search: '',
      status: null,
      filter: false,
      limit: 9,
      currentPage: 1,
      isExpand: false,
      options: [
        {
          value: 'prod',
          label: '生产中'
        },
        {
          value: 'stop',
          label: '停机中'
        },
        {
          value: 'offline',
          label: '离线'
        },
        {
          value: null,
          label: '全部'
        }
      ],
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
@import 'css/main/devices/index.scss';
</style>
