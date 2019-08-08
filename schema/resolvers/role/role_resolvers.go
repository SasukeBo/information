package role

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	var value interface{}

	switch v := params.Source.(type) {
	case models.RolePriv:
		value = v.Role
	case models.User:
		value = v.Role
	default:
		return nil, utils.LogicError{
			Message: "reloated role load error",
		}
	}

	if value.(*models.Role) == nil {
		return nil, nil
	}

	role := value.(*models.Role)
	if err := role.Get(); err != nil {
		return nil, err
	}

	return role, nil
}
