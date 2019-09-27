package resolveradmin

import (
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// CreateRolePriv _
func CreateRolePriv(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_role_priv_w", models.PrivType.Admin); err != nil {
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

	return "ok", nil
}

// DeleteRolePriv _
func DeleteRolePriv(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_role_priv_w", models.PrivType.Admin); err != nil {
		return nil, err
	}

	rolePriv := models.RolePriv{ID: params.Args["id"].(int)}
	if err := rolePriv.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}
