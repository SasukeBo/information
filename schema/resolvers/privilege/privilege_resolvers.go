package privilege

import (
	"github.com/graphql-go/graphql"

	// "github.com/astaxie/beego/logs"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// RelatedLoad _
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	var id int

	switch v := params.Source.(type) {
	case models.DeviceChargeAbility:
		id = v.Privilege.ID
	case models.RolePriv:
		id = v.Privilege.ID
	default:
		return nil, utils.LogicError{
			Message: "reloated device_charge load error",
		}
	}

	privilege := models.Privilege{ID: id}
	if err := privilege.Get(); err != nil {
		return nil, err
	}

	return privilege, nil
}
