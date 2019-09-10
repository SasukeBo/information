package types

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/resolvers/device"
	"github.com/SasukeBo/information/schema/resolvers/privilege"
	"github.com/SasukeBo/information/schema/resolvers/role"
	"github.com/SasukeBo/information/schema/resolvers/rolepriv"
	"github.com/SasukeBo/information/schema/resolvers/user"
	"github.com/SasukeBo/information/schema/resolvers/userextend"
	"github.com/SasukeBo/information/schema/scalars"
)

// SendSmsCodeResponse response type of SendSmsCode
var SendSmsCodeResponse graphql.Type

// DeviceChargeAbility 设备负责人权限类型
var DeviceChargeAbility graphql.Type

// DeviceCharge 设备负责人类型
var DeviceCharge graphql.Type

// DeviceParam 设备参数类型
var DeviceParam graphql.Type

// DeviceParamValue 设备参数值类型
var DeviceParamValue graphql.Type

// DeviceStatusLog 设备状态变更记录类型
var DeviceStatusLog graphql.Type

// Device 设备类型
var Device graphql.Type

// Privilege 权限类型
var Privilege graphql.Type

// RolePriv 角色权限关系类型
var RolePriv graphql.Type

// Role 角色类型
var Role graphql.Type

// UserExtend 用户信息拓展类型
var UserExtend graphql.Type

// UserLogin 用户登录类型
var UserLogin graphql.Type

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
				"deviceCharge": &graphql.Field{Type: DeviceCharge, Description: "设备负责人关系", Resolve: device.ChargeRelatedLoad},
				"privilege":    &graphql.Field{Type: Privilege, Description: "权限", Resolve: privilege.RelatedLoad},
			}
		}),
	})

	DeviceCharge = graphql.NewObject(graphql.ObjectConfig{
		Name: "DeviceCharge",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"device":    &graphql.Field{Type: Device, Description: "设备", Resolve: device.RelatedLoad},
				"user":      &graphql.Field{Type: User, Description: "负责人", Resolve: user.RelatedLoad},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
				"deviceChargeAbilities": &graphql.Field{
					Type:        graphql.NewList(DeviceChargeAbility),
					Description: "负责人权限",
					Resolve:     device.ChargeAbilityRelatedLoad,
				},
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
				"type":      &graphql.Field{Type: scalars.DeviceParamValueType, Description: "参数值类型"},
				"author":    &graphql.Field{Type: User, Description: "创建人", Resolve: user.RelatedLoad},
				"device":    &graphql.Field{Type: Device, Description: "设备", Resolve: device.RelatedLoad},
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
				"deviceParam": &graphql.Field{Type: DeviceParam, Description: "设备参数", Resolve: device.ParamRelatedLoad},
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
				"device":   &graphql.Field{Type: Device, Description: "设备", Resolve: device.RelatedLoad},
				"changeAt": &graphql.Field{Type: graphql.DateTime, Description: "变更时间"},
			}
		}),
	})

	Device = graphql.NewObject(graphql.ObjectConfig{
		Name: "Device",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"type":           &graphql.Field{Type: graphql.String, Description: "设备类型"},
				"name":           &graphql.Field{Type: graphql.String, Description: "设备名称"},
				"token":          &graphql.Field{Type: graphql.String, Description: "设备token，用于数据加密"},
				"status":         &graphql.Field{Type: scalars.DeviceStatus, Description: "基础状态"},
				"id":             &graphql.Field{Type: graphql.Int},
				"uuid":           &graphql.Field{Type: graphql.String, Description: "设备UUID"},
				"user":           &graphql.Field{Type: User, Description: "注册人用户", Resolve: user.RelatedLoad},
				"params":         &graphql.Field{Type: graphql.NewList(DeviceParam), Description: "设备参数", Resolve: device.ParamRelatedLoad},
				"deviceCharges":  &graphql.Field{Type: graphql.NewList(DeviceCharge), Description: "设备负责人", Resolve: device.ChargeRelatedLoad},
				"statusChangeAt": &graphql.Field{Type: graphql.DateTime, Description: "设备状态变更时间"},
				"description":    &graphql.Field{Type: graphql.String, Description: "设备描述，备注"},
				"createdAt":      &graphql.Field{Type: graphql.DateTime},
				"updatedAt":      &graphql.Field{Type: graphql.DateTime},
				"remoteIP":       &graphql.Field{Type: graphql.String},
			}
		}),
	})

	Privilege = graphql.NewObject(graphql.ObjectConfig{
		Name: "Privilege",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"name":     &graphql.Field{Type: graphql.String, Description: "权限名称"},
				"sign":     &graphql.Field{Type: graphql.String, Description: "权限签名"},
				"privType": &graphql.Field{Type: graphql.Int, Description: "权限类型"},
			}
		}),
	})

	RolePriv = graphql.NewObject(graphql.ObjectConfig{
		Name: "RolePriv",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"role":      &graphql.Field{Type: Role, Resolve: role.RelatedLoad},
				"privilege": &graphql.Field{Type: Privilege, Resolve: privilege.RelatedLoad},
			}
		}),
	})

	Role = graphql.NewObject(graphql.ObjectConfig{
		Name: "Role",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":       &graphql.Field{Type: graphql.Int},
				"roleName": &graphql.Field{Type: graphql.String, Description: "role name"},
				"status":   &graphql.Field{Type: scalars.BaseStatus, Description: "role status, can be default, publish, block and deleted"},
				"rolePrivs": &graphql.Field{
					Type:        graphql.NewList(RolePriv),
					Description: "role and privilege relationship",
					Resolve:     rolepriv.RelatedLoad,
				},
				"isAdmin":   &graphql.Field{Type: graphql.Boolean, Description: "是否为管理员角色，仅管理员角色才可以调用管理员API"},
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
				"user":  &graphql.Field{Type: User, Resolve: user.RelatedLoad},
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
				"role":       &graphql.Field{Type: Role, Description: "用户角色", Resolve: role.RelatedLoad},
				"userExtend": &graphql.Field{Type: UserExtend, Description: "用户拓展信息", Resolve: userextend.RelatedLoad},
				"status":     &graphql.Field{Type: scalars.BaseStatus, Description: "基础状态"},
				"createdAt":  &graphql.Field{Type: graphql.DateTime},
				"updatedAt":  &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	UserLogin = graphql.NewObject(graphql.ObjectConfig{
		Name: "UserLogin",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"userAgent": &graphql.Field{Type: graphql.String, Description: "UA"},
				"user":      &graphql.Field{Type: User, Description: "用户", Resolve: user.RelatedLoad},
				"remoteIP":  &graphql.Field{Type: graphql.String, Description: "头像链接"},
				"logout":    &graphql.Field{Type: graphql.Boolean, Description: "用户角色"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})
}
