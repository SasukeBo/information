package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/ulogin"
	// "github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// UserLoginListField _
var UserLoginListField = &graphql.Field{
	Type: graphql.NewList(types.UserLogin),
	Args: graphql.FieldConfigArgument{
		"todo": fields.GenArg(graphql.String, "todo", false),
	},
	Resolve: ulogin.List,
}
