package schema

import (
	"github.com/SasukeBo/information/schema/types"
	"github.com/graphql-go/graphql"
	"log"
)

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/*    test    */
		"whoIAm": types.WhoIAmType,

		/* user */
		"userGet":  types.UserGet,
		// "userList": types.UserList, // TODO: 查询条件限制

		/*    aliyun    */
		"getSmsCode": types.GetSmsCodeType,

		/*    device    */
		"deviceGet":  types.DeviceGetType,
		"deviceList": types.DeviceListType,

		/*    deviceCharge    */
		"deviceChargeGet":  types.DeviceChargeGetType,
		"deviceChargeList": types.DeviceChargeListType,

		/*    deviceParam    */
		"deviceParamGet":  types.DeviceParamGetType,
		"deviceParamList": types.DeviceParamListType,

		/*    deviceChargeAbility    */
		"deviceChargePrivGet":  types.DeviceChargePrivGetType,
		"deviceChargePrivList": types.DeviceChargePrivListType,

		/*    deviceParamValue    */
		"deviceParamValueList": types.DeviceParamValueListType,

		/*    deviceStatusLog    */
		"deviceStatusLogList": types.DeviceStatusLogListType,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*    user    */
		"register":        types.UserCreateType,
		"resetPassword":   types.ResetPasswordType,
		"loginByPassword": types.LoginByPasswordType,
		"logout":          types.LogoutType,
		// TODO:"userUpdate"

		/*    aliyun    */
		"sendSmsCode": types.SendSmsCodeType,

		/*    device    */
		"deviceCreate": types.DeviceCreateType,
		"deviceUpdate": types.DeviceUpdateType,
		"deviceDelete": types.DeviceDeleteType,
		"deviceBind":   types.DeviceBindType,

		/*    deviceCharge    */
		"deviceChargeCreate": types.DeviceChargeCreateType,
		"deviceChargeDelete": types.DeviceChargeDeleteType,
		"deviceChargeUpdate": types.DeviceChargeUpdateType,

		/*    deviceParam    */
		"deviceParamCreate": types.DeviceParamCreateType,
		"deviceParamUpdate": types.DeviceParamUpdateType,
		"deviceParamDelete": types.DeviceParamDeleteType,

		/*    deviceChargeAbility    */
		"deviceChargePrivCreate": types.DeviceChargePrivCreateType,
		"deviceChargePrivDelete": types.DeviceChargePrivDeleteType,
	},
})

// Schema is graphql schema
var Schema graphql.Schema

func init() {
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryRoot,
		Mutation: MutateRoot,
	})
	if err != nil {
		log.Fatal("failed to create schema, err: ", err)
	}
}
