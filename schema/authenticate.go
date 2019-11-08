package schema

import (
	"github.com/SasukeBo/information/resolver"
	"github.com/graphql-go/graphql"
)

/*						 query
---------------------------------- */

var currentUser = &graphql.Field{
	Type:        userType,
	Resolve:     resolver.CurrentUser,
	Description: `#### 获取当前登录用户`,
}

/*						mutation
---------------------------------- */

var signIn = &graphql.Field{
	Type: userType,
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

var signOut = &graphql.Field{
	Type:    graphql.String,
	Resolve: resolver.SignOut,
	Description: `
#### 退出登录

**注意** 需要提供operationName
	`,
}
