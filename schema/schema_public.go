package schema

import (
	"log"

	"github.com/graphql-go/graphql"

	fields "github.com/SasukeBo/information/schema/fields/public"
)

var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"deviceParamValueAdd": fields.DeviceParamValueAddField,
	},
})

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

		/* userExtend */
		"userExtendUpdate":    fields.UserExtendUpdateField,
		"userExtendBindEmail": fields.UserExtendBindEmailField,

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
		"deviceGet":            fields.DeviceGetField,
		"deviceList":           fields.DeviceListField,
		"deviceChargeGet":      fields.DeviceChargeGetField,
		"deviceChargeList":     fields.DeviceChargeListField,
		"deviceParamGet":       fields.DeviceParamGetField,
		"deviceParamList":      fields.DeviceParamListField,
		"deviceParamValueList": fields.DeviceParamValueListField,
		"deviceStatusLogList":  fields.DeviceStatusLogListField,
		/*
			FIXME: 暂时不需要的接口
			"deviceChargePrivGet":  fields.DeviceChargePrivGetField,
			"deviceChargePrivList": fields.DeviceChargePrivListField,
		*/

		/* userLogin */
		"userLoginList": fields.UserLoginListField,
		"getLastLogin":  fields.UserLoginLastField,
		"getThisLogin":  fields.UserLoginThisField,

		/* privilege */
		"privilegeList": fields.PrivilegeListField,
	},
})

// PublicSchema is graphql schema
var PublicSchema graphql.Schema

func init() {
	var err error
	PublicSchema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:        QueryRoot,
		Mutation:     MutateRoot,
		Subscription: Subscription,
	})
	if err != nil {
		log.Fatal("failed to create public schema, err: ", err)
	}
}
