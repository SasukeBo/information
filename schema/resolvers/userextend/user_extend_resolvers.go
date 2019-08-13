package userextend

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[interface{}]interface{})["currentUser"].(models.User)
	user.LoadUserExtend()
	userExtend := user.UserExtend

	if name := params.Args["name"]; name != nil {
		if err := utils.ValidateStringEmpty(name.(string), "name"); err != nil {
			return nil, err
		}
		userExtend.Name = name.(string)
	}

	if err := userExtend.Update("name"); err != nil {
		return nil, err
	}

	return userExtend, nil
}

// BindEmail _
func BindEmail(params graphql.ResolveParams) (interface{}, error) {
	return "绑定邮箱接口暂时保留", nil
}

// RelatedLoad load user_extend
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.User:
		return v.LoadUserExtend()
	case *models.User:
		return v.LoadUserExtend()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
