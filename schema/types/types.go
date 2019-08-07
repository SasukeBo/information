package types

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/scalars"
)

// SendSmsCodeResponse response type of SendSmsCode
var SendSmsCodeResponse graphql.Type

// DeviceChargeAbility 设备负责人权限
var DeviceChargeAbility graphql.Type

// DeviceCharge 设备类型
var DeviceCharge graphql.Type

// DeviceParam 设备参数
var DeviceParam graphql.Type

// DeviceParamValue 设备参数值
var DeviceParamValue graphql.Type

// DeviceStatusLog 设备状态变更记录
var DeviceStatusLog graphql.Type

// Device 设备类型
var Device graphql.Type

// Privilege 用户类型
var Privilege graphql.Type

// Response 消息体
var Response graphql.Type

// RolePriv 用户类型
var RolePriv graphql.Type

// Role 用户类型
var Role graphql.Type

// UserExtend 用户类型
var UserExtend graphql.Type

// User 用户类型
var User graphql.Type

func init() {
	// SendSmsCodeResponse 短信验证码response消息体
	SendSmsCodeResponse = graphql.NewObject(graphql.ObjectConfig{
		Name:        "SendSmsCodeResponse",
		Description: "短信验证码response消息体",
		Fields: graphql.Fields{
			"message":   &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
			"requestID": &graphql.Field{Type: graphql.String, Description: "请求ID"},
			"code": &graphql.Field{Type: graphql.String, Description: `
			请求状态码。
			· 返回OK代表请求成功
			· 其他错误码见阿里云短信服务错误码列表
			`},
			"bizID": &graphql.Field{Type: graphql.String, Description: "状态码的描述"},
		},
	})

	DeviceChargeAbility = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceChargeAbility",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":           &graphql.Field{Type: graphql.Int},
				"deviceCharge": &graphql.Field{Type: DeviceCharge, Description: "设备负责人关系"},
				"privilege":    &graphql.Field{Type: Privilege, Description: "权限"},
			}
		}),
	})

	DeviceCharge = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceCharge",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"device":    &graphql.Field{Type: Device, Description: "设备"},
				"user":      &graphql.Field{Type: User, Description: "负责人"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	DeviceParam = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceParam",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"name":      &graphql.Field{Type: graphql.String, Description: "参数名称"},
				"sign":      &graphql.Field{Type: graphql.String, Description: "参数签名"},
				"type":      &graphql.Field{Type: graphql.String, Description: "参数值类型"},
				"author":    &graphql.Field{Type: User, Description: "创建人"},
				"createdAt": &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
			}
		}),
	})

	DeviceParamValue = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceParamValue",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":          &graphql.Field{Type: graphql.Int},
				"value":       &graphql.Field{Type: graphql.String, Description: "参数值字符串"},
				"deviceParam": &graphql.Field{Type: DeviceParam, Description: "设备参数"},
				"createdAt":   &graphql.Field{Type: graphql.DateTime, Description: "创建时间"},
			}
		}),
	})

	DeviceStatusLog = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceStatusLog",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"status":   &graphql.Field{Type: scalars.DeviceStatus, Description: "设备运行状态"},
				"device":   &graphql.Field{Type: Device, Description: "设备"},
				"changeAt": &graphql.Field{Type: graphql.DateTime, Description: "变更时间"},
			}
		}),
	})

	Device = graphql.NewObject(graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"type":        &graphql.Field{Type: graphql.String, Description: "设备类型"},
				"name":        &graphql.Field{Type: graphql.String, Description: "设备名称"},
				"mac":         &graphql.Field{Type: graphql.String, Description: "设备Mac地址"},
				"token":       &graphql.Field{Type: graphql.String, Description: "设备token，用于数据加密"},
				"status":      &graphql.Field{Type: scalars.BaseStatus, Description: "基础状态"},
				"id":          &graphql.Field{Type: graphql.Int},
				"uuid":        &graphql.Field{Type: graphql.String, Description: "设备UUID"},
				"user":        &graphql.Field{Type: User, Description: "注册人用户"},
				"description": &graphql.Field{Type: graphql.String, Description: "设备描述，备注"},
				"createdAt":   &graphql.Field{Type: graphql.DateTime},
				"updatedAt":   &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	Privilege = graphql.NewObject(graphql.ObjectConfig{
		Name: "Privilege",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"privName":  &graphql.Field{Type: graphql.String},
				"privType":  &graphql.Field{Type: graphql.Int},
				"status":    &graphql.Field{Type: scalars.BaseStatus, Description: "基础状态"},
				"rolePrivs": &graphql.Field{Type: graphql.NewList(RolePriv), Description: "role and privilege relationship"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	RolePriv = graphql.NewObject(graphql.ObjectConfig{
		Name: "RolePriv",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"role":      &graphql.Field{Type: Role},
				"privilege": &graphql.Field{Type: Privilege},
			}
		}),
	})

	Role = graphql.NewObject(graphql.ObjectConfig{
		Name: "Role",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"roleName":  &graphql.Field{Type: graphql.String, Description: "role name"},
				"status":    &graphql.Field{Type: scalars.BaseStatus, Description: "role status, can be default, publish, block and deleted"},
				"rolePrivs": &graphql.Field{Type: graphql.NewList(RolePriv), Description: "role and privilege relationship"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	UserExtend = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserExtend",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":    &graphql.Field{Type: graphql.Int},
				"user":  &graphql.Field{Type: User},
				"name":  &graphql.Field{Type: graphql.String},
				"email": &graphql.Field{Type: graphql.String},
			}
		}),
	})

	User = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"uuid":       &graphql.Field{Type: graphql.String, Description: "通用唯一标识"},
				"phone":      &graphql.Field{Type: graphql.String, Description: "手机号"},
				"avatarURL":  &graphql.Field{Type: graphql.String, Description: "头像链接"},
				"role":       &graphql.Field{Type: Role, Description: "用户角色"},
				"userExtend": &graphql.Field{Type: UserExtend, Description: "用户拓展信息"},
				"status":     &graphql.Field{Type: scalars.BaseStatus, Description: "基础状态"},
				"createdAt":  &graphql.Field{Type: graphql.DateTime},
				"updatedAt":  &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})
}
