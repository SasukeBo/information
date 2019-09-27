package resolveradmin

import (
	// "fmt"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
)

// CreateRole _
func CreateRole(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_role_c", models.PrivType.Admin); err != nil {
		return nil, err
	}

	roleNameStr := params.Args["roleName"].(string)
	if err := utils.ValidateStringEmpty(roleNameStr, "roleName"); err != nil {
		return nil, err
	}

	role := models.Role{RoleName: roleNameStr}
	if err := role.Insert(); err != nil {
		return nil, err
	}

	return role, nil
}

// UpdateRole _
func UpdateRole(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_role_u", models.PrivType.Admin); err != nil {
		return nil, err
	}

	id := params.Args["id"].(int)
	role := models.Role{ID: id}
	if err := role.GetBy("id"); err != nil {
		return nil, err
	}

	if roleName := params.Args["roleName"]; roleName != nil {
		if err := utils.ValidateStringEmpty(roleName.(string), "roleName"); err != nil {
			return nil, err
		}
		role.RoleName = roleName.(string)
	}

	if status := params.Args["status"]; status != nil {
		role.Status = status.(int)
	}

	if err := role.Update("role_name", "status"); err != nil {
		return nil, err
	}

	return role, nil
}

// GetRole _
func GetRole(params graphql.ResolveParams) (interface{}, error) {
	// if err := utils.ValidateAccess(&params, "role_r"); err != nil {
	// return nil, err
	// }

	id := params.Args["id"].(int)
	role := &models.Role{ID: id}
	if err := role.GetBy("id"); err != nil {
		return nil, err
	}

	return role, nil
}

// GetRoleByName _
func GetRoleByName(params graphql.ResolveParams) (interface{}, error) {
	// if err := utils.ValidateAccess(&params, "role_r"); err != nil {
	// return nil, err
	// }

	roleNameStr := params.Args["roleName"].(string)

	role := &models.Role{RoleName: roleNameStr}
	if err := role.GetBy("role_name"); err != nil {
		return nil, err
	}

	return role, nil
}

// ListRole _
func ListRole(params graphql.ResolveParams) (interface{}, error) {
	// if err := utils.ValidateAccess(&params, "role_r"); err != nil {
	// return nil, err
	// }

	roleNamePattern := params.Args["roleNamePattern"]
	status := params.Args["status"]

	qs := models.Repo.QueryTable("role")

	if roleNamePattern != nil {
		qs = qs.Filter("role_name__icontains", roleNamePattern)
	}

	if status != nil && len(status.([]interface{})) > 0 {
		qs = qs.Filter("status__in", status)
	}

	var roles []*models.Role
	if _, err := qs.All(&roles); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Message: "get role list error",
			OriErr:  err,
		}
	}

	return roles, nil
}

// DeleteRole _
func DeleteRole(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "admin_role_d", models.PrivType.Admin); err != nil {
		return nil, err
	}

	id := params.Args["id"].(int)
	role := models.Role{ID: id}
	if err := role.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}
