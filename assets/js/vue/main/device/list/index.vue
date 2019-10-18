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
        <el-option v-for="op in options" :key="op.value" :label="op.label" :value="op.value"></el-option>
      </el-select>
    </div>

    <div class="list-table">
      <div class="table-row header-row">
        <span class="table-cell p-id">ID</span>
        <span class="table-cell">设备</span>
        <span class="table-cell">Token</span>
        <span class="table-cell">注册人</span>
        <span class="table-cell">物理地址</span>
        <span class="table-cell">良率</span>
        <span class="table-cell">稼动率</span>
        <span class="table-cell">状态</span>
      </div>

      <a
        class="table-row data-row"
        v-for="(device, index) in deviceList.devices"
        :key="'device_'+ index"
        @click.prevent="$router.push({name: 'device-show', params: {id: device.id}})"
        :href="'/device/'+device.id + '/show'"
      >
        <span class="table-cell p-id">{{device.id}}</span>
        <span class="table-cell">
          <span class="device-name">{{device.name}}</span>
          <span class="device-type">{{device.type}} - {{device.number}}</span>
        </span>

        <span class="table-cell device-token">{{device.token}}</span>

        <span class="table-cell">
          <div v-if="device.user && device.user.name" class="first">
            <img
              class="avatar"
              :src="device.user.avatarURL ? device.user.avatarURL : defaultAvatar"
            />
            {{ device.user.name }}
          </div>
          <span v-else>-</span>
        </span>

        <span class="table-cell">{{device.address}}</span>

        <span class="table-cell">{{device.statistics ? device.statistics.yield + '%' : '-'}}</span>

        <span class="table-cell">{{device.statistics ? device.statistics.activation + '%': '-'}}</span>

        <span class="table-cell">
          <i :class="['iconfont', statusMap[device.status].icon]"></i>
          <span>{{statusMap[device.status].label}}</span>
        </span>
      </a>
    </div>

    <div class="fix-footer-pagination">
      <el-pagination
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
        :current-page="currentPage"
        :page-sizes="[50, 100, 200, 500]"
        :page-size="limit"
        layout="total, sizes, prev, pager, next, jumper"
        :total="deviceList.total"
      ></el-pagination>
    </div>
  </div>
</template>
<script>
import deviceListQuery from './query.device-list.gql';
import defaultAvatar from 'images/default-avatar.png';

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
      limit: 50,
      search: this.$route.query.search || '',
      self: Boolean(this.$route.query.self) || false,
      status: undefined,
      pattern: this.$route.query.search || '',
      currentPage: 0,
      deviceList: {
        total: 0,
        devices: []
      },
      options: [
        { label: '全部', value: undefined },
        { label: '运行中', value: 'prod' },
        { label: '离线', value: 'offline' },
        { label: '停机', value: 'stop' }
      ],
      statusMap: {
        prod: { icon: 'icon-running', label: '运行中' },
        stop: { icon: 'icon-stopping', label: '停机' },
        offline: { icon: 'icon-offline', label: '离线' }
      },
      defaultAvatar
    };
  },
  computed: {
    offset() {
      return (this.currentPage - 1) * this.limit;
    }
  },
  watch: {
    search(newVal) {
      this.$router.push({
        name: this.$route.name,
        query: { ...this.$route.query, search: newVal }
      });
    },
    self(newVal) {
      this.$router.push({
        name: this.$route.name,
        query: { ...this.$route.query, self: newVal }
      });
    }
  },
  methods: {
    handleSizeChange(size) {
      this.limit = size;
    },
    handleCurrentChange(index) {
      this.currentPage = index;
    }
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

  .table-row.header-row .table-cell {
    color: $--color-font__white;
  }

  .table-cell {
    color: $--color-font__silver;
  }

  .table-cell .device-name {
    display: block;
    line-height: 1.5rem;
    color: $--color-font__white;
    font-weight: bolder;
  }

  .table-cell .device-type {
    font-size: 0.8rem;
  }

  .table-cell .avatar {
    width: 30px;
    height: 30px;
    vertical-align: middle;
    border-radius: 50%;
    margin-right: 0.5rem;
  }

  .table-cell.device-token {
    color: $--color-font__white;
    font-weight: bold;
  }

  .table-cell.p-id {
    text-align: center;
    color: $--color-font__light;
  }

  .table-cell .iconfont {
    margin-right: 10px;

    &.icon-running {
      font-size: 1.1rem;
      color: $--color-theme__success;
    }

    &.icon-offline {
      font-size: 1.2rem;
      color: $--color-theme__gray;
    }

    &.icon-stopping {
      color: $--color-theme__danger;
    }
  }

  .el-pagination {
    text-align: center;
  }

  .fix-footer-pagination {
    position: fixed;
    bottom: 0;
    width: 100%;
    background: $--color-theme__bg;
    box-shadow: -1px -2px 1px rgba(0, 0, 0, 0.1);
    left: 0;
    padding: 1rem 0;
  }

  .list-table {
    margin-bottom: 64px;
  }
}
</style>
