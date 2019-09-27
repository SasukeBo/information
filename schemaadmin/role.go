package schemaadmin

import (
	resolver "github.com/SasukeBo/information/resolveradmin"
	"github.com/SasukeBo/information/schema"
	"github.com/graphql-go/graphql"
)

// RoleCreateField create a role
var RoleCreateField = &graphql.Field{
	Type: schema.Role,
	Args: graphql.FieldConfigArgument{
		"roleName": schema.GenArg(graphql.String, "角色名称", false),
	},
	Resolve: resolver.CreateRole,
}

// RoleUpdateField create a role
var RoleUpdateField = &graphql.Field{
	Type: schema.Role,
	Args: graphql.FieldConfigArgument{
		"id":       schema.GenArg(graphql.Int, "角色ID", false),
		"roleName": schema.GenArg(graphql.String, "角色名称"),
		"status": schema.GenArg(schema.BaseStatus, `
        基础状态
        - default 默认状态
        - publish 发布状态
        - block   屏蔽（禁用）状态
        - deleted 删除状态
      `),
	},
	Resolve: resolver.UpdateRole,
}

// RoleDeleteField delete a role
var RoleDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": schema.GenArg(graphql.Int, "角色ID", false),
	},
	Resolve: resolver.DeleteRole,
}

// RoleGetField get role by id
var RoleGetField = &graphql.Field{
	Type: schema.Role,
	Args: graphql.FieldConfigArgument{
		"id": schema.GenArg(graphql.Int, "角色ID", false),
	},
	Resolve: resolver.GetRole,
}

// RoleGetByNameField get role by name
var RoleGetByNameField = &graphql.Field{
	Type: schema.Role,
	Args: graphql.FieldConfigArgument{
		"roleName": schema.GenArg(graphql.String, "角色名称", false),
	},
	Resolve: resolver.GetRoleByName,
}

// RoleListField get role by name
var RoleListField = &graphql.Field{
	Type: graphql.NewList(schema.Role),
	Args: graphql.FieldConfigArgument{
		"roleNamePattern": schema.GenArg(graphql.String, "角色名称模糊匹配"),
		"status":          schema.GenArg(graphql.NewList(schema.BaseStatus), "状态，列表"),
	},
	Resolve: resolver.ListRole,
}
