<template>
  <div class="page-home">
    <el-row style="padding: 0 0.5rem">
      <el-col style="min-width: 17.5rem">
        <div class="greeting-title">
          <span style="font-size: 0.875rem">{{ profile.name }}您好，</span>
          <span>今天是</span>
        </div>
        <div class="greeting-message">{{ date }}，欢迎回到普创智控自动化设备云平台。</div>
      </el-col>
    </el-row>

    <div class="home-body">
      <div class="home-left">
        <div class="global-card status-card">
          <span class="global-card__title">账号状态</span>
          <div class="data-item">
            <span>{{status}}</span>
            /
            <span class="data-label">身份</span>
            <span>{{ role.roleName }}</span>
            <el-tag v-if="role.isAdmin" type="danger" size="small">管理员</el-tag>
          </div>

          <div class="data-item">
            <div class="data-label" style="padding-bottom: 0.5rem">当前账号绑定联系方式为</div>

            <div class="contact-flex">
              <div class="contact-item">
                <i class="el-icon-mobile-phone"></i>
                <span class="data">{{ phone }}</span>
                <span class="show-edit-tip">点击修改</span>
              </div>

              <div class="contact-item">
                <i class="el-icon-message"></i>
                <span class="data">{{ profile.email ? profile.email : '未绑定邮箱'}}</span>
                <span v-if="!profile.email">绑定邮箱</span>
                <span class="show-edit-tip">点击修改</span>
              </div>
            </div>
          </div>

          <div class="data-item login-flex">
            <div class="login-item">
              <div class="item-title">本次登录信息</div>
              <div>
                <span class="data-label">IP</span>
                <span>{{ thisLogin.remoteIP }}</span>
              </div>
              <div>
                <span class="data-label">登录设备信息</span>
                <span>{{ thisLoginDevice }}</span>
              </div>
            </div>

            <div class="login-item">
              <div class="item-title">最近一次登录信息</div>
              <div>
                <span class="data-label">IP</span>
                <span>{{ lastLogin.remoteIP }}</span>
              </div>
              <div>
                <span class="data-label">登录设备信息</span>
                <span>{{ lastLoginDevice }}</span>
              </div>
              <div>
                <span class="data-label">登录时间</span>
                <span>{{ lastLoginTime }}</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="home-right">
        <div class="global-card">暂无内容展示</div>
      </div>
    </div>
  </div>
</template>
<script>
import { getWeekDay, parseUserAgent } from 'js/utils';
import { mapState } from 'vuex';
import { apollo } from './graphql';

export default {
  name: 'page-home',
  data() {
    return {
      lastLogin: {},
      thisLogin: {}
    };
  },
  apollo,
  computed: {
    lastLoginTime() {
      if (this.lastLogin.createdAt) {
        var time = new Date(this.lastLogin.createdAt);
        return `${time.getUTCFullYear()}年${time.getMonth() +
          1}月${time.getDate()}日${time.toLocaleTimeString()}`;
      }

      return '';
    },
    thisLoginDevice() {
      if (this.thisLogin.userAgent) {
        return parseUserAgent(this.thisLogin.userAgent);
      }

      return '';
    },
    lastLoginDevice() {
      if (this.lastLogin.userAgent) {
        return parseUserAgent(this.lastLogin.userAgent);
      }

      return '';
    },
    date() {
      var now = new Date();
      return `${now.getUTCFullYear()}年${now.getMonth() +
        1}月${now.getDate()}日${getWeekDay(now)}`;
    },
    ...mapState({
      status: state => state.user.status,
      role: state => state.user.role,
      phone: state => state.user.phone,
      profile: state => state.user.profile
    })
  }
};
</script>
<style lang="scss">
@import 'css/main/home.scss';
</style>
