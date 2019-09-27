package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"deviceParamValueSub": deviceParamValueAdd,
		"deviceStatusRefresh": deviceStatusRefresh,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		/*    user    */
		"signUp":        signUp,
		"signIn":        signIn,
		"signOut":       signOut,
		"resetPassword": resetPassword,
		"updateUser":    userUpdate,

		/*    aliyun    */
		"sendSmsCode": sendSmsCode,

		/*    device    */
		"deviceCreate": deviceCreate,
		"deviceUpdate": deviceUpdate,
		"deviceDelete": deviceDelete,

		/*    deviceCharge    */
		"deviceChargerCreate": deviceChargerCreate,
		"deviceChargerDelete": deviceChargerDelete,
		"deviceChargerUpdate": deviceChargerUpdate,

		/*    deviceParam    */
		"deviceParamCreate": deviceParamCreate,
		"deviceParamUpdate": deviceParamUpdate,
		"deviceParamDelete": deviceParamDelete,
	},
})

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		/* user */
		"currentUser": currentUser,
		"userGet":     userGet,
		"userList":    userList,

		/*    aliyun    */
		"getSmsCode": getSmsCode,

		/*    device    */
		"deviceGet":         deviceGet,
		"deviceList":        deviceList,
		"deviceTokenGet":    deviceTokenGet,
		"deviceStatusCount": deviceStatusCount,

		"deviceChargerGet":  deviceChargerGet,
		"deviceChargerList": deviceChargerList,

		"deviceParamGet":  deviceParamGet,
		"deviceParamList": deviceParamList,

		"deviceParamValueList":      deviceParamValueList,
		"deviceParamValueCount":     deviceParamValueCount,
		"deviceParamValueHistogram": deviceParamValueHistogram,

		"deviceStatusLogList":  deviceStatusLogList,
		"deviceStatusDuration": deviceStatusDuration,

		/* userLogin */
		"userLoginList": userLoginList,
		"getLastLogin":  userLoginLast,
		"getThisLogin":  userLoginThis,

		/* privilege */
		"privilegeList": privilegeList,
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
