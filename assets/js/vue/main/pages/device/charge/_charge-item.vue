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

export default {
  name: 'charge-item',
  props: ['charge'],
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    },
    deleteCharge() {
      console.log('click delete button');
    }
  }
};
</script>
