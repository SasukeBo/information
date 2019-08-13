package user

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/utils"
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
	if err := utils.ValidatePhone(phoneStr); err != nil {
		return nil, err
	}
	if sessPhone == nil || sessMsgCode == nil || phoneStr != sessPhone || msgCodeStr != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "smsCode",
			Message: "is not correct",
		}
	}
	user.Phone = phoneStr

	// validate password
	if err := utils.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}
	user.Password = utils.Encrypt(passwordStr)

	// 事务处理
	models.Repo.Begin()
	userExtend := models.UserExtend{}
	if err := userExtend.Insert(); err != nil {
		return nil, err
	}
	user.UserExtend = &userExtend

	role := models.Role{RoleName: "default"}
	if err := role.GetBy("role_name"); err != nil {
		return nil, err
	}
	user.Role = &role

	if err := user.Insert(); err != nil {
		models.Repo.Rollback()

		return nil, err
	}
	// 事务提交
	models.Repo.Commit()

	rootValue["smsCode"] = nil
	rootValue["setSession"] = []string{"smsCode"}

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
	if err := utils.ValidatePhone(phoneStr); err != nil {
		return nil, err
	}
	if sessPhone == nil || sessMsgCode == nil || phoneStr != sessPhone || msgCodeStr != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "smsCode",
			Message: "is not correct",
		}
	}
	if err := utils.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}

	user := models.User{Phone: phoneStr}
	if err := user.GetBy("phone"); err != nil {
		return nil, err
	}

	user.Password = utils.Encrypt(passwordStr)
	if err := user.Update("password"); err != nil {
		return nil, err
	}

	rootValue["currentUser"] = nil
	rootValue["smsCode"] = nil
	rootValue["setSession"] = []string{"currentUser", "smsCode"}

	return user, nil
}

// Get get user by uuid
func Get(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	user := models.User{UUID: uuid}
	if err := user.GetBy("uuid"); err != nil {
		return nil, err
	}

	return user, nil
}

// List get list of user
func List(params graphql.ResolveParams) (interface{}, error) {
	namePattern := params.Args["namePattern"]
	phone := params.Args["phone"]
	email := params.Args["email"]

	var users []*models.User

	if namePattern == nil && phone == nil && email == nil {
		return users, nil
	}

	qs := models.Repo.QueryTable("user")

	if namePattern != nil {
		qs = qs.Filter("user_extend__name__icontains", namePattern.(string))
	}

	if phone != nil {
		qs = qs.Filter("phone", phone.(string))
	}

	if email != nil {
		qs = qs.Filter("user_extend__email", email.(string))
	}

	if _, err := qs.All(&users); err != nil {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "User",
			Message: "List() error",
			OriErr:  err,
		}
	}

	return users, nil
}

// UpdateAvatar update user avatar url
func UpdateAvatar(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	avatarURL := params.Args["avatarURL"].(string)

	if err := utils.ValidateStringEmpty(avatarURL, "avatarURL"); err != nil {
		return nil, err
	}

	user.AvatarURL = avatarURL
	if err := user.Update("avatar_url"); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePassword update user password
func UpdatePassword(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	oldPassword := params.Args["oldPassword"].(string)

	if user.Password != utils.Encrypt(oldPassword) {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "User",
			Message: "password incorrect!",
		}
	}

	newPassword := params.Args["newPassword"].(string)
	if err := utils.ValidatePassword(newPassword); err != nil {
		return nil, err
	}

	user.Password = utils.Encrypt(newPassword)
	if err := user.Update("password"); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePhone update user phone
func UpdatePhone(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	user := rootValue["currentUser"].(models.User)
	password := params.Args["password"].(string)

	if user.Password != utils.Encrypt(password) {
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "User",
			Message: "password incorrect!",
		}
	}

	newPhone := params.Args["newPhone"].(string)
	if err := utils.ValidatePhone(newPhone); err != nil {
		return nil, err
	}

	smsCode := params.Args["smsCode"].(string)
	sessPhone := rootValue["phone"]
	sessSmsCode := rootValue["smsCode"]

	if sessPhone == nil || sessSmsCode == nil || newPhone != sessPhone || smsCode != sessSmsCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "smsCode",
			Message: "incorrect!",
		}
	}

	user.Phone = newPhone
	if err := user.Update("phone"); err != nil {
		return nil, err
	}

	return user, nil
}

// RelatedLoad load user
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.DeviceCharge:
		return v.LoadUser()
	case *models.DeviceCharge:
		return v.LoadUser()
	case models.DeviceParam:
		return v.LoadAuthor()
	case *models.DeviceParam:
		return v.LoadAuthor()
	case models.Device:
		return v.LoadUser()
	case *models.Device:
		return v.LoadUser()
	case models.UserExtend:
		return v.LoadUser()
	case *models.UserExtend:
		return v.LoadUser()
	case *models.UserLogin:
		return v.LoadUser()
	default:
		return nil, errors.LogicError{
			Type:    "Resolver",
			Field:   "User",
			Message: "RelatedLoad() error",
		}
	}
}
