<template>
  <div class="global-card notice-card">
    <span class="global-card__title">系统通告</span>

    <div class="global-card__body">
      <div v-for="(tip, index) in tips" :key="'tip_'+index" class="notice-item">
        <a href="javascript:;" class="notice-item__message">{{ tip.message }}</a>
        <div class="notice-item__time">{{ tip.time }}</div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'notice-card',
  data() {
    return {
      tips: [
        {
          message:
            '通知1内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容',
          time: '2019/8/19 下午1:41:22'
        },
        {
          message:
            '通知1内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容',
          time: '2019/8/18 下午1:41:22'
        },
        {
          message:
            '通知1内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容内容',
          time: '2019/8/18 下午1:41:22'
        }
      ]
    };
  }
};
</script>
