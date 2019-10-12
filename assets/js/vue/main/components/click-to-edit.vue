<template>
  <div class="click-to-edit">
    <span v-if="edit">
      <slot name="form"></slot>
    </span>
    <span v-else @click="edit = true">
      <slot name="text"></slot>
    </span>

    <i class="el-icon-check" v-if="edit" @click="save"></i>
    <i class="el-icon-close" v-if="edit" @click="edit = false"></i>
  </div>
</template>
<script>
export default {
  name: 'click-to-edit',
  data() {
    return {
      edit: false
    };
  },
  methods: {
    save() {
      this.$emit('save');
      this.edit = false;
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';
.click-to-edit {
  cursor: pointer;
  color: $--color-theme__main;

  .el-icon-check {
    color: $--color-theme__main;
  }

  .el-icon-close {
    color: $--color-theme__danger;
  }

  i {
    font-size: 1.5rem;
    font-weight: bold;
    padding: 0 0rem;
    vertical-align: middle;
  }
}
</style>
