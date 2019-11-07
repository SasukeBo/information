<template>
  <div class="product-setting">
    <div class="side-nav">
      <div>
        <div
          class="side-nav__item"
          @click="updateHash('#name')"
          :class="{active: currentHash === '#name'}"
        >名称</div>
      </div>

      <div>
        <div
          class="side-nav__item"
          @click="updateHash('#charger')"
          :class="{active: currentHash === '#charger'}"
        >生产负责人</div>
      </div>

      <div>
        <div
          class="side-nav__item"
          @click="updateHash('#customer')"
          :class="{active: currentHash === '#customer'}"
        >订货方</div>
      </div>

      <div>
        <div
          class="side-nav__item"
          @click="updateHash('#order')"
          :class="{active: currentHash === '#order'}"
        >订单信息</div>
      </div>

      <div>
        <div
          class="side-nav__item"
          @click="updateHash('#detect-items')"
          :class="{active: currentHash === '#detect-items'}"
        >检测项</div>
      </div>
    </div>

    <div class="setting-content">
      <component :is="currentComponent" :id="id"></component>
    </div>
  </div>
</template>
<script>
// components
import settingname from './_setting-name';
import settingorder from './_setting-order';
import settingcharger from './_setting-charger';
import settingdetect from './_setting-detect';
import settingcustomer from './_setting-customer';

export default {
  props: ['id'],
  name: 'setting',
  components: {
    settingname,
    settingorder,
    settingcharger,
    settingdetect,
    settingcustomer
  },
  computed: {
    currentHash() {
      return this.$route.hash ? this.$route.hash : '#name';
    },
    currentComponent() {
      switch (this.currentHash) {
        case '#name':
          return 'settingname';
        case '#charger':
          return 'settingcharger';
        case '#order':
          return 'settingorder';
        case '#detect-items':
          return 'settingdetect';
        case '#customer':
          return 'settingcustomer';
      }
    }
  },
  methods: {
    updateHash(hash) {
      if (this.currentHash !== hash)
        this.$router.push({
          name: 'product-show',
          query: { tab: 'setting' },
          hash
        });
    }
  },
  mounted() {
    NProgress.done();
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.product-setting {
  display: flex;
  justify-content: space-between;

  .side-nav {
    flex: 0 0 20%;
  }

  .side-nav__item {
    margin-bottom: 15px;
    display: inline-block;
    color: $--color-font__silver;
    cursor: pointer;
    line-height: 20px;
    padding: 6px 0;
    position: relative;
    transition: all 0.2s ease-in;

    &:after {
      display: block;
      position: absolute;
      z-index: 10;
      content: '';
      left: 50%;
      bottom: 0;
      height: 2px;
      width: 0px;
      background-color: $--color-theme__main;
      transition: all 0.2s ease-in;
    }

    &:hover,
    &.active {
      color: $--color-font__white;
      font-weight: bold;
    }

    &.active:after {
      width: 100%;
      margin-left: -50%;
    }
  }

  .setting-content {
    width: 100%;

    .content-block {
      max-width: 520px;
      margin-bottom: 40px;
    }

    .field-title {
      position: relative;
      font-size: 24px;
      line-height: 24px;
      margin: 0px 0px 25px 0px;
      color: $--color-font__white;
      font-weight: 400;

      .el-button {
        position: absolute;
        right: 0;
        width: 80px;
      }
    }

    .el-button.submit {
      width: 100%;
    }
  }
}
</style>
