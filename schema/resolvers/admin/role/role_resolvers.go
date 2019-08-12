package role

import (
	// "fmt"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
)

// Create is a gql resolver, create role
func Create(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_w"); err != nil {
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

// Update is a gql resolver, update role
func Update(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_w"); err != nil {
		return nil, err
	}

	id := params.Args["id"].(int)
	role := models.Role{ID: id}
	if err := role.GetBy("id"); err != nil {
		return nil, err
	}

	if roleName := params.Args["roleName"]; roleName != nil {
		roleNameStr := roleName.(string)
		if err := utils.ValidateStringEmpty(roleNameStr, "roleName"); err != nil {
			return nil, err
		}
		role.RoleName = roleNameStr
	}

	if status := params.Args["status"]; status != nil {
		role.Status = status.(int)
	}

	if err := role.Update("role_name", "status"); err != nil {
		return nil, err
	}

	return role, nil
}

// Get is a gql resolver, get role by id
func Get(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_r"); err != nil {
		return nil, err
	}

	id := params.Args["id"].(int)
	role := &models.Role{ID: id}
	if err := role.GetBy("id"); err != nil {
		return nil, err
	}

	return role, nil
}

// GetByName is a gql resolver, get role by name
func GetByName(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_r"); err != nil {
		return nil, err
	}

	roleNameStr := params.Args["roleName"].(string)

	role := &models.Role{RoleName: roleNameStr}
	if err := role.GetBy("role_name"); err != nil {
		return nil, err
	}

	return role, nil
}

// List _
func List(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_r"); err != nil {
		return nil, err
	}

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
			Field:   "Role",
			Message: "List() error",
			OriErr:  err,
		}
	}

	return roles, nil
}

// Delete _
func Delete(params graphql.ResolveParams) (interface{}, error) {
	if err := utils.ValidateAccess(&params, "role_w"); err != nil {
		return nil, err
	}

	id := params.Args["id"].(int)
	role := models.Role{ID: id}
	if err := role.Delete(); err != nil {
		return nil, err
	}

	return "ok", nil
}
