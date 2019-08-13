package rolepriv

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/models"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Role:
		return v.LoadRolePriv()
	case *models.Role:
		return v.LoadRolePriv()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "RolePriv",
			Message: "RelatedLoad() error",
		}
	}
}
