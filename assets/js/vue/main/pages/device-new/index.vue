<template>
  <div class="device-new">
    <div class="form-body global-card">
      <div class="form-title">
        <i class="el-icon-s-order"></i> 注册设备
      </div>
      <device-form ref="deviceForm"></device-form>
      <params-form ref="paramsForm"></params-form>
      <charge-form ref="chargeForm"></charge-form>

      <el-form class="device-form" label-width="100px">
        <el-form-item style="text-align: center">
          <el-button :loading="loading" @click="submit()" style="width: 100%;" type="primary">提交</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>
<script>
import DeviceForm from './_device-form';
import ParamsForm from './_params-form';
import ChargeForm from './_charge-form';
import gql from 'graphql-tag';

export default {
  name: 'device-new',
  components: {
    DeviceForm,
    ParamsForm,
    ChargeForm
  },
  data() {
    return {
      loading: false
    };
  },
  methods: {
    submit() {
      var _this = this;
      _this.loading = true;
      this.$refs.deviceForm.$refs.form.validate(valid => {
        if (valid) {
          this.$apollo
            .mutate({
              mutation: gql`
                mutation(
                  $name: String!
                  $type: String!
                  $description: String!
                ) {
                  deviceCreate(
                    name: $name
                    type: $type
                    description: $description
                  ) {
                    id
                    uuid
                  }
                }
              `,
              variables: {
                ..._this.$refs.deviceForm.form
              }
            })
            .then(({ data: { deviceCreate: { id, uuid } } }) => {
              _this.$refs.paramsForm.params.forEach(param => {
                _this.$apollo
                  .mutate({
                    mutation: gql`
                      mutation(
                        $name: String!
                        $sign: String!
                        $type: DeviceParamValueType!
                        $deviceID: Int!
                      ) {
                        deviceParamCreate(
                          name: $name
                          sign: $sign
                          type: $type
                          deviceID: $deviceID
                        ) {
                          id
                        }
                      }
                    `,
                    variables: {
                      ...param,
                      deviceID: parseInt(id)
                    }
                  })
                  .catch(e => {
                    console.log(e);
                  });
              });

              _this.$refs.chargeForm.charges.forEach(charge => {
                _this.$apollo
                  .mutate({
                    mutation: gql`
                      mutation(
                        $uuid: String!
                        $userUUID: String!
                        $privIDs: [Int]!
                      ) {
                        deviceChargeCreate(
                          uuid: $uuid
                          userUUID: $userUUID
                          privIDs: $privIDs
                        ) {
                          id
                        }
                      }
                    `,
                    variables: {
                      uuid,
                      ...charge
                    }
                  })
                  .catch(e => {
                    console.log(e);
                  });
              });

              _this.loading = false;
            });
        } else {
          _this.loading = false;
        }
      });
    }
  }
};
</script>
<style lang="scss">
@import 'css/vars.scss';

.device-new {
  height: calc(100% - 50px);
  overflow: auto;
  color: $--color-font__light;
}

.device-new .form-title {
  font-size: 1.5rem;
  width: 600px;
  margin: auto;
  margin-bottom: 2rem;
  padding-bottom: 0.5rem;
  border-bottom: 1px solid $--color-theme__black;
}

.device-new .form-body {
  padding: 1.5rem;
  margin: auto;
}

.device-new .device-form {
  width: 600px;
  margin: auto;
}

.charge-form .el-transfer .el-transfer-panel__header .el-checkbox {
  line-height: 0;
  padding-top: 10px;
}

.device-new .device-form {
  .el-form-item__label {
    color: $--color-font__light;
  }
}

.device-new .device-form .el-tag {
  margin-right: 0.5rem;
  cursor: pointer;
}

// .device-new .charge-form .el-transfer {
// .el-checkbox__label {
// color: $--color-font__light;
// }
// }

@media only screen and(max-width: 664px) {
  .device-new .form-body {
    padding: 1.5rem 0.5rem;
    min-width: 316px;
    margin: 0.5rem;
  }

  .device-form .el-input,
  .device-form .el-textarea {
    width: 200px;
  }

  .device-new .form-title,
  .device-new .device-form {
    width: 300px;
    margin: auto;
  }

  .device-new .form-title {
    font-size: 1.2rem;
    margin-bottom: 1.5rem;
  }

  .device-new .charge-form {
    .el-transfer__buttons {
      width: 100%;
      padding: 0.5rem;
      text-align: center;
    }

    .el-button {
      display: inline-block;
      margin: 0;

      &:first-child {
        margin-right: 1rem;
      }

      i {
        transform: rotate(90deg);
      }
    }
  }
}
</style>
