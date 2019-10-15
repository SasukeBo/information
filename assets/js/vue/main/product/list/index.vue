<template>
  <div class="product-list">
    <div class="header">产品列表</div>
    <div class="header-hr">
      <div class="create-product-btn">
        <el-button
          icon="el-icon-plus"
          circle
          type="primary"
          @click="$router.push({name: 'product-new'})"
        ></el-button>
      </div>
    </div>

    <div class="search-bar">
      <el-input
        placeholder="搜索你的产品"
        prefix-icon="el-icon-search"
        v-model="search"
        @keyup.native.enter="searchList"
      ></el-input>
      <el-checkbox v-model="self" class="custom-checkbox-vertical">只看我自己</el-checkbox>
    </div>

    <div class="list-table">
      <div class="table-row header-row">
        <span class="table-cell p-id">ID</span>
        <span class="table-cell">产品</span>
        <span class="table-cell">注册人</span>
        <span class="table-cell">订货方</span>
        <span class="table-cell">生产负责人</span>
        <span class="table-cell">目标/当前产量</span>
        <span class="table-cell">检测项数</span>
      </div>

      <div
        class="table-row data-row"
        v-for="(product, index) in productList.products"
        :key="'product_' + index"
      >
        <span class="table-cell p-id">{{product.id}}</span>
        <span class="table-cell">
          <a
            @click.prevent="$router.push({name: 'product-show', params: {id: product.id}})"
            :href="'/product/'+product.id + '/show'"
            class="name"
          >{{ product.name }}</a>
          <span class="order-num">订单号: {{ product.orderNum ? product.orderNum : '-' }}</span>
        </span>

        <a
          :href="'/product/'+product.id + '/show'"
          @click.prevent="$router.push({name: 'product-show', params: {id: product.id}})"
          class="table-cell"
        >
          <div v-if="product.register && product.register.name" class="first">
            <img
              class="avatar"
              :src="product.register.avatarURL ? product.register.avatarURL : defaultAvatar"
            />
            {{ product.register.name }}
          </div>
          <span v-else>-</span>
        </a>

        <a
          @click.prevent="$router.push({name: 'product-show', params: {id: product.id}})"
          :href="'/product/'+product.id + '/show'"
          class="table-cell"
        >
          <div class="first" v-if="product.customer">{{ product.customer }}</div>
          <span v-else>-</span>
        </a>

        <a
          @click.prevent="$router.push({name: 'product-show', params: {id: product.id}})"
          :href="'/product/'+product.id + '/show'"
          class="table-cell"
        >
          <div class="first" v-if="product.productor">{{ product.productor }}</div>
          <span v-else>-</span>
        </a>

        <a
          @click.prevent="$router.push({name: 'product-show', params: {id: product.id}})"
          :href="'/product/'+product.id + '/show'"
          class="table-cell"
        >
          <div class="first">
            <span style="color: #03a9f4">{{ product.total ? product.total : '0' }}</span>
            <span>/</span>
            <span style="color: #8fc860">{{ product.currentCount ? product.currentCount : '0' }}</span>
          </div>
        </a>

        <a
          @click.prevent="$router.push({name: 'product-show', params: {id: product.id}, query: {tab: 'setting'}, hash: '#detect-items'})"
          :href="'/product/'+product.id + '/show?tab=setting#detect-items'"
          class="table-cell"
        >
          <div
            class="first"
            style="color: #03a9f4"
          >{{product.detectItemsCount ? product.detectItemsCount : '0'}}</div>
        </a>
      </div>
    </div>
  </div>
</template>
<script>
import productListQuery from './query.product-list.gql';
import defaultAvatar from 'images/default-avatar.png';

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
      defaultAvatar,
      search: '',
      namePattern: '',
      self: false,
      productList: {
        count: 0,
        products: []
      }
    };
  },
  methods: {
    searchList() {
      this.namePattern = this.search;
    }
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

  .header-hr {
    position: relative;
    width: 100%;
    height: 1px;
    background: $--color-theme__light;
  }

  .header-hr .create-product-btn {
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
    margin: 1rem 0 2rem 0;

    .el-input {
      width: 200px;
      margin-right: 1rem;
    }

    .el-checkbox {
      display: inline-block;
      vertical-align: bottom;
    }
  }

  .list-table {
    display: table;
    width: 100%;
    border: 1px solid $--color-theme__light;
    border-bottom: none;
  }

  .table-row {
    display: table-row;
  }

  .table-cell {
    display: table-cell;
    width: auto;
    padding: 15px 10px;
    border-bottom: 1px solid $--color-theme__light;
  }

  .data-row .table-cell {
    height: 70px;
    vertical-align: middle;
    padding: 10px;
    cursor: pointer;
  }

  .table-cell.p-id {
    text-align: center;
    color: $--color-font__light;
  }

  .data-row {
    color: $--color-font__gray;
  }

  .table-cell a {
    display: block;
  }

  .table-cell .name {
    font-weight: bold;
    color: $--color-font__white;
    line-height: 1.5rem;
  }

  .table-cell .order-num {
    font-size: 13px;
  }

  a.table-cell {
    color: $--color-font__gray;

    &:hover {
      color: $--color-font__gray;
    }

    .first {
      line-height: 1.5rem;
      color: $--color-font__light;

      .avatar {
        width: 30px;
        height: 30px;
        vertical-align: middle;
        border-radius: 50%;
        margin-right: 0.5rem;
      }
    }
  }

  .header-row {
    background: linear-gradient(141deg, #aca0f2 0%, #0286c2 50%, #aca0f2 100%);
  }

  .custom-checkbox-vertical {
    color: $--color-font__light;

    .el-checkbox__input,
    .el-checkbox__label {
      display: table-cell;
    }
  }
}
</style>
