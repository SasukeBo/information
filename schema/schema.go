package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		// "deviceParamValueAdd": DeviceParamValueAddField,
		"deviceStatusRefresh": DeviceStatusRefreshField,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*    user    */
		"signUp":        SignUpField,
		"signIn":        SignInField,
		"signOut":       SignOutField,
		"resetPassword": ResetPasswordField,
		"updateUser":    UserUpdateField,

		/*    aliyun    */
		"sendSmsCode": SendSmsCodeField,

		/*    device    */
		"deviceCreate": DeviceCreateField,
		"deviceUpdate": DeviceUpdateField,
		"deviceDelete": DeviceDeleteField,

		/*    deviceCharge    */
		"deviceChargerCreate": DeviceChargerCreateField,
		"deviceChargerDelete": DeviceChargerDeleteField,
		"deviceChargerUpdate": DeviceChargerUpdateField,

		/*    deviceParam    */
		"deviceParamCreate": DeviceParamCreateField,
		"deviceParamUpdate": DeviceParamUpdateField,
		"deviceParamDelete": DeviceParamDeleteField,
	},
})

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/* user */
		"currentUser": CurrentUserField,
		"userGet":     UserGetField,
		"userList":    UserListField,

		/*    aliyun    */
		"getSmsCode": GetSmsCodeField,

		/*    device    */
		"deviceGet":         DeviceGetField,
		"deviceList":        DeviceListField,
		"deviceTokenGet":    DeviceTokenGetField,
		"deviceStatusCount": DeviceStatusCountField,

		"deviceChargerGet":  DeviceChargerGetField,
		"deviceChargerList": DeviceChargerListField,

		"deviceParamGet":  DeviceParamGetField,
		"deviceParamList": DeviceParamListField,

		"deviceParamValueList":      DeviceParamValueListField,
		"deviceParamValueCount":     DeviceParamValueCountField,
		"deviceParamValueHistogram": DeviceParamValueHistogramField,

		"deviceStatusLogList":  DeviceStatusLogListField,
		"deviceStatusDuration": DeviceStatusDurationField,

		/* userLogin */
		"userLoginList": UserLoginListField,
		"getLastLogin":  UserLoginLastField,
		"getThisLogin":  UserLoginThisField,

		/* privilege */
		"privilegeList": PrivilegeListField,
	},
})

// Root is graphql schema root
var Root graphql.Schema

func init() {
	var err error
	Root, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:        QueryRoot,
		Mutation:     MutateRoot,
		Subscription: Subscription,
	})
	if err != nil {
		log.Fatal("failed to create public schema, err: ", err)
	}
}

// GenArg 简化gql参数定义
func GenArg(gqlType graphql.Input, des string, opts ...interface{}) *graphql.ArgumentConfig {
	defaultValue := interface{}(nil)
	if len(opts) > 0 && !opts[0].(bool) {
		gqlType = graphql.NewNonNull(gqlType)
	}

	if len(opts) > 1 {
		defaultValue = opts[1]
	}

	return &graphql.ArgumentConfig{
		Type:         gqlType,
		Description:  des,
		DefaultValue: defaultValue,
	}
}
