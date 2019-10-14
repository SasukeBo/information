<template>
  <div class="product-list">
    <div class="header">产品列表</div>
    <div class="header_hr"></div>
    <div class="search-bar">
      <el-input placeholder="搜索你的产品" prefix-icon="el-icon-search" v-model="search"></el-input>
      <el-checkbox v-model="self">只看我自己</el-checkbox>
    </div>

    <div></div>
  </div>
</template>
<script>
import productListQuery from './query.product-list.gql';

export default {
  name: 'product-list',
  apollo: {
    productList: {
      query: productListQuery,
      variables() {
        return {
          namePattern: this.namePattern,
          self: this.self
        };
      }
    }
  },
  data() {
    return {
      search: '',
      namePattern: '',
      self: false,
      productList: {
        count: 0,
        products: []
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

.product-list {
  .header {
    font-size: 36px;
    padding: 16px 0 32px;
    margin-top: 27px;
  }

  .header_hr {
    position: relative;
    width: 100%;
    height: 1px;
    background: $--color-theme__light;
  }

  .search-bar {
    margin-top: 1rem;

    .el-input {
      width: 150px;
      margin-right: 1rem;
    }

    .el-checkbox {
      display: inline-block;
      vertical-align: bottom;
    }
  }
}
</style>
