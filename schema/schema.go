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
		// user
		"currentUser":   currentUser,   // 获取当前登录的用户
		"userGet":       userGet,       // 获取用户
		"userList":      userList,      // 用户列表
		"getSmsCode":    getSmsCode,    // 获取短信验证码
		"userLoginList": userLoginList, // 用户登录记录
		"getLastLogin":  userLoginLast, // 用户上次登录记录
		"getThisLogin":  userLoginThis, // 用户本次登录记录
		"privilegeList": privilegeList, // TODO: 移除

		// device
		"deviceGet":               deviceGet,               // 获取设备
		"deviceList":              deviceList,              // 获取设备列表
		"deviceTokenGet":          deviceTokenGet,          // Token查询设备 TODO: 移除
		"deviceStatusCount":       deviceStatusCount,       // 设备分运行状态统计数量
		"deviceStopLogList":       deviceStopLogList,       // 设备停机日志列表
		"deviceStopTypeCount":     deviceStopTypeCount,     // 设备停机类型次数统计
		"deviceMonthlyStatistics": deviceMonthlyStatistics, // 设备月数据分析
		"deviceStatusStatistics":  deviceStatusStatistics,  // 设备状态数据
		"logStopReasonsGet":       logStopReasonsGet,       // 停机日志获取停机原因

		// product
		"productGet":         productGet,         // 获取产品
		"productList":        productList,        // 获取产品列表
		"productHistogram":   productHistogram,   // 获取产品检测参数直方图
		"detectItemGet":      detectItemGet,      // 产品检测项获取
		"detectItemList":     detectItemList,     // 产品检测项列表
		"realTimeStatistics": realTimeStatistics, // 产品实时数据
		"productOverview":    productOverview,    // 产品总览数据
		"productDevicesGet":  productDevicesGet,  // 产品关联设备获取
	},
})

// MutateRoot is mutation root
var MutateRoot = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// User
		"signUp":        signUp,        // 注册
		"signIn":        signIn,        // 登录
		"signOut":       signOut,       // 登出
		"resetPassword": resetPassword, // 找回密码
		"userUpdate":    userUpdate,    // 更新用户信息

		// aliyun
		"sendSmsCode": sendSmsCode, // 发送短信验证码

		// device
		"deviceCreate":     deviceCreate,     // 设备创建
		"deviceUpdate":     deviceUpdate,     // 设备更新
		"deviceDelete":     deviceDelete,     // 设备移除
		"stopReasonCreate": stopReasonCreate, // 停机原因创建
		"stopReasonUpdate": stopReasonUpdate, // 停机原因更新
		"stopReasonDelete": stopReasonDelete, // 停机原因删除

		// product
		"productCreate":    productCreate,    // 产品创建
		"productDelete":    productDelete,    // 产品移除
		"productUpdate":    productUpdate,    // 产品更新
		"detectItemCreate": detectItemCreate, // 产品检测项增加
		"detectItemUpdate": detectItemUpdate, // 产品检测项更新
		"detectItemDelete": detectItemDelete, // 产品检测项删除

	},
})

// Subscription _
var Subscription = graphql.NewObject(graphql.ObjectConfig{
	Name: "Subscription",
	Fields: graphql.Fields{
		"deviceStatusUpdate": deviceStatusUpdate,
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
