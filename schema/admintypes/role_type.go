package admintypes

import (
	"github.com/SasukeBo/information/resolvers/admin/role"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/graphql-go/graphql"
)

// Role 用户类型
var Role graphql.Type

// RoleCreateType create a role
var RoleCreateType *graphql.Field

// RoleUpdateType create a role
var RoleUpdateType *graphql.Field

// RoleGetType get role by id
var RoleGetType *graphql.Field

// RoleGetByNameType get role by name
var RoleGetByNameType *graphql.Field

func init() {
	Role = graphql.NewObject(graphql.ObjectConfig{
		Name: "Role",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":        &graphql.Field{Type: graphql.Int},
				"roleName":  &graphql.Field{Type: graphql.String, Description: "role name"},
				"status":    &graphql.Field{Type: custom.BaseStatus, Description: "role status, can be default, publish, block and deleted"},
				"rolePrivs": &graphql.Field{Type: graphql.NewList(RolePriv), Description: "role and privilege relationship"},
				"createdAt": &graphql.Field{Type: graphql.DateTime},
				"updatedAt": &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	RoleCreateType = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"roleName": GenArg(graphql.String, "角色名称", false),
		},
		Resolve: role.Create,
	}

	RoleUpdateType = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"id":       GenArg(graphql.Int, "角色ID", false),
			"roleName": GenArg(graphql.String, "角色名称"),
			"status": GenArg(custom.BaseStatus, `
				基础状态
				- default 默认状态
				- publish 发布状态
				- block   屏蔽（禁用）状态
				- deleted 删除状态
			`),
		},
		Resolve: role.Update,
	}

	RoleGetType = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"id": GenArg(graphql.Int, "角色ID", false),
		},
		Resolve: role.Get,
	}

	RoleGetByNameType = &graphql.Field{
		Type: Role,
		Args: graphql.FieldConfigArgument{
			"roleName": GenArg(graphql.String, "角色名称", false),
		},
		Resolve: role.GetByName,
	}
}
