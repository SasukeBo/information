<template>
  <div class="device-details">
    <div class="details__left">
      <device-basic :uuid="uuid"></device-basic>
      <device-register :uuid="uuid"></device-register>
    </div>

    <div class="details__right">
      <param-values :uuid="uuid" v-if="uuid"></param-values>
    </div>
  </div>
</template>
<script>
import deviceBasic from './_device-basic';
import deviceRegister from './_device-register';
import paramValues from './_values';

export default {
  name: 'device-details',
  props: ['uuid'],
  components: { deviceBasic, deviceRegister, paramValues }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-details {
  display: flex;

  .details__left {
    flex: 1;
  }

  .details__right {
    flex: 3;
    padding: 1rem;
  }
}

.device-details .details__basic-info {
  padding: 1rem;

  .title {
    padding: 0.4rem 0.8rem;
    border-left: 4px solid;
    font-size: 1rem;
    margin-bottom: 0.6rem;
    background: lighten($--color-theme__bg, 10%);

    i {
      font-size: 1.5rem;
      padding-right: 0.5rem;
      vertical-align: middle;
    }
  }

  .item {
    padding: 0 1rem;
    margin-bottom: 0.3rem;
    font-size: 0.85rem;
    color: $--color-font__gray;
    display: flex;

    span:first-child {
      color: $--color-font__light;
      min-width: 100px;
    }
  }
}

</style>
