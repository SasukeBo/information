package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/uextend"
	// "github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// UserExtendUpdateField _
var UserExtendUpdateField = &graphql.Field{
	Type: types.UserExtend,
	Args: graphql.FieldConfigArgument{
		"userUUID": fields.GenArg(graphql.String, "用户UUID", false),
	},
	Resolve: uextend.Update,
}
