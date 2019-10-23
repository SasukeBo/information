package schema

import (
	"github.com/graphql-go/graphql"
	"log"
)

// Root is graphql schema root
var Root graphql.Schema

// QueryRoot is query root
var QueryRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"currentUser":             currentUser,
		"userGet":                 userGet,
		"userList":                userList,
		"getSmsCode":              getSmsCode,
		"deviceGet":               deviceGet,
		"deviceList":              deviceList,
		"deviceTokenGet":          deviceTokenGet,
		"deviceStatusCount":       deviceStatusCount,
		"deviceChargerGet":        deviceChargerGet,
		"deviceChargerList":       deviceChargerList,
		"deviceStatusLogList":     deviceStatusLogList,
		"deviceStatusDuration":    deviceStatusDuration,
		"deviceMonthlyStatistics": deviceMonthlyStatistics,
		"userLoginList":           userLoginList,
		"getLastLogin":            userLoginLast,
		"getThisLogin":            userLoginThis,
		"privilegeList":           privilegeList,
		"productGet":              productGet,
		"productList":             productList,
		"detectItemGet":           detectItemGet,
		"detectItemList":          detectItemList,
		"realTimeStatistics":      realTimeStatistics,
		"productHistogram":        productHistogram,
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"signUp":              signUp,
		"signIn":              signIn,
		"signOut":             signOut,
		"resetPassword":       resetPassword,
		"userUpdate":          userUpdate,
		"sendSmsCode":         sendSmsCode,
		"deviceCreate":        deviceCreate,
		"deviceUpdate":        deviceUpdate,
		"deviceDelete":        deviceDelete,
		"deviceChargerCreate": deviceChargerCreate,
		"deviceChargerDelete": deviceChargerDelete,
		"deviceChargerUpdate": deviceChargerUpdate,
		// "deviceProductShipCreate": deviceProductShipCreate,
		"productCreate":    productCreate,
		"productDelete":    productDelete,
		"productUpdate":    productUpdate,
		"detectItemCreate": detectItemCreate,
		"detectItemUpdate": detectItemUpdate,
		"detectItemDelete": detectItemDelete,
		// "productUpdate":       productUpdate,
	},
})

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		// "deviceParamValueSub": deviceParamValueAdd,
		"deviceStatusRefresh": deviceStatusRefresh,
	},
})

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
