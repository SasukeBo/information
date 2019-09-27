<template>
  <div class="passport-form forget-form">
    <div class="form-title form-item">找回密码</div>
    <transition name="expand">
      <div class="form-alert form-item" v-if="message">{{ message }}</div>
    </transition>
    <div class="form-body">
      <new-and-reset-form ref="resetForm" @submit="submit"></new-and-reset-form>
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
import NewAndResetForm from './new-and-reset-form';
import resetPassword from './gql/mutation.resetPassword.gql';
import { parseGQLError } from 'js/utils';

export default {
  name: 'reset',
  components: { NewAndResetForm },
  data() {
    return { message: '' };
  },
  methods: {
    submit() {
      this.$apollo
        .mutate({
          mutation: resetPassword,
          variables: this.$refs.resetForm.form
        })
        .then(({ data: { resetPassword: r } }) => {
          this.$message({
            type: 'success',
            message: '重置密码成功，请登录'
          });
          this.$router.push({ name: 'login' });
        })
        .catch(e => {
          this.message = parseGQLError(e).message;
        });
    }
  }
};
</script>
<style lang="scss">
@import 'css/authenticate/register.scss';
</style>
