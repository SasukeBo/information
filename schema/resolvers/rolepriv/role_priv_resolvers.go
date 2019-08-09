package rolepriv

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Role:
		return v.LoadRolePriv()
	case *models.Role:
		return v.LoadRolePriv()
	default:
		return nil, utils.LogicError{
			Message: "reloated role_priv load error",
		}
	}
}
