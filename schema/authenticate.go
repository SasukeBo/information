package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*						query fields
------------------------------------------ */

var currentUserField = &graphql.Field{
	Type:        User,
	Resolve:     resolver.CurrentUser,
	Description: `#### 获取当前登录用户`,
}

/*						mutation fields
------------------------------------------ */

var signInField = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"phone":    GenArg(graphql.String, "手机号", false),
		"password": GenArg(graphql.String, "密码", false),
		"remember": GenArg(graphql.Boolean, "记住登录", true, true),
	},
	Resolve: resolver.SignIn,
	Description: `
#### 使用账号密码登录

**注意** 需要提供operationName
	`,
}

var signOutField = &graphql.Field{
	Type:    graphql.String,
	Resolve: resolver.SignOut,
	Description: `
#### 退出登录

**注意** 需要提供operationName
	`,
}
