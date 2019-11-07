<template>
  <div
    class="product-devices"
    v-loading="$apollo.queries.devices.loading"
    element-loading-background="rgba(0, 0, 0, 0.3)"
  >
    <el-table
      :data="devices"
      class="device-table"
      @row-click="goDetail"
      row-class-name="device-table-row"
      cell-class-name="device-table-cell"
      header-row-class-name="device-table-header-row"
      header-cell-class-name="device-table-header-cell"
    >
      <el-table-column label="id" prop="id" width="80px"></el-table-column>
      <el-table-column label="设备">
        <template slot-scope="scope">
          <div>
            <div class="name">{{scope.row.name}}</div>
            <div class="type-number">
              <span>{{scope.row.type}}</span>
              <span>-</span>
              <span>{{scope.row.number}}</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Token" prop="token"></el-table-column>
      <el-table-column label="注册人">
        <template slot-scope="scope">
          <div>
            <img
              class="avatar"
              :src="scope.row.user && scope.row.user.avatarURL ? scope.row.user.avatarURL : defaultAvatar"
            />
            {{ scope.row.user.name }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="物理地址" prop="address"></el-table-column>
      <el-table-column label="生产指标" width="200px">
        <template slot-scope="scope">
          <div class="device-statistics">
            <div class="item">
              <div>稼动率</div>
              <div>{{ scope.row.statistics && scope.row.statistics.availability ? (scope.row.statistics.availability * 100).toFixed(2) + '%' : '-'}}</div>
            </div>

            <div class="item">
              <div>良率</div>
              <div>{{ scope.row.statistics && scope.row.statistics.quality ? (scope.row.statistics.quality * 100).toFixed(2) + '%' : '-'}}</div>
            </div>

            <div class="item">
              <div>OEE</div>
              <div>{{ scope.row.statistics && scope.row.statistics.oee ? (scope.row.statistics.oee * 100).toFixed(2) + '%' : '-'}}</div>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="状态">
        <template slot-scope="scope">
          <div>
            <i :class="['iconfont', statusMap[scope.row.status].icon]"></i>
            <span>{{statusMap[scope.row.status].label}}</span>
          </div>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>
<script>
import deviceQuery from './gql/query.devices.gql';
import defaultAvatar from 'images/default-avatar.png';

export default {
  name: 'product-devices',
  props: ['id'],
  apollo: {
    devices: {
      query: deviceQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      defaultAvatar,
      devices: [],
      statusMap: {
        prod: { icon: 'icon-running', label: '运行中' },
        stop: { icon: 'icon-stopping', label: '停机' },
        offline: { icon: 'icon-offline', label: '离线' }
      }
    };
  },
  methods: {
    goDetail(row) {
      this.$router.push({ name: 'device-show', params: { id: row.id } });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.product-devices {
  .device-table {
    margin-bottom: 20px;
    background: $--color-theme__bg;

    .el-table__empty-block {
      height: unset;
    }
  }

  .device-table .device-table-cell .avatar {
    display: inline-block;
    width: 40px;
    height: 40px;
    margin-right: 8px;
    vertical-align: middle;
    border-radius: 50%;
  }

  .device-table .device-table-cell .name {
    color: $--color-theme__white;
    font-weight: bold;
  }

  .device-table .device-table-cell .type-number {
    font-size: 0.75rem;
  }

  .device-table .device-table-header-cell {
    background: $--color-theme__dark;
    color: $--color-theme__white;
    border-top: 1px solid $--color-theme__dark;
    border-bottom: 1px solid $--color-theme__dark;
  }

  .device-table .device-table-row:hover .device-table-cell {
    color: $--color-theme__gray;

    .name {
      color: $--color-theme__dark;
    }
  }

  .device-table .device-table-cell {
    cursor: pointer;
    background: $--color-theme__bg;
    font-size: 0.85rem;
    color: $--color-font__silver;
    border-bottom: 1px solid $--color-font__gray;
  }

  .device-table .iconfont {
    margin-right: 8px;
    &.icon-running {
      font-size: 1.1rem;
      color: $--color-theme__success;
    }

    &.icon-stopping {
      color: $--color-theme__danger;
    }

    &.icon-offline {
      font-size: 1.2rem;
      color: $--color-theme__gray;
    }
  }

  .device-table .device-statistics {
    display: flex;

    .item {
      flex: auto;
      padding: 0 5px;
      border-right: 1px solid;
    }

    .item:last-child {
      border-right: none;
    }
  }
}
</style>
