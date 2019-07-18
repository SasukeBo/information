<template>
  <div class="main-login">
    <div class="title">
      <div class="title-main">设备资讯化</div>
      <div class="title-sub">普创智控</div>
    </div>

    <div class="login-form">
      <div class="form-title form-item">欢迎登录</div>
      <div class="form-alert form-item">{{ message }}</div>
      <div class="form-body">
        <el-input
          class="form-item"
          placeholder="手机号"
          size="large"
          v-model="phone"
          prefix-icon="iconfont icon-shouji"
        ></el-input>
        <el-input
          class="form-item"
          placeholder="密码"
          v-model="password"
          prefix-icon="iconfont icon-mima"
        ></el-input>
        <div class="form-item forget-password">
          <a href="/">忘记密码？</a>
        </div>
        <el-button class="form-item" type="primary" size="large">登录</el-button>
        <div class="to-register form-item">
          <a href="/">注册账号</a>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      phone: '',
      password: '',
      message: '账号或密码不正确!'
    };
  }
};
</script>
<style lang="scss">
@import 'css/passport/login.scss';
</style>
