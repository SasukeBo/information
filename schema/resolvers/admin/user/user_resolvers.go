package user

import (
  "github.com/SasukeBo/information/models"
  "github.com/graphql-go/graphql"
)

// Update doc false
func Update(params graphql.ResolveParams) (interface{}, error) {
  // TODO: valiadate user_w ability
  uuid := params.Args["uuid"].(string)

  user := models.User{UUID: uuid}
  if err := models.Repo.Read(&user, "uuid"); err != nil {
    return nil, err
  }

  phone := params.Args["phone"]
  avatarURL := params.Args["avatar_url"]
  roleID := params.Args["role_id"]
  status := params.Args["status"]

  if phone != nil {
    user.Phone = phone.(string)
  }

  if avatarURL != nil {
    user.AvatarURL = avatarURL.(string)
  }

  if roleID != nil {
    user.Role = &models.Role{ID: roleID.(int)}
  }

  if status != nil {
    user.Status = status.(int)
  }

  if _, err := models.Repo.Update(&user); err != nil {
    return nil, err
  }

  return user, nil
}
