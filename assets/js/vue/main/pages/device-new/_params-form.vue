<template>
  <div class="device-form params-form">
    <el-form label-position="left" label-width="100px">
      <el-form-item label="设备参数">
        <el-tag
          v-for="param in params"
          :key="param.sign"
          closable
          @close="removeParam(param)"
          @click="edit(param)"
        >{{param.name || param.sign}}</el-tag>
        <el-button
          style="margin-right: 1rem"
          @click="formVisible = true"
          v-show="!formVisible"
          size="small"
          type="primary"
        >增加</el-button>
        <span v-if="params.length == 0">暂无参数</span>
      </el-form-item>
    </el-form>

    <el-form
      label-position="left"
      :model="form"
      :rules="rules"
      status-icon
      label-width="100px"
      v-if="formVisible"
      ref="paramForm"
    >
      <el-form-item label="参数名称" prop="name">
        <el-input v-model="form.name"></el-input>
      </el-form-item>

      <el-form-item label="参数标识" prop="sign">
        <el-input v-model="form.sign"></el-input>
      </el-form-item>

      <el-form-item label="参数值类型" prop="type">
        <el-select v-model="form.type" placeholder="请选择">
          <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
          ></el-option>
        </el-select>
      </el-form-item>

      <el-form-item>
        <el-button size="small" type="primary" @click="saveParam()">保存</el-button>
        <el-button size="small" type="warning" @click="cancelParam()">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: 'params-form',
  data() {
    var uniqueSign = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('参数签名为必填字段!'));
      }

      var index = this.params.find(param => {
        return param.sign === value;
      });

      if (index) return callback(new Error('参数签名不可以重复使用!'));

      callback();
    };

    return {
      options: [
        { value: 'string', label: '字符串' },
        { value: 'bool', label: '布尔值' },
        { value: 'int', label: '整数值' },
        { value: 'float', label: '浮点数' }
      ],
      params: [],
      form: {
        name: '',
        sign: '',
        type: 'string'
      },
      rules: {
        sign: [{ validator: uniqueSign, trigger: 'blur' }]
      },
      formVisible: false
    };
  },
  methods: {
    removeParam(param) {
      var index = this.params.findIndex(p => p.sign === param.sign);
      this.params.splice(index, 1);
    },
    saveParam() {
      this.$refs.paramForm.validate(v => {
        if (v) {
          var value = { ...this.form };
          this.params.push(value);
          this.formVisible = false;
          this.$refs.paramForm.resetFields();
        } else {
          return false;
        }
      });
    },
    cancelParam() {
      this.formVisible = false;
      this.$refs.paramForm.resetFields();
    },
    edit(param) {
      this.form = {
        ...param
      };
      this.formVisible = true;
    }
  }
};
</script>
