<template>
  <div class="device-new">
    <div class="form-body">
      <div class="form-title">
        <i class="el-icon-back go-back-btn" @click="$router.go(-1)"></i>
        <i class="el-icon-s-order"></i> 注册设备
      </div>

      <div class="device-form">
        <el-form label-position="left" :model="form" label-width="100px" :rules="rules" ref="form">
          <el-form-item prop="name">
            <el-input v-model="form.name" placeholder="填写设备名称"></el-input>
          </el-form-item>

          <el-form-item prop="type">
            <el-input v-model="form.type" placeholder="填写设备类型"></el-input>
          </el-form-item>

          <el-form-item prop="address">
            <el-input placeholder="填写设备地址" v-model="form.address"></el-input>
          </el-form-item>

          <el-form-item prop="description">
            <el-input type="textarea" placeholder="填写设备描述" :rows="10" v-model="form.description"></el-input>
          </el-form-item>
        </el-form>
      </div>
    </div>

    <div class="deploy-order">
      <div class="deploy-order__body">
        <span class="label">批量创建</span>
        <!-- 此处注释目的是除去inline-block之间的空格间隙
        --><el-input-number v-model="form.count" :min="1"></el-input-number>
        <div style="flex: auto"></div>
        <el-button :loading="loading" type="primary" @click="submit">立即注册</el-button>
      </div>
    </div>
  </div>
</template>
<script>
import deviceCreateMutation from './gql/mutation.device-create.gql';

export default {
  name: 'device-new',
  data() {
    return {
      loading: false,
      rules: {
        name: [{ required: true, message: '设备名称必填', trigger: 'blur' }],
        type: [{ required: true, message: '设备类型必填', trigger: 'blur' }]
      },
      form: {
        count: 1,
        name: '',
        type: '',
        address: '',
        description: ''
      }
    };
  },
  methods: {
    submit() {
      this.$refs.form.validate(valid => {
        if (valid) {
          this.$apollo
            .mutate({
              mutation: deviceCreateMutation,
              variables: this.form
            })
            .then(({ data }) => {
              this.loading = false;
              this.$router.push({ name: 'device-list' });
            })
            .catch(e => console.error(e));
        }
      });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/main/device/new.scss';
</style>
