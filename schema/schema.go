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
		"sayHello": types.SayHelloType,
		"whoAmI":   types.WhoAmIType,
		/*    role    */
		"roleGet":       types.RoleGetType,
		"roleGetByName": types.RoleGetByNameType,
		/*    aliyun    */
		"getSmsCode": types.GetSmsCodeType,
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

		/*    role    */
		"roleCreate": types.RoleCreateType,
		"roleUpdate": types.RoleUpdateType,

		/*    aliyun    */
		"sendSmsCode": types.SendSmsCodeType,

		/*    device    */
		"deviceCreate": types.DeviceCreateType,
		"deviceBind":   types.DeviceBindType,
		"deviceCharge": types.DeviceChargeType,
		// "deviceUNCharge": types.DeviceUNCharge,
		// "deviceRECharge": types.DeviceRECharge,

		/*    device    */
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
		log.Fatal("failed to create new schema, err: ", err)
	}
}
