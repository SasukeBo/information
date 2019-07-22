package types

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/custom"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// User 用户类型
var User graphql.Type

// UserCreate create a user
var UserCreate *graphql.Field

func init() {
	User = graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.FieldsThunk(func() graphql.Fields {
			return graphql.Fields{
				"id":         &graphql.Field{Type: graphql.Int},
				"uuid":       &graphql.Field{Type: graphql.String, Description: "通用唯一标识"},
				"phone":      &graphql.Field{Type: graphql.String, Description: "手机号"},
				"password":   &graphql.Field{Type: graphql.String},
				"role":       &graphql.Field{Type: Role, Description: "用户角色"},
				"userExtend": &graphql.Field{Type: UserExtend, Description: "用户拓展信息"},
				"status":     &graphql.Field{Type: custom.BaseStatus, Description: "基础状态"},
				"createdAt":  &graphql.Field{Type: graphql.DateTime},
				"updatedAt":  &graphql.Field{Type: graphql.DateTime},
			}
		}),
	})

	UserCreate = &graphql.Field{
		Type: User,
		Args: graphql.FieldConfigArgument{
			"phone":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), Description: "手机号"},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), Description: "密码"},
			"msgCode":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String), Description: "手机短信验证码"},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			_uuid := uuid.New().String()
			// TODO validate msgCode by session
			// msgCode := p.Args["msgCode"].(string)
			user := models.User{
				UUID:     _uuid,
				Phone:    p.Args["phone"].(string),
				Password: p.Args["password"].(string),
			}

			err := user.Insert()
			return user, err
		},
	}
}
