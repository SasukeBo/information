package user

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "user_w"); err != nil {
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

	if err := user.Update("phone", "avatarURL", "roleID", "status"); err != nil {
		return nil, err
	}

	return user, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "user_w"); err != nil {
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
