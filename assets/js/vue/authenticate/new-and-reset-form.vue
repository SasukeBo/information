<template>
  <el-form :model="form" :rules="rules" ref="form">
    <el-form-item prop="phone">
      <el-input
        placeholder="填写手机号"
        ref="phone"
        @keyup.native.enter="beforeSubmit"
        v-model="form.phone"
        prefix-icon="iconfont icon-shouji"
      ></el-input>
    </el-form-item>
    <el-form-item prop="smsCode">
      <div class="securitycode">
        <el-input
          class="securitycode-input"
          placeholder="验证码"
          @keyup.native.enter="beforeSubmit"
          v-model="form.smsCode"
          prefix-icon="iconfont icon-securityCode-b"
        ></el-input>
        <el-button
          round
          type="success"
          class="securitycode-btn"
          @click="beforeSendSmsCode"
          :disabled="waitForNextSend !== 0"
        >
          获取验证码
          <span v-if="waitForNextSend">({{waitForNextSend}}s)</span>
        </el-button>
      </div>
    </el-form-item>
    <el-form-item prop="password">
      <el-input
        :placeholder="placeholder"
        v-model="form.password"
        @keyup.native.enter="beforeSubmit"
        type="password"
        show-password
        prefix-icon="iconfont icon-mima"
      ></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" size="large" @click="beforeSubmit">{{ submitName }}</el-button>
    </el-form-item>
  </el-form>
</template>
<script>
export default {
  name: 'newAndResetForm',
  props: ['waitForNextSend'],
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
    var validateCode = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入短信验证码!'));
      } else {
        callback();
      }
    };
    var validatePasswd = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码!'));
      } else if (value.length < 6) {
        callback(new Error('密码长度请不要低于6位'));
      } else {
        callback();
      }
    };

    return {
      placeholder: '',
      submitName: '',
      form: {
        phone: '',
        smsCode: '',
        password: ''
      },
      rules: {
        phone: [{ validator: validatePhone, trigger: 'blur' }],
        smsCode: [{ validator: validateCode, trigger: 'blur' }],
        password: [{ validator: validatePasswd, trigger: 'blur' }]
      }
    };
  },
  watch: {
    $route: {
      immediate: true,
      handler: function(newVal) {
        if (this.$route.name === 'reset_password') {
          this.placeholder = '设置新密码';
          this.submitName = '提交';
        } else if (this.$route.name === 'register') {
          this.placeholder = '设置密码';
          this.submitName = '注册';
        }
      }
    }
  },
  methods: {
    beforeSendSmsCode() {
      this.$refs['form'].validateField('phone');
      if (this.$refs['phone'].validateState !== 'error') {
        this.$emit('sendSmsCode');
      }
    },

    beforeSubmit() {
      this.$refs['form'].validate(valid => {
        if (valid) {
          this.$emit('submit');
        } else {
          console.log('submit failed');
        }
      });
    }
  }
};
</script>
