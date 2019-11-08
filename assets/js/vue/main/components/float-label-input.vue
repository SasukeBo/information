<template>
  <div class="floating-label" @click="clickable && $refs.datepicker.handleClickIcon()">
    <span :style="{display: value ? 'inline-block' : 'none'}">{{ placeholder || ''}}</span>
    <input
      class="global-input"
      :class="{disabled: !editable}"
      type="text"
      @keyup.enter="enter"
      :value="inputValue"
      :disabled="!editable || type === 'date'"
      @input="editable && $emit('input', $event.target.value)"
      :placeholder="placeholder"
      :style="{paddingTop: value ? '14px': '0', cursor}"
    />

    <el-date-picker ref="datepicker" v-if="type === 'date' " v-model="dateValue" type="date"></el-date-picker>
  </div>
</template>
<script>
import { timeFormatter } from 'js/utils';

export default {
  name: 'float-label-input',
  model: {
    prop: 'value',
    event: 'input'
  },
  props: {
    value: [String, Number, Boolean],
    type: {
      type: String,
      default: 'text'
    },
    placeholder: String,
    enter: {
      type: Function,
      default: () => undefined
    },
    editable: {
      type: Boolean,
      default: true
    }
  },
  computed: {
    cursor() {
      if (this.type === 'date') {
        return 'pointer';
      } else if (!this.editable) {
        return 'not-allowed';
      }
    },
    clickable() {
      return this.type === 'date' ? true : false;
    }
  },
  data() {
    return {
      inputValue: '',
      dateValue: ''
    };
  },
  watch: {
    value: {
      immediate: true,
      handler: function(newVal) {
        if (this.type === 'date') {
          if (newVal) this.inputValue = timeFormatter(newVal, '%y年%m月%d日');
          this.dateValue = newVal;
        } else {
          this.inputValue = newVal;
        }
      }
    },
    dateValue: {
      immediate: true,
      handler: function(newVal) {
        if (newVal) this.inputValue = timeFormatter(newVal, '%y年%m月%d日');
        if (newVal && newVal !== this.value) {
          var date = new Date(newVal);
          this.$emit('input', date.toISOString());
        }
      }
    }
  }
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

    &.disabled {
      background-color: #f5f7fa;
      border-color: #e4e7ed;
      color: #c0c4cc;
    }
  }

  .el-date-editor {
    position: absolute;
    left: 0;
    z-index: -1;
  }
}
</style>
