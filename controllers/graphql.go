package controllers

import (
	"encoding/json"
	"regexp"
	"strings"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
)

// GQLController is graphql controller
type GQLController struct {
	beego.Controller
}

// graphiqlData is the page data structure of the rendered GraphiQL page
type graphiqlData struct {
	GraphiqlVersion string
	QueryString     string
	VariablesString string
	OperationName   string
	ResultString    string
}

type gqlRootObject map[string]interface{}

// 匹配 graphql query oprtationName
var operationNameRegStr = `^query (\w+) {`

// Get http method
func (conn *GQLController) Get() {
	queryString := conn.Ctx.Input.Query("query")
	variablesString := conn.Ctx.Input.Query("variables")
	operationName := conn.Ctx.Input.Query("operationName")
	graphiqlVersion := "0.11.11"
	conn.TplName = "graphql.html"
	conn.Data["GraphiqlVersion"] = graphiqlVersion
	conn.Data["QueryString"] = queryString
	conn.Data["VariablesString"] = variablesString
	conn.Data["OperationName"] = operationName
}

// Post http method
func (conn *GQLController) Post() {
	var params struct {
		Query         string                 `json:"query"`
		Variables     map[string]interface{} `json:"variables"`
		OperationName string                 `json:"operationName"`
	}

	json.NewDecoder(conn.Ctx.Request.Body).Decode(&params)
	logs.Info(params.Query)

	if params.OperationName == "" {
		// 解决 query 中包含了 operationName 但是请求体 JSON 缺失 operationName 选项。
		reg := regexp.MustCompile(operationNameRegStr)
		if matches := reg.FindStringSubmatch(strings.TrimSpace(params.Query)); len(matches) > 1 {
			params.OperationName = matches[1]
		}
	}

	rootObject := gqlRootObject{}
	var result *graphql.Result
	if err := authenticate(conn, rootObject, params.OperationName); err != nil {
		// 返回错误信息
		result = &graphql.Result{
			Errors: []gqlerrors.FormattedError{
				gqlerrors.FormattedError{
					Message: err.Error(),
				},
			},
		}
	} else {
		gqlParams := graphql.Params{
			Schema:         schema.Schema,
			RequestString:  params.Query,
			OperationName:  params.OperationName,
			VariableValues: params.Variables,
			RootObject:     rootObject,
		}

		result = graphql.Do(gqlParams)
		setSession(conn, gqlParams.RootObject)
	}
	conn.Data["json"] = result
	conn.ServeJSON()
}

// authenticate 校验用户登录有效性
// 根据 gql 操作名称向 graphql.Params.RootObject 中放入值
// 至少会放入 currentUser 信息
// 验证失败则返回 error
func authenticate(conn *GQLController, obj gqlRootObject, name string) error {
	env := beego.AppConfig.String
	// 刷新 cookie
	currentSessionID := conn.Ctx.Input.CruSession.SessionID()

	switch name {
	case "IntrospectionQuery":
		// graphiql schema query
		return nil
	case "register":
		obj["phone"] = conn.GetSession("phone")
		obj["smsCode"] = conn.GetSession("smsCode")
		return nil
	case "loginByPassword":
		// 登录操作不需要后面的验证
		// 需要记录 IP UA
		obj["remote_ip"] = conn.Ctx.Input.IP()
		obj["user_agent"] = conn.Ctx.Input.UserAgent()
		obj["session_id"] = currentSessionID
		return nil
	case "logout":
		obj["session_id"] = currentSessionID
	}

	sessionID := conn.Ctx.Input.Cookie(env("SessionName"))
	userLogin := models.UserLogin{SessionID: sessionID}
	if err := models.Repo.Read(&userLogin, "session_id"); err != nil {
		// 查找userLogin失败，返回身份验证失败
		return utils.LogicError{
			Message: "user login not find.",
		}
	}

	if currentUserUUID := conn.GetSession("currentUserUUID"); currentUserUUID == nil {
		// 如果没有currentUserUUID
		// 通过sessionID获取userLogin
		if userLogin.Logout {
			return utils.LogicError{
				Message: "user already logout.",
			}
		}
		if !userLogin.Remembered {
			return utils.LogicError{
				Message: "user login not remembered.",
			}
		}
		user := models.User{UUID: userLogin.UserUUID}
		if err := models.Repo.Read(&user, "uuid"); err != nil {
			// 查找user失败后，返回身份验证失败
			return utils.LogicError{
				Message: "user not find.",
			}
		}
		if user.Password != userLogin.EncryptedPasswd {
			// 登录记录的密码与用户密码不匹配，验证失败
			return utils.LogicError{
				Message: "password unmatch, maybe changed.",
			}
		}
		userLogin.SessionID = currentSessionID

		conn.SetSession("currentUserUUID", user.UUID)
		obj["currentUserUUID"] = user.UUID
		models.Repo.Update(&userLogin)
	} else {
		obj["currentUserUUID"] = currentUserUUID
	}

	if userLogin.Remembered {
		expires, err := beego.AppConfig.Int("SessionCookieLifeTime")
		if err != nil {
			// 默认刷新为存活时间 7 天
			expires = 60 * 60 * 24 * 7
		}

		conn.Ctx.Output.Cookie(env("SessionName"), currentSessionID, expires)
	}

	return nil
}

// setSession 根据 graphql.Params.RootObject 中 setSession 对应的
// string 数组内容，从 RootObject 中取值并放入 Session
func setSession(conn *GQLController, obj gqlRootObject) {
	if fields := obj["setSession"]; fields != nil {
		for _, field := range fields.([]string) {
			conn.SetSession(field, obj[field])
		}
	}
	if remember := obj["remember"]; remember != nil {
		if !remember.(bool) {
			currentSessionID := conn.Ctx.Input.CruSession.SessionID()
			conn.Ctx.Output.Cookie(beego.AppConfig.String("SessionName"), currentSessionID, 0)
		}
	}
}
