<template>
  <div class="passport-form forget-form">
    <div class="form-title form-item">找回密码</div>
    <div class="form-alert form-item" v-if="show">{{ message }}</div>
    <div class="form-body">
      <el-input
        class="form-item"
        placeholder="手机号"
        v-model="phone"
        prefix-icon="iconfont icon-shouji"
      ></el-input>
      <div class="form-item securitycode">
        <el-input
          class="securitycode-input"
          placeholder="验证码"
          v-model="securityCode"
          prefix-icon="iconfont icon-securityCode-b"
        ></el-input>
        <el-button type="success" round class="securitycode-btn" @click="show = !show">获取验证码(60s)</el-button>
      </div>
      <el-input
        class="form-item"
        placeholder="新密码"
        v-model="password"
        type="password"
        show-password
        prefix-icon="iconfont icon-mima"
      ></el-input>
      <el-button class="form-item" type="primary" size="large" @click="show = !show">重置密码</el-button>
      <div class="link form-item">
        <a href="/login" @click.prevent="$router.push({path: 'login'})">
          <i class="iconfont icon-fanhui"></i>
          <span>返回登录</span>
        </a>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      show: false,
      phone: '',
      securityCode: '',
      password: '',
      message: '账号或密码不正确!'
    };
  }
};
</script>
<style lang="scss">
@import 'css/passport/register.scss';
</style>
