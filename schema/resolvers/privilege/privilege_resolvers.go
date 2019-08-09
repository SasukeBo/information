package privilege

import (
	"github.com/graphql-go/graphql"

	// "github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// List is a gql resolver, get list of privilege
func List(params graphql.ResolveParams) (interface{}, error) {
	// TODO:
	return nil, nil
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
		return nil, utils.LogicError{
			Message: "related privilege load error",
		}
	}
}
