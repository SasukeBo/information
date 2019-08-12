package privilege

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
)

// List is a gql resolver, get list of privilege
func List(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("privilege")

	if privType := params.Args["privType"]; privType != nil {
		qs = qs.Filter("priv_type", privType.(int))
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern.(string))
	}

	var privs []*models.Privilege
	if _, err := qs.All(&privs); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Privilege",
			Message: "List() error",
			OriErr:  err,
		}
	}

	return privs, nil
}

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceChargeAbility:
		return v.LoadPrivilege()
	case *models.DeviceChargeAbility:
		return v.LoadPrivilege()
	case models.RolePriv:
		return v.LoadPrivilege()
	case *models.RolePriv:
		return v.LoadPrivilege()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "Privilege",
			Message: "RelatedLoad() error",
		}
	}
}
