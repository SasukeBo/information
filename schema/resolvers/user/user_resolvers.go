package user

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/schema/resolvers"
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

	rootValue["currentUserUUID"] = nil
	rootValue["smsCode"] = nil
	rootValue["setSession"] = []string{"currentUserUUID", "smsCode"}

	return user, nil
}

// Get get user by uuid
func Get(params graphql.ResolveParams) (interface{}, error) {
	uuid := params.Args["uuid"].(string)

	user := models.User{UUID: uuid}
	if err := user.GetByUUID(); err != nil {
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
		return nil, err
	}

	return users, nil
}

// UpdateAvatar update user avatar url
func UpdateAvatar(params graphql.ResolveParams) (interface{}, error) {
	var user models.User
	if err := getUser(params, &user); err != nil {
		return nil, err
	}

	avatarURL := params.Args["avatarURL"].(string)

	user.AvatarURL = avatarURL
	if _, err := models.Repo.Update(&user, "avatar_url"); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePassword update user password
func UpdatePassword(params graphql.ResolveParams) (interface{}, error) {
	var user models.User
	if err := getUser(params, &user); err != nil {
		return nil, err
	}

	oldPassword := params.Args["oldPassword"].(string)
	if user.Password != utils.Encrypt(oldPassword) {
		return nil, utils.LogicError{
			Message: "password not correct!",
		}
	}

	newPassword := params.Args["newPassword"].(string)
	user.Password = utils.Encrypt(newPassword)
	if _, err := models.Repo.Update(&user, "password"); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePhone update user phone
func UpdatePhone(params graphql.ResolveParams) (interface{}, error) {
	rootValue := params.Info.RootValue.(map[string]interface{})
	var user models.User
	if err := getUser(params, &user); err != nil {
		return nil, err
	}

	password := params.Args["password"].(string)
	if user.Password != utils.Encrypt(password) {
		return nil, utils.LogicError{
			Message: "password not correct!",
		}
	}

	newPhone := params.Args["newPhone"].(string)
	smsCode := params.Args["smsCode"].(string)
	sessPhone := rootValue["phone"]
	sessSmsCode := rootValue["smsCode"]

	if sessPhone == nil || sessSmsCode == nil || newPhone != sessPhone || smsCode != sessSmsCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, utils.ArgumentError{
			Field:   "smsCode",
			Message: "is not correct",
		}
	}

	user.Phone = newPhone
	if _, err := models.Repo.Update(&user, "phone"); err != nil {
		return nil, err
	}

	return user, nil
}

func getUser(params graphql.ResolveParams, user *models.User) error {
	uuid := params.Info.RootValue.(map[string]interface{})["currentUserUUID"]
	if uuid == nil {
		return utils.LogicError{
			Message: "user is not authenticated",
		}
	}

	user.UUID = uuid.(string)

	if err := models.Repo.Read(user, "uuid"); err != nil {
		return err
	}

	return nil
}

// RelatedLoad load user
func RelatedLoad(params graphql.ResolveParams) (interface{}, error) {
	var id int

	switch v := params.Source.(type) {
	case models.DeviceCharge:
		id = v.User.ID
	case models.DeviceParam:
		id = v.Author.ID
	case models.Device:
		id = v.User.ID
	case models.UserExtend:
		id = v.User.ID
	default:
		return nil, utils.LogicError{
			Message: "reloated user load error",
		}
	}

	user := models.User{ID: id}
	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}
