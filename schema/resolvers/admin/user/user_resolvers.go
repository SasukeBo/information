package user

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_user_u", models.PrivType.Admin); err != nil {
		return nil, err
	}

	user := models.User{UUID: params.Args["uuid"].(string)}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	if phone := params.Args["phone"]; phone != nil {
		if err := utils.ValidatePhone(phone.(string)); err != nil {
			return nil, err
		}

		user.Phone = phone.(string)
	}

	if avatarURL := params.Args["avatarURL"]; avatarURL != nil {
		if err := utils.ValidateStringEmpty(avatarURL.(string), "avatarURL"); err != nil {
			return nil, err
		}

		user.AvatarURL = avatarURL.(string)
	}

	if roleID := params.Args["roleID"]; roleID != nil {
		user.Role = &models.Role{ID: roleID.(int)}
	}

	if status := params.Args["status"]; status != nil {
		user.Status = status.(int)
	}

	if err := user.Update("phone", "avatar_url", "role_id", "status"); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete 软删除
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_user_u", models.PrivType.Admin); err != nil {
		return nil, err
	}

	user := models.User{UUID: params.Args["uuid"].(string)}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	user.Status = models.BaseStatus.Block

	if err := user.Update("status"); err != nil {
		return nil, err
	}

	return "ok", nil
}

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_user_r", models.PrivType.Admin); err != nil {
		return nil, err
	}

	qs := models.Repo.QueryTable("user").OrderBy("-created_at")

	if limit := params.Args["limit"]; limit != nil {
		qs = qs.Limit(limit)
	}

	if offset := params.Args["offset"]; offset != nil {
		qs = qs.Offset(offset)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("user_extend__name__icontains", namePattern)
	}

	if phone := params.Args["phone"]; phone != nil {
		qs = qs.Filter("phone", phone)
	}

	if email := params.Args["email"]; email != nil {
		qs = qs.Filter("email", email)
	}

	if status := params.Args["status"]; status != nil && len(status.([]int)) > 0 {
		qs = qs.Filter("status__in", status)
	}

	if roleNamePattern := params.Args["roleNamePattern"]; roleNamePattern != nil {
		qs = qs.Filter("role__role_name__icontains", roleNamePattern)
	}

	if roleID := params.Args["roleID"]; roleID != nil {
		qs = qs.Filter("role_id", roleID)
	}

	var users []*models.User
	if _, err := qs.All(&users); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "admin user get list error",
			OriErr:  err,
		}
	}

	return users, nil
}
