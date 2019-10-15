<template>
  <div class="charger-item" v-if="charger">
    <div class="charger-item__avatar">
      <img :src="charger.user.avatarURL || '/images/avatar.jpg'" />
    </div>

    <div class="name-phone">
      <span class="name">{{ charger.user.userExtend.name }}</span>
      <span class="phone">
        <i class="el-icon-phone"></i>
        {{ charger.user.phone }}
      </span>
    </div>

    <div class="create-at">
      <div class="label">指派日期：</div>
      <div>{{ timeFormatter(charger.createdAt) }}</div>
    </div>

    <div class="privs-count">
      <span>权限数量：{{ charger.privs.length }}</span>
    </div>

    <a
      class="to-show link"
      :href="'charger/' + charger.id + '/show'"
      @click.prevent="$router.push({name: 'charger-show', params: { id: charger.id }})"
    >查看</a>

    <div class="delete-btn">
      <el-button size="small" type="danger" @click="deleteCharge">删除</el-button>
    </div>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';
import chargersQuery from './gql/query.chargers.gql';
import deviceChargerDelete from './gql/mutation.device-charger-delete.gql';

export default {
  name: 'charger-item',
  props: ['charger'],
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    },
    deleteCharge() {
      this.$apollo
        .mutate({
          mutation: deviceChargerDelete,
          variables: { id: this.charger.id },
          update: (store, { data: { id } }) => {
            var opts = {
              query: chargersQuery,
              variables: { uuid: this.$route.params.uuid }
            };
            var data = store.readQuery(opts);
            var index = data.charges.findIndex(c => c.id === id);
            data.charges.splice(index, 1);
            store.writeQuery({ ...opts, data });
          }
        })
        .then(() => {
          this.$message({ type: 'success', message: '操作成功' });
        })
        .catch(e => console.error(e));
    }
  }
};
</script>
