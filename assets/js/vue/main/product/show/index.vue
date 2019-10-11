<template>
  <div class="product-show"></div>
</template>
<script>
export default {
  name: 'product-show',
  props: ['id'],
  mounted() {
    NProgress.done()
  }
}
</script>
