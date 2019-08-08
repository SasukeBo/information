package admin

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/admin/role"
  "github.com/SasukeBo/information/schema/scalars"
  "github.com/SasukeBo/information/schema/types"
)

// RoleCreateField create a role
var RoleCreateField = &graphql.Field{
  Type: types.Role,
  Args: graphql.FieldConfigArgument{
    "roleName": fields.GenArg(graphql.String, "角色名称", false),
  },
  Resolve: role.Create,
}

// RoleUpdateField create a role
var RoleUpdateField = &graphql.Field{
  Type: types.Role,
  Args: graphql.FieldConfigArgument{
    "id":       fields.GenArg(graphql.Int, "角色ID", false),
    "roleName": fields.GenArg(graphql.String, "角色名称"),
    "status": fields.GenArg(scalars.BaseStatus, `
        基础状态
        - default 默认状态
        - publish 发布状态
        - block   屏蔽（禁用）状态
        - deleted 删除状态
      `),
  },
  Resolve: role.Update,
}

// RoleGetField get role by id
var RoleGetField = &graphql.Field{
  Type: types.Role,
  Args: graphql.FieldConfigArgument{
    "id": fields.GenArg(graphql.Int, "角色ID", false),
  },
  Resolve: role.Get,
}

// RoleGetByNameField get role by name
var RoleGetByNameField = &graphql.Field{
  Type: types.Role,
  Args: graphql.FieldConfigArgument{
    "roleName": fields.GenArg(graphql.String, "角色名称", false),
  },
  Resolve: role.GetByName,
}