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
import {
  deviceParamUpdate,
  deviceParamCreate,
  deviceParamDelete
} from './graphql';

export default {
  name: 'param-item',
  props: ['param', 'deviceID'],
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
      deviceParamDelete(this).then(() => {
        this.deleting = false;
        this.$emit('remove', this.param.id);
      });
    },
    save() {
      if (this.form.id) {
        deviceParamUpdate(this)
          .then(({ data: { deviceParamUpdate } }) => {
            this.saving = false;
            this.form = {
              edit: false,
              ...deviceParamUpdate
            };
          })
          .catch(e => console.log(e));
      } else {
        deviceParamCreate(this)
          .then(({ data: { deviceParamCreate } }) => {
            this.saving = false;
            this.reset();
            this.$emit('save', deviceParamCreate);
          })
          .catch(e => console.log(e));
      }
    }
  }
};
</script>
