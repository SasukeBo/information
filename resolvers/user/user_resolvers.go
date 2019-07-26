package user

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/resolvers"
	"github.com/SasukeBo/information/utils"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

// Create is a gql resolver, create user
func Create(params graphql.ResolveParams) (interface{}, error) {
	phoneStr := params.Args["phone"].(string)
	msgCodeStr := params.Args["smsCode"].(string)
	passwordStr := params.Args["password"].(string)

	rootValue := params.Info.RootValue.(map[string]interface{})

	_uuid := uuid.New().String()

	sessPhone := rootValue["phone"]
	sessMsgCode := rootValue["smsCode"]

	user := models.User{UUID: _uuid}

	// validate phone
	if err := resolvers.ValidatePhone(phoneStr); err != nil {
		return nil, err
	}
	if sessPhone == nil || sessMsgCode == nil || phoneStr != sessPhone || msgCodeStr != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, utils.ArgumentError{
			Field:   "smsCode",
			Message: "is not correct",
		}
	}
	user.Phone = phoneStr

	// validate password
	if err := resolvers.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}
	user.Password = utils.Encrypt(passwordStr)

	// 事务处理
	models.Repo.Begin()
	userExtend := models.UserExtend{}
	if _, err := models.Repo.Insert(&userExtend); err != nil {
		return nil, err
	}

	user.UserExtend = &userExtend
	if _, err := models.Repo.Insert(&user); err != nil {
		models.Repo.Rollback()

		return nil, err
	}
	// 事务提交
	models.Repo.Commit()

	return user, nil
}

// ResetPassword is a gql resolver, reset user password
func ResetPassword(params graphql.ResolveParams) (interface{}, error) {
	phoneStr := params.Args["phone"].(string)
	msgCodeStr := params.Args["smsCode"].(string)
	passwordStr := params.Args["password"].(string)

	rootValue := params.Info.RootValue.(map[string]interface{})

	sessPhone := rootValue["phone"]
	sessMsgCode := rootValue["smsCode"]

	// validate phone
	if err := resolvers.ValidatePhone(phoneStr); err != nil {
		return nil, err
	}
	if sessPhone == nil || sessMsgCode == nil || phoneStr != sessPhone || msgCodeStr != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, utils.ArgumentError{
			Field:   "smsCode",
			Message: "is not correct",
		}
	}
	if err := resolvers.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}

	user := models.User{Phone: phoneStr}
	if err := models.Repo.Read(&user, "phone"); err != nil {
		return nil, err
	}

	user.Password = utils.Encrypt(passwordStr)
	if _, err := models.Repo.Update(&user, "password"); err != nil {
		return nil, err
	}

	return user, nil
}
