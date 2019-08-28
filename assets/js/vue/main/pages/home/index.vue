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
        <device-card></device-card>
        <login-log-card></login-log-card>
      </div>

      <div class="home-right">
        <notice-card></notice-card>
      </div>
    </div>
  </div>
</template>
<script>
import { getWeekDay } from 'js/utils';
import { mapState } from 'vuex';
import LoginLogCard from './_login-logs-card';
import DeviceCard from './_device-card';
import NoticeCard from './_notice-card';

export default {
  name: 'page-home',
  components: {
    LoginLogCard,
    DeviceCard,
    NoticeCard
  },
  computed: {
    date() {
      var now = new Date();
      return `${now.getUTCFullYear()}年${now.getMonth() +
        1}月${now.getDate()}日${getWeekDay(now)}`;
    },
    ...mapState({
      profile: state => state.user.userExtend
    })
  }
};
</script>
<style lang="scss">
@import 'css/main/home/index.scss';
</style>
