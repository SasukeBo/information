package userlogin

import (
	"time"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	currentUserUUID := params.Info.RootValue.(map[string]interface{})["currentUserUUID"].(string)

	user := models.User{UUID: currentUserUUID}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

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
		return nil, utils.ORMError{
			Message: "user_login get list error",
			OrmErr:  err,
		}
	}

	return userLogins, nil
}
