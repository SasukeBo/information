package uextend

import (
	"regexp"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/resolvers"
	"github.com/SasukeBo/information/utils"
)

var emailRegexp = `[^\\.\\s@:](?:[^\\s@:]*[^\\s@:\\.])?@[^\\.\\s@]+(?:\\.[^\\.\\s@]+)*`

// Update _
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := resolvers.ValidateAccess(&params, "user_extend_w"); err != nil {
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
		reg := regexp.MustCompile(emailRegexp)
		if !reg.Match([]byte(email.(string))) {
			return nil, utils.LogicError{
				Message: "invalid email",
			}
		}
		userExtend.Email = email.(string)
	}

	if err := userExtend.Update("name", "email"); err != nil {
		return nil, err
	}

	return userExtend, nil
}
