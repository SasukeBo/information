<template>
  <div class="login-form passport-form">
    <div class="form-title form-item">欢迎登录</div>
    <transition name="expand">
      <div class="form-alert form-item" v-if="message">{{ message }}</div>
    </transition>
    <div class="form-body">
      <el-form :model="loginForm" :rules="rules" ref="loginForm">
        <el-form-item prop="phone">
          <el-input
            placeholder="手机号"
            @keyup.native.enter="beforeSubmit"
            v-model="loginForm.phone"
            prefix-icon="iconfont icon-shouji"
          ></el-input>
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            placeholder="密码"
            @keyup.native.enter="beforeSubmit"
            v-model="loginForm.password"
            show-password
            prefix-icon="iconfont icon-mima"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <div class="login-options">
            <el-checkbox v-model="loginForm.remember" class="remember-login">记住登录</el-checkbox>
            <div class="forget-password">
              <a
                href="/reset_password"
                @click.prevent="$router.push({path: '/reset_password'})"
              >忘记密码？</a>
            </div>
          </div>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" size="large" @click="beforeSubmit" class="passport-form__btn">登录</el-button>
        </el-form-item>
      </el-form>
      <div class="link form-item">
        <a href="/register" @click.prevent="$router.push({path: '/register'})">注册账号</a>
      </div>
    </div>
  </div>
</template>
<script>
import signIn from './gql/mutation.signIn.gql';
import { parseGQLError } from 'js/utils';

export default {
  name: 'login',

  data() {
    var reg = new RegExp(
      '^(?:\\+?86)?1(?:3\\d{3}|5[^4\\D]\\d{2}|8\\d{3}|7(?:[35678]\\d{2}|4(?:0\\d|1[0-2]|9\\d))|9[189]\\d{2}|66\\d{2})\\d{6}$'
    );
    var validatePhone = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入手机号!'));
      } else if (!reg.test(value)) {
        callback(new Error('手机号不合法!'));
      } else {
        callback();
      }
    };
    var validatePasswd = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码!'));
      } else {
        callback();
      }
    };
    return {
      rules: {
        phone: [{ validator: validatePhone, trigger: 'blur' }],
        password: [{ validator: validatePasswd, trigger: 'blur' }]
      },
      loginForm: {
        phone: '',
        password: '',
        remember: true
      },
      message: ''
    };
  },
  methods: {
    beforeSubmit() {
      this.$refs['loginForm'].validate(valid => {
        if (valid) {
          this.$apollo
            .mutate({
              mutation: signIn,
              variables: this.loginForm
            })
            .then(() => {
              this.$router.push({ name: 'home' });
            })
            .catch(e => {
              this.message = parseGQLError(e).message;
            });
        }
      });
    }
  }
};
</script>
<style lang="scss">
@import 'css/authenticate/login.scss';
</style>
