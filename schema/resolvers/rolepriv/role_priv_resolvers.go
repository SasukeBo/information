package rolepriv

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	var result interface{}

	switch v := params.Source.(type) {
	case models.Role:
		if _, err := models.Repo.LoadRelated(&v, "RolePriv"); err != nil {
			return nil, utils.ORMError{
				Message: "related load role_priv error",
				OrmErr:  err,
			}
		}
		result = v.RolePriv

	default:
		return nil, utils.LogicError{
			Message: "reloated role load error",
		}
	}

	return result, nil
}
