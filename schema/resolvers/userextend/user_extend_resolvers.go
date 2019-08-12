package userextend

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	currentUserUUID := params.Info.RootValue.(map[interface{}]interface{})["currentUserUUID"].(string)
	user := models.User{UUID: currentUserUUID}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	userExtend := user.UserExtend
	name := params.Args["name"]

	if name != nil {
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
			Field:   "UserExtend",
			Message: "RelatedLoad() error",
		}
	}
}
