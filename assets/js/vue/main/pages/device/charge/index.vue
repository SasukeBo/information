<template>
  <div class="device-charge">
    <el-tooltip effect="dark" content="增加负责人" placement="left-end">
      <i class="el-icon-plus add-charge-btn"></i>
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
<style lang="scss">
@import 'css/vars.scss';

.device-charge {
  margin: 1rem;
  margin-top: 3rem;
  position: relative;
}

.device-charge .add-charge-btn {
  font-size: 1.2rem;
  font-weight: bolder;
  padding: 1rem;
  border-radius: 50%;
  background: $--color-theme__main;
  box-shadow: 0 0 3px darken($--color-theme__main, 80%),
    0 0 5px $--color-theme__white;
  position: absolute;
  transition: background 0.3s ease;
  top: -24px;
  right: 32px;
  cursor: pointer;
  margin: auto;
  z-index: 1;

  &:hover {
    background: darken($--color-theme__main, 15%);
  }
}

.device-charge .charge-body {
  padding-left: 1rem;
  padding-right: 1rem;

  .empty-tip {
    text-align: center;
  }
}

.device-charge .charge-search {
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid $--color-theme__black;

  .el-input {
    max-width: 300px;
  }
}

.device-charge .charge-item {
  margin-bottom: 1rem;

  &:last-child {
    margin-bottom: 0;
  }
}

@media only screen and (max-width: 700px) {
  .device-charge .add-charge-btn {
    right: 0;
  }
}

@media only screen and (max-width: 400px) {
  .device-charge .add-charge-btn {
    right: calc(50% - 25px);
  }

  .device-charge .charge-search .el-input {
    max-width: 100%;
  }

  .device-charge .charge-body {
    padding-top: 2rem;
  }
}
</style>
