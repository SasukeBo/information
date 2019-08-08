package public

import (
  "github.com/graphql-go/graphql"

  "github.com/SasukeBo/information/schema/fields"
  "github.com/SasukeBo/information/schema/resolvers/auth"
  "github.com/SasukeBo/information/schema/types"
)

// CurrentUserField 获取当前登录的用户
var CurrentUserField = &graphql.Field{
  Type:    types.User,
  Resolve: auth.CurrentUser,
}

// LoginByPasswordField login by password
var LoginByPasswordField = &graphql.Field{
  Type: graphql.String,
  Args: graphql.FieldConfigArgument{
    "phone":    fields.GenArg(graphql.String, "手机号", false),
    "password": fields.GenArg(graphql.String, "密码", false),
    "remember": fields.GenArg(graphql.Boolean, "记住登录", true, true),
  },
  Resolve:     auth.LoginByPassword,
  Description: "请求时需要加上 operationName",
}

// LogoutField the system
var LogoutField = &graphql.Field{
  Type:        graphql.String,
  Resolve:     auth.Logout,
  Description: "请求时需要加上 operationName",
}
