package rolepriv

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// Create _
func Create(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_w"); err != nil {
		return nil, err
	}

	role := models.Role{ID: params.Args["roleID"].(int)}
	if err := role.GetBy("id"); err != nil {
		return nil, err
	}

	priv := models.Privilege{ID: params.Args["privID"].(int)}
	if err := priv.Get(); err != nil {
		return nil, err
	}

	rolePriv := models.RolePriv{Role: &role, Privilege: &priv}
	if err := rolePriv.Insert(); err != nil {
		return nil, err
	}

	return rolePriv, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_w"); err != nil {
		return nil, err
	}

	rolePriv := models.RolePriv{ID: params.Args["id"].(int)}
	if err := rolePriv.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}
