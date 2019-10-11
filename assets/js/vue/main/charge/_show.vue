<template>
  <div class="charge-show" v-if="!$apollo.queries.charge.loading">
    <div class="charge-show__title">
      <i class="el-icon-back go-back-btn" @click="$router.go(-1)"></i>
      <span class="title">设备负责人</span>
    </div>

    <div class="charge-show__body global-card">
      <div class="device-data">
        <i class="el-icon-s-platform"></i>
        <span class="label">设备</span>
        <span class="name">{{ charge.device.name }}</span>
        <span class="label">类型</span>
        <span class="type">{{ charge.device.type }}</span>
      </div>

      <div class="user-data">
        <div class="avatar-name">
          <img class="avatar item" :src="charge.user.avatarURL || '/images/avatar.jpg'" />
          <span class="name item">{{ charge.user.userExtend.name }}</span>
          <span class="phone item">
            <i class="el-icon-mobile-phone"></i>
            {{ charge.user.phone}}
          </span>
          <span class="email item">
            <i class="iconfont icon-185078emailmailstreamline"></i>
            {{ charge.user.userExtend.email}}
          </span>
        </div>
      </div>

      <div class="privs-data">
        <div class="title">
          <i class="el-icon-s-management"></i>
          <span>权限</span>
          <el-button
            round
            icon="el-icon-plus"
            size="small"
            type="primary"
            @click="addAbilityFormShow = true"
          >增加</el-button>
        </div>

        <el-form
          :inline="true"
          v-if="addAbilityFormShow"
          size="mini"
          :model="form"
          :rules="rules"
          ref="form"
        >
          <el-form-item prop="privilegeID">
            <el-select
              v-model="form.privilegeID"
              remote
              filterable
              placeholder="选择权限"
              :remote-method="querySearchPrivs"
            >
              <el-option
                v-for="priv in devicePrivOptions"
                :key="priv.id"
                :label="priv.name"
                :value="priv.id"
              ></el-option>
            </el-select>
          </el-form-item>

          <el-form-item>
            <el-button round icon="el-icon-finished" type="primary" @click="submit">保存</el-button>
          </el-form-item>

          <el-form-item>
            <el-button round icon="el-icon-close" type="info" @click="hideForm">取消</el-button>
          </el-form-item>
        </el-form>

        <ability-item v-for="a in charge.abilities" :key="a.id" :priv="a"></ability-item>
      </div>
    </div>
  </div>
</template>
<script>
import AbilityItem from './_ability-item';

import chargeQuery from './gql/query.charger.gql';
import devicePrivsQuery from './gql/query.device-privs.gql';
import abilityCreate from './gql/mutation.ability-create.gql';

import { parseGQLError } from 'js/utils'

export default {
  name: 'charge-show',
  props: ['id'],
  components: { AbilityItem },
  apollo: {
    charge: {
      query: chargeQuery,
      variables() {
        return this.variables;
      }
    },
    devicePrivs: {
      query: devicePrivsQuery,
      variables() {
        return {
          privType: 'device',
          namePattern: this.privNamePattern
        };
      }
    }
  },
  data() {
    return {
      rules: {
        privilegeID: [
          { required: true, trigger: 'blur', message: '请选择权限' }
        ]
      },
      form: {
        privilegeID: null,
        deviceChargeID: this.id
      },
      privNamePattern: '',
      devicePrivs: [],
      charge: {},
      variables: { id: this.id },
      addAbilityFormShow: false
    };
  },
  computed: {
    devicePrivOptions() {
      return this.devicePrivs.filter(dp => {
        return (
          this.charge.abilities.findIndex(a => a.privilege.id === dp.id) === -1
        );
      });
    }
  },
  methods: {
    addAbility() {},
    querySearchPrivs(query) {
      this.privNamePattern = query;
    },
    hideForm() {
      this.addAbilityFormShow = false;
      this.$refs.form.resetFields();
    },
    submit() {
      this.$refs.form.validate(valid => {
        if (valid)
          this.$apollo
            .mutate({
              mutation: abilityCreate,
              variables: this.form,
              update: (store, { data: { ability } }) => {
                var opts = {
                  query: chargeQuery,
                  variables: this.variables
                };
                var data = store.readQuery(opts);
                data.charge.abilities.unshift(ability);
                store.writeQuery({ ...opts, data });
              }
            })
            .then(() => {
              this.hideForm();
            })
            .catch(e =>
              this.$message({
                type: 'error',
                message: parseGQLError(e).message
              })
            );
      });
    }
  }
};
</script>
<style lang="scss">
@import 'css/main/charge/_show.scss';
</style>
