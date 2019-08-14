package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/rolepriv"
	// "github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// RolePrivCreateField _
var RolePrivCreateField = &graphql.Field{
	Type: types.RolePriv,
	Args: graphql.FieldConfigArgument{
		"roleID": fields.GenArg(graphql.Int, "角色ID", false),
		"privID": fields.GenArg(graphql.Int, "权限ID", false),
	},
	Resolve: rolepriv.Create,
}

// RolePrivDeleteField _
var RolePrivDeleteField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"id": fields.GenArg(graphql.Int, "角色权限关系ID", false),
	},
	Resolve: rolepriv.Delete,
}
