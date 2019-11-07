<template>
  <div class="admin">
    <p>网站后台页面需要全新设计</p>
    <router-view></router-view>
  </div>
</template>
<script>
export default {
  name: 'admin'
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.admin {
  width: 1000px;
  margin: auto;
}
</style>
