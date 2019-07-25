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

	userLogin := models.UserLogin{}
	userLogin.UserUUID = user.UUID
	userLogin.Remembered = remember
	userLogin.EncryptedPasswd = user.Password

	if sessionID := rootValue["session_id"]; sessionID != nil {
		userLogin.SessionID = sessionID.(string)
	}

	if remoteIP := rootValue["remote_ip"]; remoteIP != nil {
		userLogin.RemoteIP = remoteIP.(string)
	}

	if userAgent := rootValue["user_agent"]; userAgent != nil {
		userLogin.UserAgent = userAgent.(string)
	}

	if _, err := models.Repo.Insert(&userLogin); err != nil {
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
	sessionID := rootValue["currentUserUUID"]
	if currentUserUUID == nil {
		return nil, utils.LogicError{
			Message: "user not authenticated.",
		}
	}

	var userLogin models.UserLogin
	err := models.Repo.QueryTable("user_login").Filter("user_uuid", currentUserUUID.(string)).Filter("session_id", sessionID.(string)).One(&userLogin)
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

	rootValue["currentUserUUID"] = nil
	rootValue["setSession"] = []string{"currentUserUUID"}

	return "ok", nil
}
