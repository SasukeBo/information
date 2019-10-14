<template>
  <div class="product-setting-order content-block">
    <div class="field-title">订单信息</div>
    <float-label-input :enter="submit" v-model="product.orderNum" placeholder="订单编号"></float-label-input>
    <float-label-input :enter="submit" v-model="product.total" placeholder="订货量"></float-label-input>
    <float-label-input
      type="date"
      :enter="submit"
      v-model="product.finishTime"
      placeholder="预计完成日期"
    ></float-label-input>
    <el-button type="primary" :loading="loading" @click="submit">保存修改</el-button>
  </div>
</template>
<script>
import FloatLabelInput from 'js/vue/main/components/float-label-input';
// graphql
import productDetailsQuery from './gql/query.product-details.gql';
import productUpdateMutate from './gql/mutate.product-update.gql';

export default {
  name: 'setting-order',
  components: { FloatLabelInput },
  props: ['id'],
  apollo: {
    product: {
      query: productDetailsQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      product: {},
      loading: false
    };
  },
  methods: {
    submit() {
      this.loading = true;
      this.$apollo
        .mutate({
          mutation: productUpdateMutate,
          variables: {
            id: this.id,
            orderNum: this.product.orderNum,
            total: this.product.total,
            finishTime: this.product.finishTime
          }
        })
        .then(({ data }) => {
          this.$message({ type: 'success', message: '保存成功。' });
          this.loading = false;
        })
        .catch(e => {
          this.$message({ type: 'error', message: e });
          this.loading = false;
        });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.product-setting-order .el-date-editor {
  width: 100%;
  margin-bottom: 1rem;

  .el-input__inner {
    padding: 0 15px;
    height: 50px;
    color: $--color-font__dark;
  }
}
</style>
