package auth

import (
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// LoginByPassword 使用账号密码方式登录
func LoginByPassword(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	phoneStr := params.Args["phone"].(string)
	passwordStr := params.Args["password"].(string)
	remember := params.Args["remember"].(bool)

	user := models.User{Phone: phoneStr}
	if err := user.GetBy("phone"); err != nil {
		return nil, err
	}

	if user.Password != utils.Encrypt(passwordStr) {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Auth",
			Message: "password incorrect.",
		}
	}

	sessionID := rootValue["session_id"]
	if sessionID == nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Auth",
			Message: "missing session_id.",
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

		if err := userLogin.Insert(); err != nil {
			return nil, err
		}
	} else if err == nil {
		// 存在 user_uuid 和 session_id 匹配的 user_login
		// 更新此 user_login logout 字段为 false
		userLogin.Logout = false
		if err := userLogin.Update("logout"); err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	rootValue["currentUser"] = user
	rootValue["setSession"] = []string{"currentUser"}
	rootValue["remember"] = remember

	return user.UUID, nil
}

// Logout 登出系统
func Logout(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	sessionID := rootValue["session_id"]
	currentUser := rootValue["currentUser"].(models.User)

	var userLogin models.UserLogin
	err := models.Repo.QueryTable("user_login").Filter("user_id", currentUser.ID).Filter("session_id", sessionID.(string)).One(&userLogin)
	if err == orm.ErrMultiRows {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Auth",
			Message: "Logout() error",
			OriErr:  err,
		}
	}

	if err == orm.ErrNoRows {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Auth",
			Message: "user not authenticated.",
		}
	}
	userLogin.Logout = true
	if err := userLogin.Update("logout"); err != nil {
		return nil, err
	}

	rootValue["currentUser"] = nil
	rootValue["setSession"] = []string{"currentUser"}

	return "ok", nil
}

// CurrentUser 获取当前用户
func CurrentUser(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	currentUser := rootValue["currentUser"].(models.User)

	return currentUser, nil
}
