package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/privilege"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// PrivilegeListField get list of privilege
var PrivilegeListField = &graphql.Field{
	Type: graphql.NewList(types.Privilege),
	Args: graphql.FieldConfigArgument{
		"privType":    fields.GenArg(scalars.PrivType, "权限类型"),
		"namePattern": fields.GenArg(graphql.String, "权限名称模糊匹配"),
	},
	Resolve: privilege.List,
}
