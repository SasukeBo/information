<template>
  <div class="user-device">
    <div class="page-left toggle-options">
      <div class="left-title">
        <span>筛选</span>
        <i class="el-icon-caret-bottom" @click="isExpand = !isExpand"></i>
      </div>

      <transition name="expand">
        <div v-show="isExpand" class="expand-panel">
          <div class="global-card status-btns">
            <div class="status-btn">
              <div class="status-btn__icon">
                <i class="iconfont icon-production animation"></i>
              </div>

              <div class="status-btn__label">在生产</div>
            </div>

            <div class="status-btn">
              <div class="status-btn__icon">
                <i class="iconfont icon-production"></i>
              </div>

              <div class="status-btn__label">已停机</div>
            </div>

            <div class="status-btn">
              <div class="status-btn__icon">
                <i class="iconfont icon-off-line"></i>
              </div>

              <div class="status-btn__label">已离线</div>
            </div>
          </div>

          <div class="global-card name-search">
            <el-input placeholder="搜索设备" prefix-icon="el-icon-search" v-model="search"></el-input>
          </div>

          <div class="global-card ship-btns">
            <div class="ship-btns__label">关系类型：</div>
            <el-checkbox-group v-model="checkboxGroup">
              <el-checkbox-button label="owner">创建的</el-checkbox-button>
              <el-checkbox-button label="charger">负责的</el-checkbox-button>
            </el-checkbox-group>
          </div>
        </div>
      </transition>
    </div>

    <div class="page-right">
      <div></div>
      <div class="right-title">设备列表</div>

      <div class="device-card-list">
        <device-card v-for="i in 10" :device="device" :key="'device-card_' + i"></device-card>
      </div>
    </div>
  </div>
</template>
<script>
import DeviceCard from './_device-card';

export default {
  name: 'devices',
  components: { DeviceCard },
  data() {
    return {
      search: '',
      checkboxGroup: [],
      isExpand: false,
      device: {
        name: '自动化设备',
        type: '载带包装',
        mac: '2C:6E:85:2B:CB:18',
        token: 'a60d6ceda891bd',
        uuid: '205b75f3-2dfe-46a1-8bd9-a4edb16c1f60',
        user: {
          uuid: '52b4396d-0528-4446-989c-0280253a7a82',
          userExtend: { name: 'sasuke' }
        },
        status: 'producting',
        description: '自动化载带包装',
        createdAt: '2019-08-14 09:31:50.915787+08',
        updatedAt: '2019-08-14 09:31:50.915789+08'
      }
    };
  }
};
</script>
<style lang="scss">
@import 'css/main/devices/index.scss';
</style>
