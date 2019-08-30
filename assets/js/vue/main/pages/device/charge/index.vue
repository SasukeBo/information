<template>
  <div class="device-charge">
    <el-tooltip effect="dark" content="增加负责人" placement="left-end">
      <i class="el-icon-plus add-charge-btn" @click="$router.push({ name: 'charge-new' })"></i>
    </el-tooltip>

    <div class="global-card charge-body">
      <div class="charge-search">
        <el-input v-model="search" placeholder="搜索名称" prefix-icon="el-icon-search"></el-input>
      </div>

      <charge-item
        @remove="handleRemove"
        v-for="charge in charges"
        :key="charge.id"
        :charge="charge"
      ></charge-item>
      <div v-if="charges.length === 0" class="empty-tip">暂无负责人</div>
    </div>
  </div>
</template>
<script>
import { Tooltip } from 'element-ui';
import ChargeItem from './_charge-item';
import chargesQuery from './gql/query.charges.gql';

export default {
  name: 'device-charge',
  props: ['uuid'],
  components: {
    ElTooltip: Tooltip,
    ChargeItem
  },
  apollo: {
    charges: {
      query: chargesQuery,
      variables() {
        return { uuid: this.uuid };
      }
    }
  },
  data() {
    return {
      charges: [],
      search: ''
    };
  },
  methods: {
    handleRemove(id) {
      var index = this.charges.findIndex(c => c.id === id);
      this.charges.splice(index, 1);
    }
  }
};
</script>
