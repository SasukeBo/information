<template>
  <div class="charge-new">
    <div class="charge-new__title">
      <i class="el-icon-back go-back-btn" @click="$router.go(-1)"></i>
      <span class="title">增加设备负责人</span>
    </div>

    <div class="charge-new__body global-card">
      <el-form :model="form" label-position="left" label-width="100px" ref="chargeForm">
        <el-form-item label="负责人" prop="name">
          <el-select
            v-model="form.userUUID"
            remote
            filterable
            placeholder="搜索用户"
            :remote-method="querySearchUsers"
          >
            <el-option
              v-for="user in users"
              :key="user.uuid"
              :label="user.userExtend.name"
              :value="user.uuid"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="负责人权限" prop="privIDs">
          <el-transfer
            :titles="['可选权限', '已有权限']"
            v-model="form.privIDs"
            :data="devicePrivs"
            :props="{key: 'id', label: 'name'}"
          ></el-transfer>
        </el-form-item>

        <el-form-item>
          <el-button size="small" type="primary" class="submit" @click="submit()">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import { Autocomplete, Transfer } from 'element-ui';

import usersQuery from './gql/query.users.gql';
import chargesQuery from '../device/charge/gql/query.charges.gql';
import devicePrivsQuery from './gql/query.device-privs.gql';
import chargeCreate from './gql/mutation.charge-create.gql';

export default {
  name: 'charge-new',
  props: ['uuid'],
  components: {
    ElAutocomplete: Autocomplete,
    ElTransfer: Transfer
  },
  apollo: {
    users: {
      query: usersQuery,
      variables() {
        return { namePattern: this.namePattern };
      }
    },
    devicePrivs: {
      query: devicePrivsQuery,
      variables: { privType: 'device' }
    }
  },
  data() {
    return {
      form: {
        deviceUUID: this.uuid,
        userUUID: '',
        privIDs: []
      },
      namePattern: '',
      devicePrivs: [],
      users: []
    };
  },
  methods: {
    querySearchUsers(queryString) {
      this.namePattern = queryString;
    },
    submit() {
      this.$apollo
        .mutate({
          mutation: chargeCreate,
          variables: this.form,
          update: (store, { data: { deviceCharge } }) => {
            try {
              var data = store.readQuery({
                query: chargesQuery,
                variables: { uuid: this.uuid }
              });
              data.charges.unshift(deviceCharge);
              store.writeQuery({
                query: chargesQuery,
                variables: { uuid: this.uuid },
                data
              });
            } catch (e) {
              console.log(e.name);
            }
          }
        })
        .then(({ data }) => {
          this.$message({ type: 'success', message: '操作成功！' });
          this.$router.push({
            name: 'device-charges',
            params: { uuid: this.uuid }
          });
        })
        .catch(e => {
          this.$message({ type: 'error', message: e.message });
        });
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/charge/_new.scss';
</style>
