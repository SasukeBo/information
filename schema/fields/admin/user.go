package admin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/admin/user"
	"github.com/SasukeBo/information/schema/scalars"
	"github.com/SasukeBo/information/schema/types"
)

// UserUpdateField update a user
var UserUpdateField = &graphql.Field{
	Type: types.User,
	Args: graphql.FieldConfigArgument{
		"uuid":       fields.GenArg(graphql.String, "用户UUID", false),
		"phone":      fields.GenArg(graphql.String, "用户手机号"),
		"avatar_url": fields.GenArg(graphql.String, "用户头像地址"),
		"role_id":    fields.GenArg(graphql.Int, "用户角色ID"),
		"status":     fields.GenArg(scalars.BaseStatus, "用户基础状态"),
	},
	Resolve: user.Update,
}
