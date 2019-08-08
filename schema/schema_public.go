package schema

import (
  "log"

  "github.com/graphql-go/graphql"

  fields "github.com/SasukeBo/information/schema/fields/public"
)

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootMutation",
  Fields: graphql.Fields{
    /*    user    */
    "register":           fields.UserCreateField,
    "resetPassword":      fields.ResetPasswordField,
    "loginByPassword":    fields.LoginByPasswordField,
    "logout":             fields.LogoutField,
    "userUpdateAvatar":   fields.UserUpdateAvatarField,
    "userUpdatePassword": fields.UserUpdatePasswordField,
    "userUpdatePhone":    fields.UserUpdatePhoneField,

    /*    aliyun    */
    "sendSmsCode": fields.SendSmsCodeField,

    /*    device    */
    "deviceCreate": fields.DeviceCreateField,
    "deviceUpdate": fields.DeviceUpdateField,
    "deviceDelete": fields.DeviceDeleteField,
    "deviceBind":   fields.DeviceBindField,

    /*    deviceCharge    */
    "deviceChargeCreate": fields.DeviceChargeCreateField,
    "deviceChargeDelete": fields.DeviceChargeDeleteField,
    "deviceChargeUpdate": fields.DeviceChargeUpdateField,

    /*    deviceParam    */
    "deviceParamCreate": fields.DeviceParamCreateField,
    "deviceParamUpdate": fields.DeviceParamUpdateField,
    "deviceParamDelete": fields.DeviceParamDeleteField,

    /*    deviceChargeAbility    */
    "deviceChargePrivCreate": fields.DeviceChargePrivCreateField,
    "deviceChargePrivDelete": fields.DeviceChargePrivDeleteField,
  },
})

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
  Name: "RootQuery",
  Fields: graphql.Fields{
    /* user */
    "currentUser": fields.CurrentUserField,
    "userGet":     fields.UserGetField,
    "userList":    fields.UserListField,

    /*    aliyun    */
    "getSmsCode": fields.GetSmsCodeField,

    /*    device    */
    "deviceGet":  fields.DeviceGetField,
    "deviceList": fields.DeviceListField,

    /*    deviceCharge    */
    "deviceChargeGet":  fields.DeviceChargeGetField,
    "deviceChargeList": fields.DeviceChargeListField,

    /*    deviceParam    */
    "deviceParamGet":  fields.DeviceParamGetField,
    "deviceParamList": fields.DeviceParamListField,

    /*    deviceChargeAbility    */
    "deviceChargePrivGet":  fields.DeviceChargePrivGetField,
    "deviceChargePrivList": fields.DeviceChargePrivListField,

    /*    deviceParamValue    */
    "deviceParamValueList": fields.DeviceParamValueListField,

    /*    deviceStatusLog    */
    "deviceStatusLogList": fields.DeviceStatusLogListField,
  },
})

// PublicSchema is graphql schema
var PublicSchema graphql.Schema

func init() {
  var err error
  PublicSchema, err = graphql.NewSchema(graphql.SchemaConfig{
    Query:    QueryRoot,
    Mutation: MutateRoot,
  })
  if err != nil {
    log.Fatal("failed to create public schema, err: ", err)
  }
}
