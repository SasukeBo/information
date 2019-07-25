<template>
  <div class="passport-form register-form">
    <div class="form-title form-item">注册账号</div>
    <div class="form-alert form-item" v-if="message">{{ message }}</div>
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
        <el-button type="success" round class="securitycode-btn" @click="showCaptcha = true" :disabled="phone === ''">获取验证码</el-button>
      </div>
      <el-input
        class="form-item"
        placeholder="密码"
        v-model="password"
        type="password"
        show-password
        prefix-icon="iconfont icon-mima"
      ></el-input>
      <el-button class="form-item" type="primary" size="large" @click="message = '账号或密码不正确!'" :disabled="!canSubmitRegister">注册</el-button>
      <div class="link form-item">
        已有账号，
        <a href="/login" @click.prevent="$router.push({path: 'login'})">直接登录</a>
      </div>
    </div>

    <in-slide-captcha
      :showCaptcha.sync="showCaptcha"
      v-if="showCaptcha"
      @verified="sendSmsCode()"
    ></in-slide-captcha>
  </div>
</template>
<script>
import InSlideCaptcha from '../slide-captcha';
import gql from 'graphql-tag';
import apollo from './apollo'

export default {
  name: 'register',
  components: {
    InSlideCaptcha
  },
  ...apollo,
  data() {
    return {
      showCaptcha: false,
      phone: '',
      securityCode: '',
      password: '',
      message: ''
    };
  },
  computed: {
    canSubmitRegister() {
      if (this.phone && this.securityCode && this.password) return true
      return false
    }
  },
  methods: {
    sendSmsCode() {
      var _this = this
      _this.$apollo.mutate({
        mutation: gql`
        mutation sendSmsCode ($phone: String!) {
          sendSmsCode(phone: $phone) {
            message
            code
          }
        }
        `,
        variables: {
          phone: _this.phone
        }
      }).then(({data: {sendSmsCode: res}}) => {
        if (res.message === 'OK') _this.message = "send success"
        else console.log(res.code)
      }).catch(e => {
        console.log(e)
      })
    }
  }
};
</script>
<style lang="scss">
@import 'css/passport/register.scss';
</style>
