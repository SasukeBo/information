package schema

import (
	"log"

	"github.com/graphql-go/graphql"

	fields "github.com/SasukeBo/information/schema/fields/public"
)

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"deviceParamValueAdd": fields.DeviceParamValueAddField,
		"deviceStatusRefresh": fields.DeviceStatusRefreshField,
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

		/*    aliyun    */
		"sendSmsCode": fields.SendSmsCodeField,

		/*    device    */
		"deviceCreate": fields.DeviceCreateField,
		"deviceUpdate": fields.DeviceUpdateField,
		"deviceDelete": fields.DeviceDeleteField,

		/*    deviceCharge    */
		"deviceChargerCreate": fields.DeviceChargerCreateField,
		"deviceChargerDelete": fields.DeviceChargerDeleteField,
		"deviceChargerUpdate": fields.DeviceChargerUpdateField,

		/*    deviceParam    */
		"deviceParamCreate": fields.DeviceParamCreateField,
		"deviceParamUpdate": fields.DeviceParamUpdateField,
		"deviceParamDelete": fields.DeviceParamDeleteField,
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
		"deviceGet":         fields.DeviceGetField,
		"deviceList":        fields.DeviceListField,
		"deviceTokenGet":    fields.DeviceTokenGetField,
		"deviceStatusCount": fields.DeviceStatusCountField,

		"deviceChargerGet":  fields.DeviceChargerGetField,
		"deviceChargerList": fields.DeviceChargerListField,

		"deviceParamGet":  fields.DeviceParamGetField,
		"deviceParamList": fields.DeviceParamListField,

		"deviceParamValueList":      fields.DeviceParamValueListField,
		"deviceParamValueCount":     fields.DeviceParamValueCountField,
		"deviceParamValueHistogram": fields.DeviceParamValueHistogramField,

		"deviceStatusLogList":  fields.DeviceStatusLogListField,
		"deviceStatusDuration": fields.DeviceStatusDurationField,

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
