<template>
  <div class="global-card device-card">
    <span class="global-card__title">设备概览</span>

    <div class="data-item device-groups">
      <div class="device-group">
        <div class="data-label device-group__title">所有设备</div>

        <div style="display: flex">
          <div class="device-item">
            <i class="iconfont icon-running"></i>
            <span>{{ total && total.prod ? total.prod : '-' }} 台生产</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-stopping"></i>
            <span>{{ total && total.stop ? total.stop : '-' }} 台停机</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-off-line"></i>
            <span>{{ total && total.offline ? total.offline : '-' }} 台离线</span>
          </div>
        </div>
      </div>

      <div class="device-group">
        <div class="data-label device-group__title">我创建的</div>

        <div style="display: flex">
          <div class="device-item">
            <i class="iconfont icon-running"></i>
            <span>{{ register && register.prod ? register.prod : '-' }} 台生产</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-stopping"></i>
            <span>{{ register && register.stop ? register.stop : '-' }} 台停机</span>
          </div>

          <div class="device-item">
            <i class="iconfont icon-off-line"></i>
            <span>{{ register && register.offline ? register.offline : '-' }} 台离线</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import countQuery from './gql/query.device-status-count.gql';

export default {
  name: 'device-card',
  apollo: {
    total: {
      query: countQuery,
      update(data) {
        return data.count;
      }
    },
    register: {
      query: countQuery,
      variables: { filter: 'register' },
      update(data) {
        return data.count;
      }
    }
  },
  data() {
    return {
      total: undefined,
      register: undefined
    };
  }
};
</script>
