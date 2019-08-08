package userextend

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
)

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	currentUserUUID := params.Info.RootValue.(map[interface{}]interface{})["currentUserUUID"].(string)
	user := models.User{UUID: currentUserUUID}
	if err := user.GetByUUID(); err != nil {
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
