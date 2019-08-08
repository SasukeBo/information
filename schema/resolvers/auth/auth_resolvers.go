package auth

import (
  "github.com/SasukeBo/information/models"
  "github.com/SasukeBo/information/utils"
  "github.com/astaxie/beego/orm"
  "github.com/graphql-go/graphql"
)

// LoginByPassword 使用账号密码方式登录
func LoginByPassword(params graphql.ResolveParams) (interface{}, error) {
  rootValue := params.Info.RootValue.(map[string]interface{})
  phoneStr := params.Args["phone"].(string)
  passwordStr := params.Args["password"].(string)
  remember := params.Args["remember"].(bool)

  user := models.User{Phone: phoneStr}
  if err := models.Repo.Read(&user, "phone"); err != nil {
    return nil, utils.LogicError{
      Message: "account not exist!",
    }
  }

  if user.Password != utils.Encrypt(passwordStr) {
    return nil, utils.LogicError{
      Message: "password not correct!",
    }
  }

  sessionID := rootValue["session_id"]
  if sessionID == nil {
    return nil, utils.LogicError{
      Message: "missing session_id",
    }
  }

  var userLogin models.UserLogin

  if err := models.Repo.QueryTable("user_login").Filter("user_id", user.ID).Filter("session_id", sessionID).One(&userLogin); err == orm.ErrNoRows {
    // 不存在 user_uuid 和 session_id 匹配的 user_login
    // 创建一个 user_login
    userLogin.User = &user
    userLogin.Remembered = remember
    userLogin.EncryptedPasswd = user.Password
    userLogin.SessionID = sessionID.(string)

    if remoteIP := rootValue["remote_ip"]; remoteIP != nil {
      userLogin.RemoteIP = remoteIP.(string)
    }

    if userAgent := rootValue["user_agent"]; userAgent != nil {
      userLogin.UserAgent = userAgent.(string)
    }

    if _, err := models.Repo.Insert(&userLogin); err != nil {
      return nil, err
    }
  } else if err == nil {
    // 存在 user_uuid 和 session_id 匹配的 user_login
    // 更新此 user_login logout 字段为 false
    userLogin.Logout = false
    models.Repo.Update(&userLogin)
  } else {
    return nil, err
  }

  rootValue["currentUserUUID"] = user.UUID
  rootValue["setSession"] = []string{"currentUserUUID"}
  rootValue["remember"] = remember

  return user.UUID, nil
}

// Logout 登出系统
func Logout(params graphql.ResolveParams) (interface{}, error) {
  rootValue := params.Info.RootValue.(map[string]interface{})
  currentUserUUID := rootValue["currentUserUUID"]
  sessionID := rootValue["session_id"]
  if currentUserUUID == nil {
    return nil, utils.LogicError{
      Message: "user not authenticated.",
    }
  }

  user := models.User{UUID: currentUserUUID.(string)}
  if err := models.Repo.Read(&user, "uuid"); err != nil {
    return nil, utils.LogicError{
      Message: "user not find.",
    }
  }

  var userLogin models.UserLogin
  err := models.Repo.QueryTable("user_login").Filter("user_id", user.ID).Filter("session_id", sessionID.(string)).One(&userLogin)
  if err == orm.ErrMultiRows {
    return nil, utils.LogicError{
      Message: "returned mutil rows not one.",
    }
  }
  if err == orm.ErrNoRows {
    return nil, utils.LogicError{
      Message: "user not authenticated.",
    }
  }
  userLogin.Logout = true
  models.Repo.Update(&userLogin, "logout")

  rootValue["currentUserUUID"] = nil
  rootValue["setSession"] = []string{"currentUserUUID"}

  return "ok", nil
}

// CurrentUser 获取当前用户
func CurrentUser(params graphql.ResolveParams) (interface{}, error) {
  rootValue := params.Info.RootValue.(map[string]interface{})
  currentUserUUID := rootValue["currentUserUUID"]
  if currentUserUUID == nil {
    return nil, utils.LogicError{
      Message: "user not authenticated.",
    }
  }

  user := models.User{UUID: currentUserUUID.(string)}
  if err := models.Repo.Read(&user, "uuid"); err != nil {
    return nil, utils.LogicError{
      Message: "user not found.",
    }
  }

  return user, nil
}
