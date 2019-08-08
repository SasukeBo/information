package role

import (
  "github.com/SasukeBo/information/models"
  "github.com/SasukeBo/information/schema/resolvers"
  "github.com/graphql-go/graphql"
)

// Create is a gql resolver, create role
func Create(params graphql.ResolveParams) (interface{}, error) {
  // TODO 权限验证
  roleNameStr := params.Args["roleName"].(string)
  if err := resolvers.ValidateStringEmpty(roleNameStr, "roleName"); err != nil {
    return nil, err
  }
  role := &models.Role{RoleName: roleNameStr}
  _, err := models.Repo.Insert(role)
  return role, err
}

// Update is a gql resolver, update role
func Update(params graphql.ResolveParams) (interface{}, error) {
  // TODO 权限验证
  id := params.Args["id"].(int)
  role := &models.Role{ID: id}
  err := models.Repo.Read(role)
  if err != nil {
    return role, err
  }

  if roleName := params.Args["roleName"]; roleName != nil {
    roleNameStr := roleName.(string)
    err := resolvers.ValidateStringEmpty(roleNameStr, "roleName")
    if err != nil {
      return role, err
    }
    role.RoleName = roleNameStr
  }

  if status := params.Args["status"]; status != nil {
    role.Status = status.(int)
  }

  _, err = models.Repo.Update(role)
  return role, err
}

// Get is a gql resolver, get role by id
func Get(params graphql.ResolveParams) (interface{}, error) {
  id := params.Args["id"].(int)
  role := &models.Role{ID: id}
  err := models.Repo.Read(role)
  return role, err
}

// GetByName is a gql resolver, get role by name
func GetByName(params graphql.ResolveParams) (interface{}, error) {
  roleNameStr := params.Args["roleName"].(string)
  if err := resolvers.ValidateStringEmpty(roleNameStr, "roleName"); err != nil {
    return nil, err
  }
  role := &models.Role{RoleName: roleNameStr}
  err := models.Repo.Read(role, "RoleName")
  return role, err
}
