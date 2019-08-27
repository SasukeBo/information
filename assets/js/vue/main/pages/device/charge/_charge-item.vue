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
<style lang="scss" scoped>
@import 'css/vars.scss';

.charge-item {
  display: flex;
  flex-flow: wrap;
  background: $--color-theme__light-black;
  padding: 0.5rem;
  border-radius: 4px;
}

.charge-item .charge-item__avatar {
  flex: 1;
  text-align: left;
  margin-right: 1rem;

  img {
    width: 50px;
    height: 50px;
    border-radius: 4px;
  }
}

.charge-item .name-phone {
  flex: 3;
  padding-right: 1rem;
  line-height: 50px;

  .name {
    font-size: 1.2rem;
    padding-bottom: 0.5rem;
    padding-right: 0.5rem;
  }
}

.charge-item .create-at {
  flex: 3;

  .label {
    font-size: 0.8rem;
    padding-bottom: 0.5rem;
  }
}

.charge-item .link {
  flex: 1;
  line-height: 50px;
}

.charge-item .delete-btn {
  flex: 1;

  &:before {
    content: '';
    display: inline-block;
    vertical-align: middle;
    height: 50px;
  }

  .el-button {
    vertical-align: middle;
  }
}

@media only screen and (max-width: 700px) {
  .charge-item {
    display: block;
    padding: 0.5rem 1rem;
  }

  .charge-item .charge-item__avatar {
    display: inline-block;
    vertical-align: middle;

    img {
      width: 100px;
      height: 100px;
    }
  }

  .charge-item .name-phone {
    display: inline-block;
    vertical-align: middle;
    line-height: 2rem;

    .name {
      display: block;
    }

    .phone {
      display: block;
    }
  }

  .charge-item .create-at {
    margin-top: 0.5rem;
  }

  .charge-item .create-at div {
    display: inline-block;
  }

  .link {
    display: inline-block;
    margin-right: 1rem;
  }

  .delete-btn {
    display: inline-block;
  }
}
</style>
