<template>
  <div class="product-overview">
    <div class="box-row">
      <a
        class="box-link global-card"
        href="product/5/show?tab=devices"
        @click.prevent="$router.push({name: 'product-show', query: {tab: 'devices'}})"
      >
        <div class="box-container">
          <div class="box-title">生产设备总数</div>
          <div class="box-content">
            <span class="content">100台</span> /
            <span class="content__append">86台再生产</span>
          </div>
        </div>
      </a>

      <a
        class="box-link global-card"
        href="product/5/show?tab=instances"
        @click.prevent="$router.push({name: 'product-show', query: {tab: 'instances'}})"
      >
        <div class="box-container">
          <div class="box-title">生产总数</div>
          <div class="box-content">
            <span class="content">65536个</span> /
            <span class="content__append">良率90%</span>
          </div>
        </div>
      </a>

      <a
        class="box-link global-card"
        href="product/5/show?tab=instances"
        @click.prevent="$router.push({name: 'product-show', query: {tab: 'instances'}})"
      >
        <div class="box-container">
          <div class="box-title">今日生产总数</div>
          <div class="box-content">
            <span class="content">1024个</span> /
            <span class="content__append">良率99%</span>
          </div>
        </div>
      </a>
    </div>

    <div class="detail-row">
      <div class="detail-group">
        <div class="item">
          <span class="label">注册人</span>
          <div class="value" v-if="product.register">
            <img
              class="register-avatar"
              :src="product.register.avatarURL ? product.register.avatarURL : defaultAvatar"
            />
            <span>{{ product.register.name }}</span>
            (
            <span style="color: #8a9099">{{product.register.phone}}</span> )
          </div>
          <div class="value" v-else>未知</div>
        </div>

        <div class="item">
          <span class="label">注册时间</span>
          <span class="value">{{ timeFormatter(product.createdAt) }}</span>
        </div>

        <div class="item">
          <span class="label">Token</span>
          <span class="value">{{ product.token }}</span>
        </div>

        <div class="item">
          <span class="label">生产负责人</span>
          <click-to-edit class="value" @save="save('productor')">
            <template v-slot:text>{{ product.productor ? product.productor : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.productor" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="item">
          <span class="label">负责人联系电话</span>
          <click-to-edit class="value" @save="save('productorContact')">
            <template
              v-slot:text
            >{{ product.productorContact ? product.productorContact : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.productorContact" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>
      </div>

      <div class="detail-group">
        <div class="item">
          <span class="label">订单编号</span>
          <click-to-edit class="value" @save="save('orderNum')">
            <template v-slot:text>{{ product.orderNum ? product.orderNum : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.orderNum" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="item">
          <span class="label">需求方</span>
          <click-to-edit class="value" @save="save('customer')">
            <template v-slot:text>{{ product.customer ? product.customer : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.customer" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="item">
          <span class="label">需求方联系电话</span>
          <click-to-edit class="value" @save="save('customerContact')">
            <template v-slot:text>{{ product.customerContact ? product.customerContact : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.customerContact" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="item">
          <span class="label">需求生产总量</span>
          <click-to-edit class="value" @save="save('total')">
            <template v-slot:text>{{ product.total ? product.total : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="product.total" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="item">
          <span class="label">计划完成于</span>
          <click-to-edit class="value" @save="save('finishTime')">
            <template
              v-slot:text
            >{{ product.finishTime ? timeFormatter(product.finishTime) : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-date-picker
                v-model="product.finishTime"
                size="mini"
                type="date"
                placeholder="选择日期"
                prefix-icon="none"
                :editable="false"
                :clearable="false"
              ></el-date-picker>
            </template>
          </click-to-edit>
        </div>
      </div>

      <div
        class="detail-group detect-items"
        v-if="product.detectItems && product.detectItems.length"
      >
        <div
          class="item"
          v-for="item in product.detectItems"
          :key="item.sign"
          style="overflow: hidden;"
        >
          <span class="label" style="width: 60px;">检测项</span>
          <div class="value" style="color: #c0c4cc;">
            <span style="color: #fff; font-weight: bold; padding-right: 0.5rem">{{ item.sign }}</span>
            <i class="iconfont icon-shangxiankongzhi" style="color: #f06d6b"></i>
            <span style="padding-right: 0.5rem;">{{ item.upperLimit }}</span>
            <i class="iconfont icon-xiaxiankongzhi" style="color: #606266"></i>
            <span>{{ item.lowerLimit }}</span>
          </div>
        </div>

        <a
          class="edit-detect-items"
          :href="'/product/' + id + '/edit'"
          @click.prevent="$router.push({name: 'product-show', query: {tab: 'setting'}, hash: '#detect-items'})"
        >修改检测项</a>
      </div>
    </div>
  </div>
</template>
<script>
import productDetailsQuery from './gql/query.product-details.gql';
import productUpdateMutate from './gql/mutate.product-update.gql';
import { timeFormatter, parseGQLError } from 'js/utils';
import defaultAvatar from 'images/default-avatar.png';
import ClickToEdit from 'js/vue/main/components/click-to-edit';

export default {
  name: 'product-overview',
  props: ['id'],
  components: { ClickToEdit },
  apollo: {
    product: {
      query: productDetailsQuery,
      variables() {
        return {
          id: this.id
        };
      }
    }
  },
  data() {
    return {
      product: {},
      defaultAvatar
    };
  },
  methods: {
    timeFormatter(str) {
      return timeFormatter(str, '%y年%m月%d日');
    },
    save(field) {
      var variables = { id: this.id };
      variables[field] = this.product[field];
      this.$apollo
        .mutate({
          mutation: productUpdateMutate,
          variables
        })
        .then(() => {
          this.$message({ type: 'success', message: '更新成功' });
        })
        .catch(e => {
          var err = parseGQLError(e);
          this.$message({ type: 'error', message: err.message });
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

.product-overview {
  .box-row {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
  }

  .box-link {
    width: 30%;
    margin-bottom: 1rem;
    cursor: pointer;
    box-shadow: unset;
    transition: box-shadow ease-in 0.15s;

    &:hover {
      box-shadow: $--shadow__global-card;
    }
  }

  .box-container {
    transition: all ease-in 0.15s;
    height: 120px;
    padding: 24px;
    border-radius: 2px;

    .box-title {
      color: $--color-font__gray;
      margin-bottom: 8px;
    }

    .box-content .content {
      display: inline-block;
      font-weight: 300;
      font-size: 36px;
      line-height: 36px;
      color: $--color-theme__main;
    }

    .box-content .content__append {
      display: inline-block;
      color: $--color-theme__success;
    }
  }

  .detail-row {
    margin-top: 30px;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;

    .register-avatar {
      display: inline-block;
      height: 30px;
      width: 30px;
      border-radius: 2px;
      vertical-align: middle;
      margin-right: 0.5rem;
    }
  }

  .detail-row .detail-group {
    width: 30%;
  }

  .detail-row .item {
    min-height: 40px;
    display: flex;
    color: $--color-font__light;
    align-items: center;

    .label {
      display: inline-block;
      width: 110px;
      color: $--color-font__gray;

      &:after {
        content: ':';
        display: inline-block;
      }
    }

    .value {
      flex: auto;
    }

    .value.click-to-edit .el-input {
      width: 100px;
    }
  }

  .detail-row .click-to-edit .el-date-editor .el-input__inner {
    padding: 0 0.8rem;
    text-align: center;
  }

  .detail-row .detect-items {
    position: relative;

    &:hover .edit-detect-items {
      display: inline;
    }

    .edit-detect-items {
      display: none;
      position: absolute;
      right: 0;
      top: -15px;
      font-size: 12px;
    }
  }
}
</style>
