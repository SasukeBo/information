package role

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.RolePriv:
		return v.LoadRole()
	case *models.RolePriv:
		return v.LoadRole()
	case models.User:
		return v.LoadRole()
	case *models.User:
		return v.LoadRole()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "load related source type unmatched error.",
		}
	}
}
