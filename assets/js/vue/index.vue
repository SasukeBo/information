<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>
<script>
export default {
  name: 'app',
  created() {
    this.$socket.connect();
  },
  beforeDestroy() {
    this.$socket.close();
    alert('destroy');
  }
};
</script>

<style lang="scss">
</style>
