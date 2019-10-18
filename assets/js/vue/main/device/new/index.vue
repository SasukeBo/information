<template>
  <div class="device-new">
    <div class="header">创建设备</div>
    <div class="header_hr"></div>
    <div class="error-message" :style="{height: message ? '52px' : '0'}">
      <div class="message">{{ message }}</div>
      <el-button type="text" icon="el-icon-close" @click="message=''"></el-button>
    </div>

    <div class="form-body">
      <form class="device-form">
        <div class="field-title block-title">公共字段</div>
        <float-label-input v-model="form.name" placeholder="设备名称"></float-label-input>
        <float-label-input v-model="form.type" placeholder="设备类型"></float-label-input>
        <div class="field-title">选择设备生产的产品</div>
        <el-select
          remote
          filterable
          clearable
          v-model="form.productID"
          placeholder="搜索产品"
          :remote-method="searchProduct"
          :searchLoading="searchLoading"
        >
          <el-option
            v-for="(p, index) in products"
            :key="'product_' + index"
            :label="p.name"
            :value="p.id"
          ></el-option>
        </el-select>

        <div class="field-title">批量创建设备</div>
        <div class="flex-form-item batch-num">
          <el-input-number v-model="num" @change="changeNum" :min="1" label="批量创建"></el-input-number>
          <span style="padding-left: 1rem;">/</span>
          <span style="padding-left: 1rem;">台</span>
        </div>

        <div class="field-title block-title">私有字段</div>

        <div
          v-for="(form, index) in privateForms"
          :key="'private-form_' + index"
          class="flex-form-item private-form-item"
        >
          <float-label-input v-model="form.address" placeholder="设备地址"></float-label-input>
          <span>-</span>
          <float-label-input v-model="form.number" placeholder="设备编号"></float-label-input>
          <span>设备 {{index + 1}}</span>
        </div>

        <el-button class="submit-btn" type="primary" @click="submit" :loading="loading">提交</el-button>
      </form>
    </div>
  </div>
</template>
<script>
import FloatLabelInput from 'js/vue/main/components/float-label-input';
import productListQuery from './query.product-list.gql';
import deviceCreateMutate from './mutate.device-create.gql';

export default {
  name: 'device-new',
  components: {
    FloatLabelInput
  },
  data() {
    return {
      num: 1,
      message: '',
      products: [],
      loading: false,
      searchLoading: false,
      form: {
        type: '',
        name: '',
        productID: undefined
      },
      privateForms: [{ address: '', number: '' }]
    };
  },
  methods: {
    searchProduct(query) {
      this.searchLoading = true;
      if (query !== '') {
        this.$apollo
          .query({
            query: productListQuery,
            variables: { query }
          })
          .then(({ data }) => {
            this.products = data.productList.products;
            this.searchLoading = false;
          })
          .catch(e => {
            this.searchLoading = false;
            this.products = [];
            console.log(e);
          });
      } else {
        this.searchLoading = false;
        this.products = [];
      }
    },
    changeNum(currentValue, oldValue) {
      if (currentValue > oldValue) {
        for (var i = 0; i < currentValue - oldValue; i++) {
          this.privateForms.push({ address: '', number: '' });
        }
      } else if (currentValue < oldValue && this.privateForms.length > 1) {
        for (var i = 0; i < oldValue - currentValue; i++) {
          this.privateForms.pop();
        }
      }
    },
    submit() {
      this.loading = true;
      if (this.form.name === '') {
        this.message = '请填写设备名称';
        this.loading = false;
        document.getElementById('app').scrollTo(0, 0);
        return;
      }

      if (this.form.type === '') {
        this.message = '请填写设备类型';
        this.loading = false;
        document.getElementById('app').scrollTo(0, 0);
        return;
      }

      this.$apollo
        .mutate({
          mutation: deviceCreateMutate,
          variables: {
            ...this.form,
            privateForms: this.privateForms
          }
        })
        .then(({ data }) => {
          this.$message({
            type: 'success',
            message: '成功创建' + data.deviceCreate + '台设备'
          });
          this.$router.push({
            name: 'device-list',
            query: { search: this.form.name, self: true }
          });
        })
        .catch(e => {
          this.message = e.message;
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

.device-new {
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

  .form-body .device-form {
    width: 520px;
    margin-right: auto;
  }

  .device-form .field-title {
    border-bottom: 1px solid $--color-theme__light;
    margin-bottom: 1rem;
    padding-bottom: 0.5rem;
  }

  .device-form .field-title.block-title {
    font-size: 1.2rem;
    font-weight: bold;
    margin: 2rem 0;
  }

  .device-form .el-select {
    margin-bottom: 15px;
    width: 200px;
  }

  .device-form .el-input-number {
    width: 200px;
  }

  .device-form .flex-form-item {
    display: flex;
    flex-wrap: wrap;
  }

  .device-form .flex-form-item.batch-num {
    align-items: baseline;

    .el-input-number .el-input-number__decrease,
    .el-input-number .el-input-number__increase {
      bottom: 1px;
      display: flex;
      align-items: center;

      i {
        flex: auto;
      }
    }
  }

  .device-form .flex-form-item.private-form-item {
    justify-content: space-between;
    align-items: center;
    margin-bottom: 1rem;

    .floating-label {
      width: 200px;
      margin-bottom: 0;
    }
  }

  .device-form .submit-btn {
    width: 100%;
    margin-top: 1rem;
  }
}
</style>
