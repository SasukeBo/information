<template>
  <div class="passport-form forget-form">
    <div class="form-title form-item">找回密码</div>
    <transition name="expand">
      <div class="form-alert form-item" v-if="message">{{ message }}</div>
    </transition>
    <div class="form-body">
      <new-and-reset-form
        ref="form"
        :waitForNextSend="waitForNextSend"
        @sendSmsCode="showCaptcha = true"
        @submit="submit"
      ></new-and-reset-form>
      <div class="link form-item">
        <a href="/login" @click.prevent="$router.push({path: 'login'})">
          <i class="iconfont icon-fanhui"></i>
          <span>返回登录</span>
        </a>
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
import InSlideCaptcha from './slide-captcha';
import NewAndResetForm from './new-and-reset-form';
import gql from './graphql';

export default {
  name: 'reset',
  components: {
    InSlideCaptcha,
    NewAndResetForm
  },

  data() {
    return {
      showCaptcha: false,
      message: '',
      waitForNextSend: 0
    };
  },
  methods: {
    submit() {
      gql.resetPassword(this);
    },
    sendSmsCode() {
      // 设置等待60s
      this.waitForNextSend = 60;
      var interval = setInterval(() => {
        this.waitForNextSend--;
        if (!this.waitForNextSend) clearInterval(interval);
      }, 1000);
      gql.sendSmsCode(this);
    }
  }
};
</script>
<style lang="scss">
@import 'css/authenticate/register.scss';
</style>
