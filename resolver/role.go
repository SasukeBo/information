package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
)

// LoadRole _
func LoadRole(params graphql.ResolveParams) (interface{}, error) {
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
		return nil, models.Error{Message: "load related role failed."}
	}
}
