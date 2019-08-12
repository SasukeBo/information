package uextend

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "user_extend_w"); err != nil {
		return nil, err
	}

	userUUID := params.Args["userUUID"].(string)
	user := models.User{UUID: userUUID}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	var userExtend *models.UserExtend
	var err error
	if userExtend, err = user.LoadUserExtend(); err != nil {
		return nil, err
	}

	if name := params.Args["name"]; name != nil {
		userExtend.Name = name.(string)
	}

	if email := params.Args["email"]; email != nil {
		if err := utils.ValidateEmail(email.(string)); err != nil {
			return nil, err
		}
		userExtend.Email = email.(string)
	}

	if err := userExtend.Update("name", "email"); err != nil {
		return nil, err
	}

	return userExtend, nil
}
