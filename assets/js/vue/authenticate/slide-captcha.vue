<template>
  <div class="captcha-dialog" ref="dialog">
    <div class="modal" @click="closeDialog()"></div>
    <div ref="captcha" class="captcha">
      <i class="close-dialog el-icon-close" @click="closeDialog()"></i>
      <div class="slide-tip">手机浏览器请关闭滑屏前进后退功能完成验证</div>
    </div>
  </div>
</template>
<script>
import 'js/utils/jigsaw.min.js';

export default {
  props: ['showCaptcha'],
  data() {
    return {};
  },
  mounted() {
    var _this = this;
    var captcha = _this.$refs.captcha;

    document.addEventListener(
      'touchmove',
      e => {
        e.preventDefault();
      },
      false
    );

    window.jigsaw.init({
      el: captcha,
      onSuccess: function() {
        setTimeout(() => {
          _this.closeDialog();
          _this.$emit('verified');
        }, 1000);
      },
      url: '/images',
      limit: 200
    });
    setTimeout(() => (_this.$refs.dialog.style.opacity = 1), 100);
  },
  methods: {
    closeDialog() {
      this.$refs.dialog.style.opacity = 0;
      setTimeout(
        () => this.$emit('update:showCaptcha', !this.showCaptcha),
        300
      );
    }
  }
};
</script>
<style lang="scss">
.captcha-dialog {
  opacity: 0;
  position: absolute;
  width: 100%;
  height: 100%;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  transition: all 0.3s ease;
}

.captcha-dialog .modal {
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.7);
  border-radius: 4px;
}

.captcha-dialog .captcha {
  position: absolute !important;
  top: 0;
  right: 0;
  left: 0;
  bottom: 0;
  height: 214px;
  background: #fff !important;
  margin: auto !important;
}

.captcha-dialog .captcha .slide-tip {
  position: absolute;
  bottom: -29px;
  margin: auto;
  left: 0;
  right: 0;
  text-align: center;
  color: #afafaf;
  display: none;
}

.captcha-dialog .close-dialog {
  position: absolute;
  width: 20px;
  height: 20px;
  right: -14px;
  top: -14px;
  background: #fff;
  text-align: center;
  line-height: 20px;
  border-radius: 50%;
  cursor: pointer;
}

@media only screen and (max-width: 440px) {
  .captcha-dialog .captcha .slide-tip {
    display: block;
  }
}
</style>
