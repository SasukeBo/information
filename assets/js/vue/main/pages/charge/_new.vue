<template>
  <div class="charge-new">
    <div class="charge-new__title">
      <i class="el-icon-back go-back-btn" @click="$router.go(-1)"></i>
      <span class="title">增加设备负责人</span>
    </div>

    <div class="charge-new__body global-card">
      <el-form :model="form" label-position="left" label-width="100px" ref="chargeForm">
        <el-form-item label="负责人" prop="name">
          <el-autocomplete
            v-model="form.name"
            :fetch-suggestions="querySearchUsers"
            value-key="name"
            placeholder="选择指派人"
            @select="handleSelect"
          ></el-autocomplete>
        </el-form-item>

        <el-form-item label="负责人权限" prop="privIDs">
          <el-transfer :titles="['可选权限', '已有权限']" v-model="form.privIDs" :data="devicePrivs"></el-transfer>
        </el-form-item>

        <el-form-item>
          <el-button size="small" type="primary" class="submit" @click="saveCharge()">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import { apollo } from '../device-new/graphql';
import { newApollo } from './graphql';
import { Autocomplete, Transfer } from 'element-ui';

export default {
  name: 'charge-new',
  props: ['uuid'],
  components: {
    ElAutocomplete: Autocomplete,
    ElTransfer: Transfer
  },
  apollo: { ...apollo, ...newApollo },
  data() {
    return {
      form: {
        name: '',
        userUUID: '',
        privIDs: []
      },
      devicePrivs: [],
      users: []
    };
  },
  methods: {
    querySearchUsers(queryString, cb) {
      var users = this.users.map(user => {
        return { name: user.userExtend.name, uuid: user.uuid };
      });
      var results = queryString
        ? users.filter(this.createUserFilter(queryString))
        : users;
      cb(results);
    },
    createUserFilter(queryString) {
      return user => {
        return user.name.toLowerCase().indexOf(queryString.toLowerCase()) === 0;
      };
    },
    handleSelect(user) {
      this.form.userUUID = user.uuid;
    },
    saveCharge() {
      console.log('saveCharge');
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/charge/_new.scss';
</style>
