<template>
  <div class="device-form charge-form">
    <el-form label-position="left" label-width="100px">
      <el-form-item label="设备负责人">
        <el-tag
          v-for="charge in charges"
          :key="charge.uuid"
          closable
          @click="edit(charge)"
          @close="removeCharge(charge)"
        >{{ charge.name }}</el-tag>
        <el-button
          @click="formVisible = true"
          size="small"
          type="primary"
          v-show="!formVisible"
          style="margin-right: 1rem"
        >增加</el-button>
        <span v-if="charges.length == 0" style="color: #606266">暂无负责人</span>
      </el-form-item>
    </el-form>

    <el-form
      :model="form"
      label-position="left"
      label-width="100px"
      v-if="formVisible"
      ref="chargeForm"
    >
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
        <el-button size="small" type="primary" @click="saveCharge()">保存</el-button>
        <el-button size="small" type="warning" @click="cancelCharge()">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
import { apollo } from './graphql';
import gql from 'graphql-tag';
import { Autocomplete, Transfer } from 'element-ui';

export default {
  name: 'charge-form',
  components: {
    ElAutocomplete: Autocomplete,
    ElTransfer: Transfer
  },
  apollo,
  data() {
    return {
      devicePrivs: [],
      formVisible: false,
      form: {
        name: '',
        userUUID: '',
        privIDs: []
      },
      charges: []
    };
  },
  methods: {
    removeCharge(charge) {
      var index = this.charges.findIndex(c => c.userUUID === charge.userUUID);
      this.charges.splice(index, 1);
    },
    querySearchUsers(queryString, callback) {
      this.$apollo
        .query({
          query: gql`
            query($namePattern: String) {
              userList(namePattern: $namePattern) {
                uuid
                userExtend {
                  name
                }
              }
            }
          `,
          variables: {
            namePattern: queryString
          }
        })
        .then(({ data: { userList } }) => {
          var data = userList.map(user => {
            return { name: user.userExtend.name, uuid: user.uuid };
          });
          callback(data);
        });
    },
    handleSelect(data) {
      this.form.userUUID = data.uuid;
    },
    saveCharge() {
      var charge = {
        ...this.form
      };
      this.charges.push(charge);
      this.formVisible = false;
      this.$refs.chargeForm.resetFields();
    },
    cancelCharge() {
      this.formVisible = false;
      this.$refs.chargeForm.resetFields();
    },
    edit(charge) {
      this.form = {
        ...charge
      };
      this.formVisible = true;
    }
  }
};
</script>
