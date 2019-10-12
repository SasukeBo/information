<template>
  <div
    class="product-show"
    v-loading="$apollo.queries.product.loading"
    element-loading-background="unset"
  >
    <div class="header">
      <i class="iconfont icon-chanpin"></i>
      <div class="header-content">
        <div>{{ product.name }}</div>
        <div class="header-content__name">产品信息</div>
        <a
          class="to-product-edit"
          :href="'/product/' + id + '/edit'"
          @click.prevent="$router.push({name: 'product-edit', params: {id}})"
        >编辑</a>
      </div>
    </div>

    <div class="tabs">
      <div class="tab-container">
        <ul class="nav-tabs">
          <li>
            <a :class="{active: currentTab === 'overview'}" @click="updateTab('overview')">总览</a>
          </li>
          <li>
            <a :class="{active: currentTab === 'devices'}" @click="updateTab('devices')">生产设备</a>
          </li>
          <li>
            <a :class="{active: currentTab === 'instances'}" @click="updateTab('instances')">产品实例</a>
          </li>
        </ul>
      </div>
    </div>

    <div class="tab-hr">
      <div class="create-product-btn">
        <el-button
          icon="el-icon-plus"
          circle
          type="primary"
          @click="$router.push({name: 'product-new'})"
        ></el-button>
      </div>
    </div>

    <component :is="currentTab" :id="id"></component>
  </div>
</template>
<script>
// gql
import productGetQuery from './gql/query.product-name.gql';
// components
import overview from './_overview.vue';
import instances from './_instance.vue';
import devices from './_devices.vue';

export default {
  name: 'product-show',
  props: ['id'],
  components: {
    overview,
    instances,
    devices
  },
  apollo: {
    product: {
      query: productGetQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      product: {}
    };
  },
  computed: {
    currentTab() {
      return this.$route.query.tab ? this.$route.query.tab : 'overview';
    }
  },
  methods: {
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

.product-show {
  .header {
    font-size: 36px;
    padding: 16px 0 32px;
    margin-top: 27px;
    position: relative;

    .iconfont {
      position: absolute;
      font-size: 4rem;
      padding-right: 1rem;
      color: $--color-theme__main;
    }

    .header-content {
      padding-left: 5rem;
      position: relative;

      .to-product-edit {
        position: absolute;
        font-size: 14px;
        bottom: 0;
        display: inline-block;
        left: 150px;
        line-height: 18px;
      }
    }

    .header-content__name {
      font-size: 13px;
    }
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

  .tab-hr .create-product-btn {
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
