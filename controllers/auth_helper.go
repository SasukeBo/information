package controllers

import (
	"github.com/SasukeBo/information/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
)

// authenticate 校验用户登录有效性
// 根据 gql 操作名称向 graphql.Params.RootObject 中放入值
// 至少会放入 currentUser 信息
// 验证失败则返回 error
// FIXME: 设计有漏洞，拿到sessionID可以冒充登录
func authenticate(ctx *context.Context) error {
	o := orm.NewOrm()
	// 如果当前 session 中有 currentUser
	if currentUser, ok := ctx.Input.Session("currentUser").(models.User); ok {
		if err := o.Read(&currentUser, "id"); err != nil {
			return models.Error{Message: "get user failed.", OriErr: err}
		}

		// 且用户当前状态正常
		if currentUser.Status == models.BaseStatus.Block {
			return models.Error{Message: "user has been blocked"}
		}

		// 则验证成功
		ctx.Output.Session("currentUser", currentUser)
		return nil
	}

	env := beego.AppConfig.String
	sessionID := ctx.Input.Cookie(env("SessionName"))

	// 通过请求头获取的浏览器保存的 session ID
	userLogin := models.UserLogin{SessionID: sessionID}

	// 获取 userLogin
	if err := o.Read(&userLogin, "session_id"); err != nil {
		// 查找userLogin失败，返回身份验证失败
		return models.Error{Message: "user not authenticated.", OriErr: err}
	}

	// 用户已经登出
	if userLogin.Logout {
		return models.Error{Message: "user already logout."}
	}

	// 用户没有记住登录
	if !userLogin.Remembered {
		return models.Error{Message: "user not keep login."}
	}

	var currentUser *models.User
	var err error

	if currentUser, err = userLogin.LoadUser(); err != nil {
		return err
	}
	// 判断密码是否匹配
	if currentUser.Password != userLogin.EncryptedPasswd {
		// 登录记录的密码与用户密码不匹配，验证失败
		return models.Error{Message: "password incorrect."}
	}

	// 当前session ID
	sessionID = ctx.Input.CruSession.SessionID()
	if userLogin.SessionID != sessionID {
		userLogin.Logout = true
		if _, err := o.Update(&userLogin, "logout"); err != nil {
			return models.Error{Message: "update user_login failed.", OriErr: err}
		}
		userLogin.ID = 0

		newUserLogin := models.UserLogin{
			EncryptedPasswd: userLogin.EncryptedPasswd,
			UserAgent:       userLogin.UserAgent,
			User:            userLogin.User,
			RemoteIP:        userLogin.RemoteIP,
			SessionID:       sessionID,
			Remembered:      userLogin.Remembered,
		}

		if _, err := o.Insert(&newUserLogin); err != nil {
			return models.Error{Message: "insert user_login failed.", OriErr: err}
		}
	}

	ctx.Output.Session("currentUser", *currentUser)
	expires, err := beego.AppConfig.Int("SessionCookieLifeTime")
	if err != nil {
		// 默认刷新为存活时间 7 天
		expires = 60 * 60 * 24 * 7
	}

	ctx.Output.Cookie(env("SessionName"), sessionID, expires)
	return nil
}

// AuthFilter 所有请求前的验证过滤器
func AuthFilter(ctx *context.Context) {
	if err := authenticate(ctx); err != nil {
		ctx.Output.Session("auth_error", err)
	}
}

// CleanAuthErrorFilter 执行完 controller 后清除每次请求的认证错误信息，防止造成下次请求误判当前 session 有认证 error
func CleanAuthErrorFilter(ctx *context.Context) {
	ctx.Output.Session("auth_error", nil)
}
