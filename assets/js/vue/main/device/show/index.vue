<template>
  <div class="device-show">
    <div class="header">
      <div class="header_left">
        <i class="iconfont icon-shebei"></i>
        <div class="title">
          <div>
            <span>{{ device.name }}</span>
            <span style="font-size: 1.3rem">设备信息</span>
          </div>

          <div class="subtitle">
            <el-tooltip effect="dark" content="设备Token" placement="bottom">
              <span class="item">{{ device.token }}</span>
            </el-tooltip>

            <el-tooltip effect="dark" content="设备类型" placement="bottom">
              <span class="item">{{ device.type }}</span>
            </el-tooltip>

            <el-tooltip effect="dark" content="设备编号" placement="bottom">
              <span class="item">{{ device.number }}</span>
            </el-tooltip>
          </div>
        </div>
      </div>

      <div class="header_right">
        <el-tooltip effect="dark" content="修改信息" placement="top">
          <i class="operation-btn el-icon-edit" @click="() => {null}"></i>
        </el-tooltip>

        <el-tooltip effect="dark" content="删除设备" placement="top">
          <i class="operation-btn el-icon-delete" @click="confirmDelete"></i>
        </el-tooltip>
      </div>
    </div>

    <div class="tabs">
      <div class="tab-container">
        <ul class="nav-tabs">
          <li>
            <a :class="{active: currentTab === 'overview'}" @click="updateTab('overview')">总览</a>
          </li>
          <li>
            <a :class="{active: currentTab === 'statistics'}" @click="updateTab('statistics')">生产数据</a>
          </li>
          <li>
            <a :class="{active: currentTab === 'logs'}" @click="updateTab('logs')">停机日志</a>
          </li>
        </ul>
      </div>
    </div>

    <div class="tab-hr">
      <div class="create-device-btn">
        <el-button
          icon="el-icon-plus"
          circle
          type="primary"
          @click="$router.push({name: 'device-new'})"
        ></el-button>
      </div>
    </div>

    <component :is="currentTab" :id="id"></component>
  </div>
</template>
<script>
import deviceQuery from './gql/query.device.gql';
// components
import overview from './_overview';
import statistics from './statistics';
import logs from './status-log/';

export default {
  name: 'device-show',
  props: ['id'],
  components: {
    logs,
    overview,
    statistics
  },
  apollo: {
    device: {
      query: deviceQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      device: {}
    };
  },
  computed: {
    currentTab() {
      return this.$route.query.tab ? this.$route.query.tab : 'overview';
    }
  },
  methods: {
    confirmDelete() {
      var content =
        '确认删除该设备吗？删除该设备后，此设备生产的产品数据仍然会保留，可通过产品详情页查看保留的生产数据。';
      this.$confirm(content, '确认删除？', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(() => {})
        .catch(() => this.$message({ type: 'info', message: '已取消删除' }));
    },
    updateTab(tab) {
      if (this.currentTab !== tab)
        this.$router.push({ name: this.$route.name, query: { tab } });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-show {
  .header {
    margin-top: 27px;
    padding: 16px 0 32px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }

  .header_left {
    position: relative;
    flex-grow: 2;
    margin-right: 50px;
  }

  .header_left .title {
    padding-left: 5rem;
    font-size: 36px;
  }

  .header_left .subtitle {
    font-size: 13px;
    color: $--color-font__gray;

    .item {
      margin-right: 24px;
    }
  }

  .header_left .iconfont {
    position: absolute;
    font-size: 4rem;
    padding-right: 1rem;
    color: $--color-theme__main;
  }

  .header_right {
    display: flex;
    margin-top: 18px;
  }

  .header_right .operation-btn {
    color: $--color-font__silver;
    display: block;
    cursor: pointer;
    width: 50px;
    height: 50px;
    font-size: 1.5rem;
    line-height: 50px;
    text-align: center;
    padding: 1px 7px 2px;
  }

  .header_right .operation-btn.el-icon-edit:hover {
    color: $--color-theme__main;
  }

  .header_right .operation-btn.el-icon-delete:hover {
    color: $--color-theme__danger;
  }

  .tabs {
    margin: 0px auto 48px auto;
  }

  .tabs .nav-tabs {
    overflow-y: hidden;
    overflow-x: auto;
    transition: all 0.3s ease-in;
    margin: 0;
    white-space: nowrap;
    padding-left: 0;
    list-style: none;
    cursor: default;
  }

  .tabs .nav-tabs li {
    float: none;
    display: inline-block;
    width: auto;

    &:first-child a {
      margin-left: 0;
    }
  }

  .tabs .nav-tabs a {
    font-size: 1em;
    margin: 0 14px;
    padding: 14px 2px;
    color: $--color-font__silver;
    cursor: pointer;
    display: block;
    position: relative;

    &.active,
    &:hover {
      font-weight: bold;
      color: $--color-font__white;
    }

    &:after {
      position: absolute;
      bottom: 0;
      left: 50%;
      content: '';
      display: block;
      height: 2px;
      width: 0;
      z-index: 1;
      transition: ease-in all 0.15s;
    }

    &.active:after {
      margin-left: -50%;
      width: 100%;
      background-color: $--color-theme__main;
    }
  }

  .tab-hr {
    height: 1px;
    position: relative;
    top: -49px;
    background: $--color-theme__light;
  }

  .tab-hr .create-device-btn {
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
}
</style>
