<template>
  <div class="product-setting-detect content-block">
    <div class="field-title">
      产品检测项
      <el-button type="primary" size="small" @click="addNewItem" icon="el-icon-plus">增加</el-button>
    </div>

    <div class="detect-item" v-for="(item, i) in detectItems" :key="'item_' + i">
      <float-label-input v-model="item.sign" placeholder="检测项名称"></float-label-input>
      <float-label-input v-model="item.upperLimit" placeholder="值上限"></float-label-input>
      <float-label-input v-model="item.lowerLimit" placeholder="值下限"></float-label-input>
      <div class="btn-groups">
        <el-button class="save" icon="iconfont icon-save" type="text" @click="saveItem(item, i)"></el-button>
        <el-button class="delete" icon="el-icon-delete" type="text" @click="deleteItem(item, i)"></el-button>
      </div>
    </div>
  </div>
</template>
<script>
// graphql
import detectItemListQuery from './gql/query.detect-item-list.gql';
import detectItemUpdateMutate from './gql/mutate.detect-item-update.gql';
import detectItemCreateMutate from './gql/mutate.detect-item-create.gql';
import detectItemDeleteMutate from './gql/mutate.detect-item-delete.gql';

// components
import FloatLabelInput from 'js/vue/main/components/float-label-input';

export default {
  name: 'setting-detect',
  props: ['id'],
  components: { FloatLabelInput },
  apollo: {
    detectItems: {
      query: detectItemListQuery,
      variables() {
        return { id: this.id };
      },
      update: data => data.detectItemList.detectItems
    }
  },
  data() {
    return {
      detectItems: []
    };
  },
  methods: {
    saveItem(item, index) {
      if (!item.sign)
        this.$message({ type: 'error', message: '检测项名称为必填字段！' });

      if (item.id) {
        this.$apollo
          .mutate({
            mutation: detectItemUpdateMutate,
            variables: {
              ...item
            }
          })
          .then(({ data }) =>
            this.$message({ type: 'success', message: '检测项保存成功!' })
          )
          .catch(e => this.$message({ type: 'error', message: e }));
      } else {
        this.$apollo
          .mutate({
            mutation: detectItemCreateMutate,
            variables: {
              productID: this.id,
              ...item
            }
          })
          .then(({ data: { item } }) => {
            this.detectItems[index].id = item.id;
            this.$message({ type: 'success', message: '检测项保存成功!' });
          })
          .catch(e => this.$message({ type: 'error', message: e }));
      }
    },
    deleteItem(item, index) {
      if (item.id) {
        this.$apollo
          .mutate({
            mutation: detectItemDeleteMutate,
            variables: { id: item.id }
          })
          .then(() => {
            this.detectItems.splice(index, 1);
            this.$message({ type: 'success', message: '检测项删除成功!' });
          })
          .catch(e => this.$message({ type: 'error', message: e }));
      } else {
        this.detectItems.splice(index, 1);
      }
    },
    addNewItem() {
      this.detectItems.unshift({});
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.product-setting-detect .detect-item {
  display: flex;
  justify-content: space-between;
  position: relative;
  margin-bottom: 1rem;

  .floating-label {
    max-width: 25%;
    margin: 0;

    .global-input {
      height: 40px;
    }
  }

  .btn-groups {
    min-width: 80px;
    text-align: center;
    border-left: 2px solid;
  }

  .el-button {
    height: 40px;

    i {
      font-size: 1.1rem;
    }

    &.delete {
      color: darken($--color-theme__danger, 10%);

      &:hover {
        color: $--color-theme__danger;
      }
    }
  }
}
</style>
