<template>
  <div class="global-card login-log-card">
    <div class="login-groups">
      <div class="login-item">
        <div class="data-label login-item__title">本次登录信息</div>

        <div>
          <div class="data-item">
            <i class="iconfont icon-ip"></i>
            <div class="table-cell">{{ thisLogin.remoteIP }}</div>
          </div>

          <div class="data-item">
            <i class="iconfont icon-monitor-tablet-and-s"></i>
            <div class="table-cell">{{ thisLoginDevice }}</div>
          </div>
        </div>
      </div>

      <div class="login-item">
        <div class="data-label login-item__title">
          最近一次登录信息
          <span class="login-time">{{ lastLoginTime }}</span>
        </div>

        <div>
          <div class="data-item">
            <i class="iconfont icon-ip"></i>
            <div class="table-cell">{{ lastLogin ? lastLogin.remoteIP : '-' }}</div>
          </div>

          <div class="data-item">
            <i class="iconfont icon-monitor-tablet-and-s"></i>
            <div class="table-cell">{{ lastLoginDevice }}</div>
          </div>
        </div>

        <a href="javascript:;" class="login-logs-more">
          <span style="padding-right: 0.5rem">查看更多</span>
          <i class="el-icon-d-arrow-right"></i>
        </a>
      </div>
    </div>
  </div>
</template>
<script>
import thisLoginQuery from './gql/query.thisLogin.gql';
import lastLoginQuery from './gql/query.lastLogin.gql';
import { parseUserAgent, parseGQLError } from 'js/utils';

export default {
  name: 'login-logs-card',
  data() {
    return {
      lastLogin: {},
      thisLogin: {}
    };
  },
  apollo: {
    thisLogin: { query: thisLoginQuery },
    lastLogin: {
      query: lastLoginQuery,
      error(e) {
        var err = parseGQLError(e);
        console.warn(err.message, ', ', err.originMessage);
      }
    }
  },
  computed: {
    lastLoginTime() {
      if (this.lastLogin && this.lastLogin.createdAt) {
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
      if (this.lastLogin && this.lastLogin.userAgent) {
        return parseUserAgent(this.lastLogin.userAgent);
      }

      return '-';
    }
  }
};
</script>
