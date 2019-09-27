package resolver

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// ListUserLogin _
func ListUserLogin(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	qs := models.Repo.QueryTable("user_login").OrderBy("-created_at").Filter("user_id", user.ID)

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit.(int))
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset.(int))
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("created_at__lt", beforeTime.(time.Time))
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("created_at__gt", afterTime.(time.Time))
	}

	var userLogins []*models.UserLogin
	if _, err := qs.All(&userLogins); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get user_login list error",
			OriErr:  err,
		}
	}

	return userLogins, nil
}

// LastUserLogin _
func LastUserLogin(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})

	sessionID := rootValue["session_id"]
	user := rootValue["currentUser"].(models.User)

	qs := models.Repo.QueryTable("user_login").OrderBy("-created_at").Limit(1)
	cond := models.NewCond().And("user_id", user.ID).AndNot("session_id", sessionID)
	qs = qs.SetCond(cond)

	var lastLogin models.UserLogin
	if err := qs.One(&lastLogin); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get last user_login error",
			OriErr:  err,
		}
	}

	return lastLogin, nil
}

// ThisUserLogin _
func ThisUserLogin(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	sessionID := rootValue["session_id"].(string)

	userLogin := models.UserLogin{SessionID: sessionID}
	if err := userLogin.GetBy("session_id"); err != nil {
		return nil, err
	}

	return userLogin, nil
}
