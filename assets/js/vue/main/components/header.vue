<template>
  <div class="com-header">
    <div class="com-header__left">
      <div class="menu-button">
        <i
          :class="['el-icon-s-unfold', menuOpen ? 'transform-rotate' : '']"
          @click="$emit('toggle-menu')"
        ></i>
      </div>
      <div class="title">
        <div>普创智控</div>
        <div>protron</div>
      </div>
    </div>
    <div class="com-header__center">
      <el-input placeholder="搜索" style="display: none;" prefix-icon="el-icon-search"></el-input>
    </div>

    <div class="com-header__right">
      <div class="topbar-entry">
        <img class="topbar-entry__avatar" :src="avatarURL ? avatarURL : '~images/avatar.jpg'" />

        <div class="topbar-body global-card">
          <div class="topbar-body__header">
            <img :src="avatarURL ? avatarURL : '~images/avatar.jpg'" class="avatar" />
            <span style="font-size: 1.5rem">{{ name ? name : phone }}</span>
            <el-tag
              :type="role.isAdmin ? 'warning' : 'success'"
              size="mini"
              v-if="role"
            >{{ role.roleName }}</el-tag>
          </div>

          <div class="topbar-body__item">
            <i class="el-icon-mobile-phone"></i>
            <span>{{ phone }}</span>
            <span class="hover-show">点击修改</span>
          </div>

          <div class="topbar-body__item" v-if="email">
            <i class="iconfont icon-185078emailmailstreamline"></i>
            <span>{{ email }}</span>
            <span class="hover-show">点击修改</span>
          </div>

          <div class="topbar-body__item">
            <i class="el-icon-setting"></i>
            <span>设置中心</span>
          </div>

          <div class="topbar-body__footer" @click="logout">
            <span>退出登录</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { mapState } from 'vuex';
import gql from 'graphql-tag';

export default {
  name: 'com-header',
  props: ['menuOpen'],
  computed: {
    ...mapState({
      avatarURL: state => state.user.avatarURL,
      phone: state => state.user.phone,
      role: state => state.user.role,
      name: state => state.user.name,
      email: state => state.user.email
    })
  },
  methods: {
    logout() {
      this.$apollo
        .mutate({
          mutation: gql`
            mutation signOut {
              signOut
            }
          `
        })
        .then(() => {
          this.$store.dispatch('user/logout');
          this.$router.push({
            name: 'login',
            query: { return_to: this.$route.name, ...this.$route.params }
          });
        });
    }
  }
};
</script>
