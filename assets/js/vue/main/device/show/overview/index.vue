<template>
  <div
    class="device-overview"
    v-loading="$apollo.queries.device.loading"
    element-loading-background="unset"
  >
    <div class="overview__row row">
      <div class="overview__col global-card">
        <div style="float: left; width: 50%">
          <div class="col-title">运行状态</div>
          <div class="col-content">
            <span class="label" v-if="device.status">{{statusMap[device.status].label}}</span>
          </div>
        </div>

        <i :class="['iconfont', statusMap[device.status].icon]" v-if="device.status"></i>
      </div>

      <div class="overview__col global-card">
        <div class="col-title">稼动率</div>
        <div class="col-content">
          <span v-if="device.statistics">{{ device.statistics.activation}}</span>
          <span v-else>0</span>
          %
        </div>
      </div>

      <div class="overview__col global-card">
        <div class="col-title">良率</div>
        <div class="col-content">
          <span v-if="device.statistics">{{ device.statistics.yield}}</span>
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
          <div class="value" v-if="device.user">
            <img
              class="avatar"
              :src="device.user.avatarURL ? device.user.avatarURL : defaultAvatar"
            />
            <div>{{ device.user.name }}</div>（
            <div style="color: #909399">{{ device.user.phone }}</div>）
          </div>
        </div>

        <div class="col-line">
          <div class="label">描述</div>
          <click-to-edit class="value" @save="save('description')">
            <template v-slot:text>{{ device.description ? device.description : '[点击填写]' }}</template>
            <template v-slot:form>
              <el-input v-model="device.description" size="mini"></el-input>
            </template>
          </click-to-edit>
        </div>
      </div>

      <div class="details__col"></div>
    </div>
  </div>
</template>
<script>
// graphql
import deviceOverviewQuery from '../gql/query.device-overview.gql';
import deviceUpdateMutate from '../gql/mutate.device-update.gql';

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
      statusMap: {
        prod: { icon: 'icon-running', label: '运行中' },
        stop: { icon: 'icon-stopping', label: '停机' },
        offline: { icon: 'icon-offline', label: '离线' }
      }
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

    .icon-running {
      color: $--color-theme__success;
    }

    .icon-stopping {
      color: $--color-theme__danger;
    }

    .icon-offline {
      color: $--color-theme__gray;
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

  .device-details {
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

    .value .avatar {
      display: inline-block;
      height: 30px;
      width: 30px;
      border-radius: 2px;
      vertical-align: middle;
      margin-right: 0.5rem;
    }

    .label:after {
      content: ':';
      padding: 0 5px;
    }
  }
}
</style>
