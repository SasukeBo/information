package schemaadmin

import (
	resolver "github.com/SasukeBo/information/resolveradmin"
	"github.com/SasukeBo/information/schema"
	"github.com/graphql-go/graphql"
)

// RolePrivCreateField _
var RolePrivCreateField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"roleID": schema.GenArg(graphql.Int, "角色ID", false),
		"privID": schema.GenArg(graphql.Int, "权限ID", false),
	},
	Resolve: resolver.CreateRolePriv,
}

// RolePrivDeleteField _
var RolePrivDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": schema.GenArg(graphql.Int, "角色权限关系ID", false),
	},
	Resolve: resolver.DeleteRolePriv,
}
