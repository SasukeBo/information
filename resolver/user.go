package resolver

import (
	"github.com/SasukeBo/information/models"
	"github.com/SasukeBo/information/utils"
	"github.com/astaxie/beego/orm"
	"github.com/graphql-go/graphql"
)

// SignUp _
func SignUp(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	phoneStr := params.Args["phone"].(string)
	msgCodeStr := params.Args["smsCode"].(string)
	passwordStr := params.Args["password"].(string)
	name := params.Args["name"].(string)

	rootValue := params.Info.RootValue.(map[string]interface{})

	sessPhone := rootValue["phone"]
	sessMsgCode := rootValue["smsCode"]

	user := models.User{Name: name}

	// validate phone
	if err := utils.ValidatePhone(phoneStr); err != nil {
		return nil, err
	}
	if sessPhone == nil || sessMsgCode == nil || phoneStr != sessPhone || msgCodeStr != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, models.Error{Message: "smsCode incorrect."}
	}
	user.Phone = phoneStr

	// validate password
	if err := utils.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}
	user.Password = utils.Encrypt(passwordStr)

	// 事务处理
	role := models.Role{RoleName: "default"}
	if err := o.Read(&role, "role_name"); err != nil {
		return nil, models.Error{Message: "get role failed.", OriErr: err}
	}
	user.Role = &role

	if _, err := o.Insert(&user); err != nil {
		return nil, models.Error{Message: "insert user failed.", OriErr: err}
	}

	rootValue["smsCode"] = nil
	rootValue["setSession"] = []string{"smsCode"}

	return user, nil
}

// ResetPassword is a gql resolver, reset user password
func ResetPassword(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
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
		return nil, models.Error{Message: "smsCode incorrect."}
	}
	if err := utils.ValidatePassword(passwordStr); err != nil {
		return nil, err
	}

	user := models.User{Phone: phoneStr}
	if err := o.Read(&user, "phone"); err != nil {
		return nil, models.Error{Message: "get user failed.", OriErr: err}
	}

	user.Password = utils.Encrypt(passwordStr)
	if _, err := o.Update(&user, "password"); err != nil {
		return nil, models.Error{Message: "update password failed.", OriErr: err}
	}

	rootValue["currentUser"] = nil
	rootValue["smsCode"] = nil
	rootValue["setSession"] = []string{"currentUser", "smsCode"}

	return user, nil
}

// GetUser _
func GetUser(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	id := params.Args["id"].(int)

	user := models.User{ID: id}
	if err := o.Read(&user, "id"); err != nil {
		return nil, models.Error{Message: "get user failed.", OriErr: err}
	}

	return user, nil
}

// ListUser _
func ListUser(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	namePattern := params.Args["namePattern"]
	phone := params.Args["phone"]
	email := params.Args["email"]

	var users []*models.User

	qs := o.QueryTable("user")

	if namePattern != nil {
		qs = qs.Filter("UserExtend__name__icontains", namePattern)
	}

	if phone != nil {
		qs = qs.Filter("phone", phone)
	}

	if email != nil {
		qs = qs.Filter("UserExtend__email", email)
	}

	if _, err := qs.All(&users); err != nil {
		return nil, models.Error{Message: "list user failed.", OriErr: err}
	}

	return users, nil
}

// UpdateUser _
func UpdateUser(params graphql.ResolveParams) (interface{}, error) {
	o := orm.NewOrm()
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	avatarURL := params.Args["avatarURL"].(string)

	if err := utils.ValidateStringEmpty(avatarURL, "avatarURL"); err != nil {
		return nil, err
	}

	user.AvatarURL = avatarURL
	if _, err := o.Update(&user, "avatar_url"); err != nil {
		return nil, models.Error{Message: "update user avatar_url failed.", OriErr: err}
	}

	return user, nil
}

/*
// UpdatePassword _
func UpdatePassword(params graphql.ResolveParams) (interface{}, error) {
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)
	oldPassword := params.Args["oldPassword"].(string)

	if user.Password != utils.Encrypt(oldPassword) {
		return nil, errors.Error{
			Type:    "Resolver",
			Field:   "password",
			Message: "password incorrect.",
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
		return nil, errors.Error{
			Type:    "Resolver",
			Field:   "password",
			Message: "password incorrect.",
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
		return nil, errors.Error{
			Type:    "Resolver",
			Field:   "smsCode",
			Message: "smsCode incorrect.",
		}
	}

	user.Phone = newPhone
	if err := user.Update("phone"); err != nil {
		return nil, err
	}

	return user, nil
}
*/

// LoadUser _
func LoadUser(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Device:
		return v.LoadUser()
	case *models.Device:
		return v.LoadUser()
	case *models.UserLogin:
		return v.LoadUser()
	default:
		return nil, models.Error{Message: "load related user failed."}
	}
}
