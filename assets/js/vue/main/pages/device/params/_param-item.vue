<template>
  <div class="param-item">
    <div class="table-cell" style="width: 15%">
      <span class="span-block" v-if="!form.edit">{{ form.name }}</span>
      <el-input autofocus v-model="form.name" v-else @keyup.enter.native="save"></el-input>
    </div>

    <div class="table-cell" style="width: 15%">
      <span class="span-block" v-if="!form.edit">{{ form.sign }}</span>
      <el-input v-model="form.sign" v-else @keyup.enter.native="save"></el-input>
    </div>

    <div class="table-cell" style="width: 15%">
      <span class="span-block" v-if="!form.edit">{{ typeMap[form.type] }}</span>
      <el-select v-model="form.type" v-else placeholder="请选择">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>

    <div class="table-cell" style="width: 20%; white-space: nowrap">
      <span class="span-block">{{ timeFormatter(form.createdAt) }}</span>
    </div>

    <div class="table-cell" style="width: 15%">
      <span class="span-block">{{ form.author ? form.author.userExtend.name : '-' }}</span>
    </div>

    <div class="table-cell" style="width: 20%">
      <transition mode="out-in" name="slide-fade">
        <div v-if="!form.edit" key="delete">
          <el-button type="primary" size="small" @click="form.edit = true">编辑</el-button>
          <el-button type="danger" size="small" @click="remove" :loading="deleting">删除</el-button>
        </div>
        <div v-else key="save">
          <el-button type="primary" size="small" :loading="saving" @click="save">保存</el-button>
          <el-button type="warning" size="small" @click="cancel">取消</el-button>
        </div>
      </transition>
    </div>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';
import paramDelete from './gql/mutation.param-delete.gql';
import paramUpdate from './gql/mutation.param-update.gql';
import paramsQuery from './gql/query.params.gql';

export default {
  name: 'param-item',
  props: ['param', 'uuid'],
  data() {
    return {
      saving: false,
      deleting: false,
      form: {},
      options: [
        { value: 'string', label: '字符串' },
        { value: 'bool', label: '布尔值' },
        { value: 'int', label: '整数值' },
        { value: 'float', label: '浮点数' }
      ],
      typeMap: {
        string: '字符串',
        bool: '布尔值',
        int: '整数值',
        float: '浮点数'
      },
      paramListQueryOpts: {
        query: paramsQuery,
        variables: this.$parent.queryVariables
      }
    };
  },
  watch: {
    param: {
      immediate: true,
      handler: function(newVal) {
        if (newVal) this.form = { edit: false, ...this.param };
      }
    }
  },
  created() {
    if (!this.param) this.reset();
  },
  methods: {
    timeFormatter(timeStr) {
      return timeFormatter(timeStr);
    },
    cancel() {
      if (!this.form.id) {
        this.reset();
        this.$emit('cancel');
        return;
      }

      this.form = { edit: false, ...this.param };
    },
    reset() {
      this.form = {
        edit: true,
        name: '',
        sign: '',
        type: ''
      };
    },
    remove() {
      this.$apollo
        .mutate({
          mutation: paramDelete,
          variables: { id: this.param.id },
          update: (store, { data: { id } }) => {
            var data = store.readQuery(this.paramListQueryOpts);
            var index = data.deviceParams.findIndex(p => p.id === id);
            data.deviceParams.splice(index, 1);
            store.writeQuery({ ...this.paramListQueryOpts, data });
          }
        })
        .then(data => {
          this.$message({ type: 'success', message: data });
        });
    },
    save() {
      if (this.form.id) {
        this.$apollo
          .mutate({
            mutation: paramUpdate,
            variables: this.form,
            update: (store, { data: { deviceParam } }) => {
              var data = store.readQuery(this.paramListQueryOpts);
              var index = data.deviceParams.findIndex(
                dp => dp.id === deviceParam.id
              );
              data.deviceParams[index] = deviceParam;
              store.writeQuery({ ...this.paramListQueryOpts, data });
            }
          })
          .then(() => {
            this.saving = false;
          })
          .catch(e => console.log(e));
      } else {
        this.$apollo
          .mutate({
            mutation: paramCreate,
            variables: {
              deviceUUID: this.uuid,
              ...this.form
            },
            update: (store, { data: { deviceParam } }) => {
              var data = store.readQuery(this.paramListQueryOpts);
              data.deviceParams.unshift(deviceParam);
              store.writeQuery({ ...this.paramListQueryOpts, data });
            }
          })
          .then(() => {
            this.saving = false;
            this.$emit('save');
            this.reset();
          })
          .catch(e => console.log(e));
      }
    }
  }
};
</script>
