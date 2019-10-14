<template>
  <div class="product-form">
    <div class="header">{{ action === 'new' ? '创建产品' : '修改产品'}}</div>
    <div class="header_hr"></div>
    <div class="error-message" :style="{height: message ? '52px' : '0'}">
      <div class="message">{{ message }}</div>
      <el-button type="text" icon="el-icon-close" @click="message=''"></el-button>
    </div>
    <div class="form">
      <form class="flex-layout-form">
        <float-label-input v-model="form.name" placeholder="产品名称"></float-label-input>
        <float-label-input v-model="form.orderNum" placeholder="订单号"></float-label-input>
        <float-label-input v-model="form.customer" placeholder="需求方"></float-label-input>
        <float-label-input v-model="form.customerContact" placeholder="需求方联系电话"></float-label-input>
        <float-label-input v-model="form.total" placeholder="计划生产总数"></float-label-input>
        <float-label-input v-model="form.productor" placeholder="生产负责人"></float-label-input>
        <float-label-input v-model="form.productorContact" placeholder="生产负责人联系电话"></float-label-input>
        <float-label-input type="date" v-model="form.finishTime" placeholder="计划完成时间"></float-label-input>

        <div class="detect-items">
          <span class="title">检测项</span>
          <el-button type="primary" size="small" @click="addNewItem">增加</el-button>
        </div>

        <div class="detect-item" v-for="(item, i) in detectItems" :key="'item_' + i">
          <float-label-input v-model="item.sign" placeholder="检测项名称"></float-label-input>
          <float-label-input v-model="item.upperLimit" placeholder="值上限"></float-label-input>
          <float-label-input v-model="item.lowerLimit" placeholder="值下限"></float-label-input>
          <el-button type="danger" icon="el-icon-delete" @click="detectItems.splice(i, 1)">删除</el-button>
        </div>

        <el-button
          :loading="loading"
          class="submit-btn"
          type="primary"
          @click="submit"
        >{{ action === 'new' ? '注册产品' : '保存修改'}}</el-button>
      </form>
    </div>
  </div>
</template>
<script>
import FloatLabelInput from 'js/vue/main/components/float-label-input';
import productCreateMutate from './mutate.create-product.gql';

export default {
  name: 'product-form',
  props: ['id'],
  components: {
    FloatLabelInput
  },
  data() {
    return {
      form: {
        name: ''
      },
      name: '',
      message: '',
      detectItems: [],
      loading: false
    };
  },
  computed: {
    action() {
      if (this.$route.name === 'product-edit') return 'edit';
      else return 'new';
    }
  },
  methods: {
    addNewItem() {
      this.detectItems.push({
        sign: '',
        upperLimit: undefined,
        lowerLimit: undefined
      });
    },
    submit() {
      this.loading = true;
      if (!this.form.name) {
        this.message = '请填写产品名称！';
        this.loading = false;
        return;
      }

      for (var i = 0; i < this.detectItems.length; i++) {
        if (!this.detectItems[i].sign) {
          this.message = '信息有误，请正确填写检测项名称！';
          this.loading = false;
          return;
        }
      }

      this.$apollo
        .mutate({
          mutation: productCreateMutate,
          variables: {
            ...this.form,
            detectItems: this.detectItems
          }
        })
        .then(({ data }) =>
          this.$router.push({
            name: 'product-show',
            params: { id: data.productCreate.id }
          })
        )
        .catch(e => {
          console.log(e);
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

.product-form {
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

  .error-message {
    background: $--color-theme__danger;
    border-radius: 4px;
    margin-top: 5px;
    transition: height 0.3s ease-in-out;
    overflow: hidden;

    .message {
      padding: 1rem;
      display: inline-block;
    }

    .el-button {
      float: right;
      margin: 6px 1rem;
      font-size: 1rem;
      color: $--color-theme__white;
    }
  }

  .form {
    margin-top: 48px;
    width: 520px;
  }

  .flex-layout-form {
    justify-content: space-between;
    flex-wrap: wrap;
    display: flex;
  }

  .detect-items {
    width: 100%;
    border-bottom: 1px solid $--color-theme__light;
    margin-bottom: 15px;
    padding: 0 0 0.5rem;
    display: flex;

    .title {
      line-height: 32px;
      flex: auto;
    }
  }
  .detect-item {
    display: flex;
    justify-content: space-between;
    position: relative;

    .floating-label {
      width: 30%;
    }

    .el-button {
      position: absolute;
      right: -110px;
      top: 5px;
    }
  }

  .submit-btn {
    width: 100%;
  }
}
</style>
