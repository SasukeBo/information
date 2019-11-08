<template>
  <div class="product-setting-name content-block">
    <div class="field-title">产品名称</div>
    <float-label-input :enter="submit" v-model="product.name" placeholder="请填写产品名称"></float-label-input>
    <el-button class="submit" type="primary" :loading="loading" @click="submit">保存修改</el-button>
  </div>
</template>
<script>
import FloatLabelInput from 'js/vue/main/components/float-label-input';
// graphql
import productNameQuery from './gql/query.product-name.gql';
import productUpdateMutate from './gql/mutate.product-update.gql';

export default {
  name: 'setting-name',
  components: { FloatLabelInput },
  props: ['id'],
  apollo: {
    product: {
      query: productNameQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      product: { name: '' },
      name: '',
      loading: false
    };
  },
  methods: {
    submit() {
      if (!this.product.name)
        this.$message({ type: 'error', message: '请填写产品名称！' });
      this.loading = true;
      this.$apollo
        .mutate({
          mutation: productUpdateMutate,
          variables: {
            id: this.id,
            name: this.product.name
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
