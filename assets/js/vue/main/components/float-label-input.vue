<template>
  <div class="floating-label">
    <span :style="{display: value ? 'inline-block' : 'none'}">{{ placeholder || ''}}</span>
    <input
      class="global-input"
      type="text"
      :value="value"
      @input="$emit('input', $event.target.value)"
      :placeholder="placeholder"
      :style="{paddingTop: value ? '14px': '0'}"
    />
  </div>
</template>
<script>
export default {
  model: {
    prop: 'value',
    event: 'input'
  },
  props: ['value', 'placeholder'],
  name: 'float-label-input'
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.floating-label {
  display: inline-block;
  width: 100%;
  position: relative;
  margin-bottom: 15px;

  span {
    position: absolute;
    top: 5px;
    left: 16px;
    pointer-events: none;
    font-size: 11px;
    color: #9da2a6;
    font-weight: bold;
    line-height: 18px;
    display: none;
    z-index: 100;
  }

  .global-input {
    width: 100%;
    border: 1px solid $--color-border__1;
    border-radius: 3px;
    vertical-align: middle;
    height: 50px;
    padding: 0px 15px;
    color: $--color-font__dark;
    font-size: 15px;
    background-color: $--color-theme__white;
    font-weight: inherit;
    transition: all 0.1s ease-in-out 0s;

    &:focus {
      border-color: $--color-theme__main;
    }
  }
}
</style>
