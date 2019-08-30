<template>
  <div class="charge-item" v-if="charge">
    <div class="charge-item__avatar">
      <img :src="charge.user.avatarURL || '/images/avatar.jpg'" />
    </div>

    <div class="name-phone">
      <span class="name">{{ charge.user.userExtend.name }}</span>
      <span class="phone">
        <i class="el-icon-phone"></i>
        {{ charge.user.phone }}
      </span>
    </div>

    <div class="create-at">
      <div class="label">指派日期：</div>
      <div>{{ timeFormatter(charge.createdAt) }}</div>
    </div>

    <a
      class="to-show link"
      :href="'charge/' + charge.id + '/show'"
      @click.prevent="$router.push({name: 'charge-show', params: { id: charge.id }})"
    >查看</a>

    <a
      class="to-edit link"
      :href="'charge/' + charge.id + '/edit'"
      @click.prevent="$router.push({ name: 'charge-edit', params: {id: charge.id}})"
    >编辑</a>

    <div class="delete-btn">
      <el-button size="small" type="danger" @click="deleteCharge">删除</el-button>
    </div>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';
import chargsQuery from './gql/query.charges.gql';
import deviceChargeDelete from './gql/mutation.device-charge-delete.gql';

export default {
  name: 'charge-item',
  props: ['charge'],
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    },
    deleteCharge() {
      this.$apollo
        .mutate({
          mutation: deviceChargeDelete,
          variables: { id: this.charge.id },
          update: (store, { data: { id } }) => {
            var opts = {
              query: chargsQuery,
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
