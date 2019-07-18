<template>
  <div class="login-form passport-form">
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
        show-password
        prefix-icon="iconfont icon-mima"
      ></el-input>
      <div class="form-item forget-password">
        <a href="/forget_password" @click.prevent="$router.push({path: '/forget_password'})">忘记密码？</a>
      </div>
      <el-button class="form-item" type="primary" size="large">登录</el-button>
      <div class="link form-item">
        <a href="/register" @click.prevent="$router.push({path: '/register'})">注册账号</a>
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
