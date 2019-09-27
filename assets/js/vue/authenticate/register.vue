<template>
  <div class="passport-form register-form">
    <div class="form-title form-item">注册账号</div>
    <transition name="expand">
      <div class="form-alert form-item" v-if="message">{{ message }}</div>
    </transition>
    <div class="form-body">
      <new-and-reset-form ref="registerForm" @submit="submit"></new-and-reset-form>
      <div class="link form-item">
        已有账号，
        <a href="/login" @click.prevent="$router.push({path: 'login'})">直接登录</a>
      </div>
    </div>
  </div>
</template>
<script>
import NewAndResetForm from './new-and-reset-form';
import register from './gql/mutation.register.gql';
import { parseGQLError } from 'js/utils';

export default {
  name: 'register',
  components: { NewAndResetForm },
  data() {
    return {
      message: ''
    };
  },
  methods: {
    submit() {
      this.$apollo
        .mutate({
          mutation: register,
          variables: this.$refs.registerForm.form
        })
        .then(({ data: { register: r } }) => {
          this.$message({
            type: 'success',
            message: '恭喜您，注册成功，请登录。'
          });
          this.$router.push({ name: 'login' });
        })
        .catch(e => {
          app.message = parseGQLError(e).message;
        });
    }
  }
};
</script>
<style lang="scss">
@import 'css/authenticate/register.scss';
</style>
