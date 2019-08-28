<template>
  <div class="device-charge">
    <el-tooltip effect="dark" content="增加负责人" placement="left-end">
      <i class="el-icon-plus add-charge-btn" @click="$router.push({ name: 'charge-new' })"></i>
    </el-tooltip>

    <div class="global-card charge-body">
      <div class="charge-search">
        <el-input v-model="search" placeholder="搜索名称" prefix-icon="el-icon-search"></el-input>
      </div>

      <charge-item v-for="charge in charges" :key="charge.id" :charge="charge"></charge-item>
      <div v-if="charges.length === 0" class="empty-tip">暂无负责人</div>
    </div>
  </div>
</template>
<script>
import { apollo } from './graphql';
import { Tooltip } from 'element-ui';
import ChargeItem from './_charge-item';

export default {
  name: 'device-charge',
  props: ['uuid'],
  components: {
    ElTooltip: Tooltip,
    ChargeItem
  },
  apollo,
  data() {
    return {
      charges: [],
      search: ''
    };
  }
};
</script>
