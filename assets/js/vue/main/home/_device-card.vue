<template>
  <div class="global-card device-card">
    <span class="global-card__title">设备概览</span>

    <div class="data-item device-groups">
      <div class="device-group">
        <div class="data-label device-group__title">所有设备</div>

        <div style="display: flex">
          <div class="device-item">
            <i class="iconfont icon-shebei"></i>
            <span>共 {{ own.total }} 台</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-on-line"></i>
            <span>在线 {{ own.online }} 台</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-off-line"></i>
            <span>离线 {{ own.offline }} 台</span>
          </div>
        </div>
      </div>

      <div class="device-group">
        <div class="data-label device-group__title">我创建的</div>

        <div style="display: flex">
          <div class="device-item">
            <i class="iconfont icon-shebei"></i>
            <span>共 {{ charge.total }} 台</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-on-line"></i>
            <span>在线 {{ charge.online }} 台</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-off-line"></i>
            <span>离线 {{ charge.offline }} 台</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: 'device-card',
  data() {
    return {
      own: {
        total: 10,
        online: 8,
        offline: 2
      },
      charge: {
        total: 23,
        online: 23,
        offline: 0
      }
    };
  }
};
</script>
