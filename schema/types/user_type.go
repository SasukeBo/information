package types

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// UserType 用户类型
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"uuid":        &graphql.Field{Type: graphql.String},
		"account":     &graphql.Field{Type: graphql.String},
		"password":    &graphql.Field{Type: graphql.String},
		"userProfile": &graphql.Field{Type: UserProfileType},
		"role":        &graphql.Field{Type: RoleType},
		"status":      &graphql.Field{Type: custom.BaseStatus},
		"createdAt":   &graphql.Field{Type: graphql.DateTime},
		"updatedAt":   &graphql.Field{Type: graphql.DateTime},
	},
})

var (
	// account non null argument config
	accountNAC  = &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}
	passwordNAC = &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}
	roleIDAC    = &graphql.ArgumentConfig{Type: graphql.Int}
	realNameNAC = &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)}
)

// UserCreate create a user
var UserCreate = &graphql.Field{
	Type: UserType,
	Args: graphql.FieldConfigArgument{
		"account":  gNString,
		"password": gNString,
		"roleID":   gInt,
		"status":   gBaseStatus,
		"realName": gNString,
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_uuid := uuid.New().String()
		user := models.User{
			UUID:        _uuid,
			Account:     p.Args["account"].(string),
			Password:    p.Args["password"].(string),
			Role:        &models.Role{ID: p.Args["roleID"].(int)},
			UserProfile: &models.UserProfile{UUID: _uuid, RealName: p.Args["realName"].(string)},
			Status:      p.Args["status"].(models.BaseStatus),
		}

		err := user.Insert()
		return user, err
	},
}
