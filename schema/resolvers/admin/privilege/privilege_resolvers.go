package privilege

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
)

// List get list of privilege
func List(params graphql.ResolveParams) (interface{}, error) {
	qs := models.Repo.QueryTable("privilege")

	if privType := params.Args["privType"]; privType != nil {
		qs = qs.Filter("priv_type", privType)
	}

	if namePattern := params.Args["namePattern"]; namePattern != nil {
		qs = qs.Filter("name__icontains", namePattern)
	}

	var privs []*models.Privilege
	if _, err := qs.All(&privs); err != nil {
		return nil, errors.LogicError{
			Type:    "Model",
			Message: "get privilege list error",
			OriErr:  err,
		}
	}

	return privs, nil
}