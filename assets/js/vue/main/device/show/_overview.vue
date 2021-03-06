<template>
  <div
    class="device-overview"
    v-loading="$apollo.queries.device.loading"
    element-loading-background="unset"
  >
    <div class="overview__row row">
      <div class="overview__col global-card">
        <div class="col-title">OEE</div>
        <div class="col-content">
          <span
            v-if="device.statistics && device.statistics.oee"
          >{{ (device.statistics.oee * 100).toFixed(2) }}</span>
          <span v-else>0</span>
          %
        </div>
      </div>

      <div class="overview__col global-card">
        <div class="col-title">稼动率</div>
        <div class="col-content">
          <span
            v-if="device.statistics && device.statistics.availability"
          >{{ (device.statistics.availability * 100).toFixed(2) }}</span>
          <span v-else>0</span>
          %
        </div>
      </div>

      <div class="overview__col global-card">
        <div class="col-title">良率</div>
        <div class="col-content">
          <span
            v-if="device.statistics && device.statistics.quality"
          >{{ (device.statistics.quality * 100).toFixed(2) }}</span>
          <span v-else>0</span>
          %
        </div>
      </div>
    </div>

    <div class="device-details row">
      <div class="details__col">
        <div class="col-line">
          <div class="label">名称</div>
          <click-to-edit class="value" @save="save('name')">
            <template v-slot:text>{{ device.name ? device.name : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="device.name" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="col-line">
          <div class="label">编号</div>
          <click-to-edit class="value" @save="save('number')">
            <template v-slot:text>{{ device.number ? device.number : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="device.number" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="col-line">
          <div class="label">类型</div>
          <click-to-edit class="value" @save="save('type')">
            <template v-slot:text>{{ device.type ? device.type : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="device.type" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="col-line">
          <div class="label">Token</div>
          <div class="value" style="color: #fff;font-weight: bold">{{ device.token }}</div>
        </div>
      </div>

      <div class="details__col">
        <div class="col-line">
          <div class="label">创建时间</div>
          <div class="value">{{ timeFormatter(device.createdAt) }}</div>
        </div>

        <div class="col-line">
          <div class="label">物理位置</div>
          <click-to-edit class="value" @save="save('address')">
            <template v-slot:text>{{ device.address ? device.address : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="device.address" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>

        <div class="col-line">
          <div class="label">注册人</div>
          <div class="value register" v-if="device.user">
            <img
              class="avatar"
              :src="device.user.avatarURL ? device.user.avatarURL : defaultAvatar"
            />
            <div>{{ device.user.name }}</div>
            <div style="color: #909399">{{ device.user.phone }}</div>
          </div>
        </div>

        <div class="col-line">
          <div class="label">生产产品</div>
          <div class="value" v-if="device.product">
            <a
              :href="'/product/' + device.product.id + '/show'"
              target="_blank"
            >{{device.product.name}}</a>
          </div>
        </div>
      </div>

      <div class="details__col">
        <div class="col-line">
          <div class="label" style="width: 110px;">最大生产速率</div>
          <click-to-edit class="value" @save="save('prodSpeed')">
            <template v-slot:text>{{ device.prodSpeed ? device.prodSpeed : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input-number v-model="device.prodSpeed" :precision="2" :step="1" size="mini"></el-input-number>
            </template>
          </click-to-edit>
        </div>
      </div>
    </div>
  </div>
</template>
  <script>
// graphql
import deviceOverviewQuery from './gql/query.device-overview.gql';
import deviceUpdateMutate from './gql/mutate.device-update.gql';

import defaultAvatar from 'images/default-avatar.png';
import { timeFormatter, parseGQLError } from 'js/utils';
import ClickToEdit from 'js/vue/main/components/click-to-edit';

export default {
  name: 'device-overview',
  props: ['id'],
  components: { ClickToEdit },
  apollo: {
    device: {
      query: deviceOverviewQuery,
      variables() {
        return { id: this.id };
      }
    }
  },
  data() {
    return {
      device: {},
      defaultAvatar,
      statusUpdater: undefined
    };
  },
  methods: {
    timeFormatter(str) {
      return timeFormatter(str, '%y年%m月%d日');
    },
    save(field) {
      var variables = { id: this.id };
      variables[field] = this.device[field];
      this.$apollo
        .mutate({
          mutation: deviceUpdateMutate,
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
  },
  beforeDestroy() {
    clearInterval(this.statusUpdater);
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-overview {
  .overview__col {
    width: 30%;
    padding: 24px;
    border-radius: 2px;
    margin-bottom: 0;
    box-shadow: none;

    .iconfont {
      float: left;
      width: 50%;
      text-align: center;
      line-height: 79px;
      font-size: 4rem;
      padding-left: 1rem;
    }
  }

  .overview__col .col-title {
    margin-bottom: 8px;
    color: $--color-font__gray;
    font-size: 15px;
  }

  .overview__col .col-content {
    color: $--color-theme__main;
    font-size: 36px;
    line-height: 50px;
  }

  .device-details.row {
    margin-top: 30px;
    color: $--color-font__light;
  }

  .device-details .details__col {
    width: 30%;

    .col-line {
      display: flex;
      min-height: 40px;
      align-items: center;
    }

    .label {
      width: 80px;
      color: $--color-font__gray;
    }

    .value {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      flex: auto;
    }

    .value.register {
      position: relative;
      display: block;
    }

    .value .avatar {
      height: 30px;
      width: 30px;
      border-radius: 2px;
      margin-right: 0.5rem;
      float: left;
    }

    .label:after {
      content: ':';
      padding: 0 5px;
    }
  }
}
</style>
