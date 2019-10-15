<template>
  <div class="device-list">
    <div class="header">设备列表</div>
    <div class="header-hr">
      <div class="create-device-btn">
        <el-button
          icon="el-icon-plus"
          circle
          type="primary"
          @click="$router.push({name: 'device-new'})"
        ></el-button>
      </div>
    </div>

    <div class="search-bar">
      <el-input
        placeholder="设备类型/设备名称/地址/设备编号模糊搜索"
        prefix-icon="el-icon-search"
        v-model="pattern"
        class="filter__search"
        @keyup.native.enter="search = pattern"
      ></el-input>

      <el-checkbox v-model="self" class="custom-checkbox-vertical">只看我自己</el-checkbox>

      <el-select v-model="status" placeholder="状态筛选" class="filter__status">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>
  </div>
</template>
<script>
import deviceListQuery from './query.device-list.gql';

export default {
  name: 'device-list',
  apollo: {
    deviceList: {
      query: deviceListQuery,
      variables() {
        return {
          limit: this.limit,
          offset: this.offset,
          search: this.search,
          self: this.self,
          status: this.status
        };
      }
    }
  },
  data() {
    return {
      limit: 0,
      offset: 0,
      search: '',
      self: false,
      status: undefined,
      pattern: '',
      deviceList: {
        total: 0,
        devices: []
      }
    };
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-list {
  .header {
    font-size: 36px;
    padding: 16px 0 32px;
    margin-top: 27px;
  }

  .header-hr {
    position: relative;
    width: 100%;
    height: 1px;
    background: $--color-theme__light;
  }

  .header-hr .create-device-btn {
    position: absolute;
    right: 11px;
    top: -25px;

    .el-button {
      box-shadow: 0 0 14px 0px #000;
      transition: background ease 0.3s;
    }

    .el-button i {
      font-size: 1.5rem;
    }
  }

  .search-bar {
    margin: 3rem 0 2rem 0;

    .filter__search {
      width: 320px;
      margin-right: 1rem;
    }

    .custom-checkbox-vertical {
      display: inline-block;
      vertical-align: bottom;
      color: $--color-font__light;

      .el-checkbox__input,
      .el-checkbox__label {
        display: table-cell;
      }
    }

    .filter__status {
      float: right;
    }
  }
}
</style>
