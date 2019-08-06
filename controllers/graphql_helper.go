package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/logs"
)

/*
	gqlGetSession 获取 session 中的数据到 gqlRootObject 中。
	conn - controller 对象
	obj - gqlRootObject graphql 根对象
	opName - string graphq 操作名
*/
func gqlGetSession(conn *GQLController, obj gqlRootObject, opName string) {
	switch opName {
	case "IntrospectionQuery", "sendSmsCode":
		// graphiql schema query
		break

	case "register", "resetPassword", "getSmsCode":
		obj["phone"] = conn.GetSession("phone")
		obj["smsCode"] = conn.GetSession("smsCode")

	case "loginByPassword":
		// 登录操作不需要后面的验证
		// 需要记录 IP UA
		obj["remote_ip"] = conn.Ctx.Input.IP()
		obj["user_agent"] = conn.Ctx.Input.UserAgent()
		fallthrough

	case "logout":
		obj["session_id"] = conn.Ctx.Input.CruSession.SessionID()
	}

	obj["currentUserUUID"] = conn.GetSession("currentUserUUID")
}

// gqlSetSession 根据 graphql.Params.RootObject 中 set session 对应的
// string 数组内容，从 RootObject 中取值并放入 Session
func gqlSetSession(conn *GQLController, obj gqlRootObject) {
	if fields := obj["setSession"]; fields != nil {
		for _, field := range fields.([]string) {
			conn.SetSession(field, obj[field])
		}
	}
	if remember := obj["remember"]; remember != nil {
		if !remember.(bool) {
			// 如果用户登录为不记住登录，则将cookie过期时间设置为 回话
			currentSessionID := conn.Ctx.Input.CruSession.SessionID()
			conn.Ctx.Output.Cookie(beego.AppConfig.String("SessionName"), currentSessionID, 0)
		}
	}
}
