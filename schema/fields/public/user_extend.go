package public

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/schema/fields"
	"github.com/SasukeBo/information/schema/resolvers/userextend"
	"github.com/SasukeBo/information/schema/types"
)

// UserExtendUpdateField _
var UserExtendUpdateField = &graphql.Field{
	Type: types.UserExtend,
	Args: graphql.FieldConfigArgument{
		"name": fields.GenArg(graphql.String, "姓名"),
	},
	Resolve:     userextend.Update,
	Description: "更新用户附加信息",
}

// UserExtendBindEmailField _
var UserExtendBindEmailField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"email": fields.GenArg(graphql.String, "姓名"),
	},
	Resolve:     userextend.BindEmail,
	Description: "用户绑定邮箱",
}
