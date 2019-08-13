package ulogin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "user_login_r"); err != nil {
		return nil, err
	}

	qs := models.Repo.QueryTable("user_login").OrderBy("-created_at")

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	if userUUID := params.Args["userUUID"]; userUUID != nil {
		user := models.User{UUID: userUUID.(string)}
		if err := user.GetBy("uuid"); err != nil {
			return utils.EmptyResult, err
		}

		qs = qs.Filter("user_id", user.ID)
	}

	if remoteIP := params.Args["remoteIP"]; remoteIP != nil {
		qs = qs.Filter("remote_ip", remoteIP)
	}

	if beforeTime := params.Args["beforeTime"]; beforeTime != nil {
		qs = qs.Filter("created_at__lt", beforeTime)
	}

	if afterTime := params.Args["afterTime"]; afterTime != nil {
		qs = qs.Filter("created_at__gt", afterTime)
	}

	var userLogins []*models.UserLogin

	if _, err := qs.All(&userLogins); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "get user_login list error",
			OriErr:  err,
		}
	}

	return userLogins, nil
}
