<template>
  <div class="device-new">
    <div class="form-body global-card">
      <div class="form-title">
        <i class="el-icon-back go-back-btn" @click="$router.go(-1)"></i>
        <i class="el-icon-s-order"></i> 注册设备
      </div>

      <div class="device-form">
        <el-form label-position="left" :model="form" label-width="100px" :rules="rules" ref="form">
          <el-form-item label="设备名称" prop="name">
            <el-input v-model="form.name"></el-input>
          </el-form-item>

          <el-form-item label="设备类型" prop="type">
            <el-input v-model="form.type"></el-input>
          </el-form-item>

          <el-form-item label="设备描述" prop="description">
            <el-input type="textarea" :rows="10" v-model="form.description"></el-input>
          </el-form-item>

          <el-form-item style="text-align: center">
            <el-button :loading="loading" @click="submit()" style="width: 100%;" type="primary">提交</el-button>
          </el-form-item>
        </el-form>
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
        name: '测试设备',
        type: '测试',
        description: '测试注册设备交互是否正常'
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
              this.$router.push({
                name: 'device-show',
                params: { uuid: data.device.uuid }
              });
            })
            .catch(e => console.error(e));
        }
      });
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/device/new.scss';
</style>
