<template>
  <div class="device-params">
    <el-tooltip effect="dark" content="增加参数" placement="left-end">
      <i class="el-icon-plus add-param-btn" @click="addParam"></i>
    </el-tooltip>

    <div class="global-card params-body">
      <div class="responsive-table">
        <div class="table-header">
          <div>名称</div>
          <div>标识</div>
          <div>值类型</div>
          <div>创建时间</div>
          <div>创建人</div>
          <div style="width: 100px"></div>
        </div>

        <param-item
          @cancel="cancelParam"
          @save="handleSave"
          v-if="showFormItem"
          key="new-item"
          :uuid="uuid"
        ></param-item>

        <param-item
          @remove="removeItem"
          v-for="param in deviceParams"
          :key="param.id"
          :param="param"
        ></param-item>
      </div>
    </div>
  </div>
</template>
<script>
import { Tooltip } from 'element-ui';
import ParamItem from './_param-item';
import deviceParamsQuery from './gql/query.params.gql'

export default {
  name: 'device-params',
  props: ['uuid'],
  components: {
    ParamItem,
    ElTooltip: Tooltip
  },
  apollo: {
    deviceParams: {
      query: deviceParamsQuery,
      variables() {
        return this.queryVariables;
      }
    }
  },
  data() {
    return {
      deviceParams: [],
      queryVariables: {
        deviceUUID: this.uuid,
        namePattern: ''
      },
      showFormItem: false
    };
  },
  methods: {
    addParam() {
      this.showFormItem = true;
    },
    cancelParam() {
      this.showFormItem = false;
    },
    handleSave(param) {
      this.showFormItem = false;
    },
    removeItem(id) {
      var index = this.params.findIndex(p => p.id === id);
      this.params.splice(index, 1);
    }
  }
};
</script>
