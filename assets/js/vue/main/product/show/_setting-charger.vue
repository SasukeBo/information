<template>
  <div class="product-setting-charger content-block">
    <div class="field-title">生产负责人</div>
    <float-label-input v-model="product.productor" :enter="submit" placeholder="填写姓名"></float-label-input>
    <float-label-input v-model="product.productorContact" :enter="submit" placeholder="填写手机号"></float-label-input>
    <el-button type="primary" :loading="loading" @click="submit">保存修改</el-button>
  </div>
</template>
<script>
import FloatLabelInput from 'js/vue/main/components/float-label-input';
// graphql
import productDetailsQuery from './gql/query.product-details.gql';
import productUpdateMutate from './gql/mutate.product-update.gql';

export default {
  name: 'setting-charger',
  props: ['id'],
  components: { FloatLabelInput },
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
      loading: false,
      product: {}
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
            productor: this.product.productor,
            productorContact: this.product.productorContact
          }
        })
        .then(({ data }) => {
          this.loading = false;
          this.$message({ type: 'success', message: '保存成功' });
        })
        .catch(e => {
          this.loading = false;
          this.$message({ type: 'error', message: e });
        });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
