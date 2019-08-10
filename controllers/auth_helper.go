package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// authenticate 校验用户登录有效性
// 根据 gql 操作名称向 graphql.Params.RootObject 中放入值
// 至少会放入 currentUser 信息
// 验证失败则返回 error
func authenticate(ctx *context.Context) error {
	currentUserUUID := ctx.Input.Session("currentUserUUID")

	// 如果当前 session 中有 currentUserUUID，则验证成功
	if currentUserUUID != nil {
		user := models.User{UUID: currentUserUUID.(string)}
		if err := user.GetBy("uuid"); err != nil {
			return err
		}

		if user.Status == models.BaseStatus.Block {
			return utils.LogicError{
				Message: "Account has been blocked",
			}
		}

		return nil
	}

	env := beego.AppConfig.String
	sessionID := ctx.Input.Cookie(env("SessionName"))

	// 通过请求头获取的浏览器保存的 session ID
	userLogin := models.UserLogin{SessionID: sessionID}

	// 获取 userLogin
	if err := models.Repo.Read(&userLogin, "session_id"); err != nil {
		// 查找userLogin失败，返回身份验证失败
		return utils.LogicError{
			Message: "user not authenticated.",
		}
	}

	// 用户已经登出
	if userLogin.Logout {
		return utils.LogicError{
			Message: "user already logout.",
		}
	}

	// 用户没有记住登录
	if !userLogin.Remembered {
		return utils.LogicError{
			Message: "user login not remembered.",
		}
	}

	var user *models.User
	var err error

	if user, err = userLogin.LoadUser(); err != nil {
		return err
	}
	// 判断密码是否匹配
	if user.Password != userLogin.EncryptedPasswd {
		// 登录记录的密码与用户密码不匹配，验证失败
		return utils.LogicError{
			Message: "password unmatch, maybe changed.",
		}
	}

	// 当前session ID
	sessionID = ctx.Input.CruSession.SessionID()

	userLogin.SessionID = sessionID

	ctx.Output.Session("currentUserUUID", user.UUID)
	models.Repo.Update(&userLogin)

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
