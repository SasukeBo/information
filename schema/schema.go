package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"deviceParamValueSub": deviceParamValueSubField,
		"deviceStatusRefresh": deviceStatusRefreshField,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*    user    */
		"signUp":        signUpField,
		"signIn":        signInField,
		"signOut":       signOutField,
		"resetPassword": resetPasswordField,
		"updateUser":    userUpdateField,

		/*    aliyun    */
		"sendSmsCode": sendSmsCodeField,

		/*    device    */
		"deviceCreate": deviceCreateField,
		"deviceUpdate": deviceUpdateField,
		"deviceDelete": deviceDeleteField,

		/*    deviceCharge    */
		"deviceChargerCreate": deviceChargerCreateField,
		"deviceChargerDelete": deviceChargerDeleteField,
		"deviceChargerUpdate": deviceChargerUpdateField,

		/*    deviceParam    */
		"deviceParamCreate": deviceParamCreateField,
		"deviceParamUpdate": deviceParamUpdateField,
		"deviceParamDelete": deviceParamDeleteField,
	},
})

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/* user */
		"currentUser": currentUserField,
		"userGet":     userGetField,
		"userList":    userListField,

		/*    aliyun    */
		"getSmsCode": getSmsCodeField,

		/*    device    */
		"deviceGet":         deviceGetField,
		"deviceList":        deviceListField,
		"deviceTokenGet":    deviceTokenGetField,
		"deviceStatusCount": deviceStatusCountField,

		"deviceChargerGet":  deviceChargerGetField,
		"deviceChargerList": deviceChargerListField,

		"deviceParamGet":  deviceParamGetField,
		"deviceParamList": deviceParamListField,

		"deviceParamValueList":      deviceParamValueListField,
		"deviceParamValueCount":     deviceParamValueCountField,
		"deviceParamValueHistogram": deviceParamValueHistogramField,

		"deviceStatusLogList":  deviceStatusLogListField,
		"deviceStatusDuration": deviceStatusDurationField,

		/* userLogin */
		"userLoginList": userLoginListField,
		"getLastLogin":  userLoginLastField,
		"getThisLogin":  userLoginThisField,

		/* privilege */
		"privilegeList": privilegeListField,
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
