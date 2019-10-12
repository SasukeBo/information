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
	phone := params.Args["phone"].(string)
	user := &models.User{Phone: phone}
	if err := o.Read(user, "phone"); err == nil {
		return nil, models.Error{Message: "phone has already been registered."}
	}

	msgCode := params.Args["smsCode"].(string)
	password := params.Args["password"].(string)
	name := params.Args["name"].(string)
	user.Name = name

	rootValue := params.Info.RootValue.(map[string]interface{})

	sessPhone := rootValue["phone"]
	sessMsgCode := rootValue["smsCode"]

	// 接口类型是可以和具体类型直接比较的
	if sessPhone == nil || sessMsgCode == nil || phone != sessPhone || msgCode != sessMsgCode {
		// 用户发送验证码的手机号与提交注册时的手机号不匹配，按照验证码不正确处理
		return nil, models.Error{Message: "smsCode incorrect."}
	}
	user.Phone = phone

	// validate password
	if err := utils.ValidatePassword(password); err != nil {
		return nil, err
	}
	user.Password = utils.Encrypt(password)

	// 事务处理
	role := models.Role{RoleName: "default"}
	if err := o.Read(&role, "role_name"); err != nil {
		return nil, models.Error{Message: "get role failed.", OriErr: err}
	}
	user.Role = &role

	if _, err := o.Insert(user); err != nil {
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
	rootValue := params.Info.RootValue.(map[string]interface{})
	user := rootValue["currentUser"].(models.User)
	newPhone := params.Args["newPhone"]
	smsCode := params.Args["smsCode"]
	updates := []string{}

	if avatarURL := params.Args["avatarURL"]; avatarURL != nil {
		user.AvatarURL = avatarURL.(string)
		updates = append(updates, "avatar_url")
	}

	newPassword := params.Args["newPassword"]
	password := params.Args["password"]
	if newPassword != nil {
		if password == nil || user.Password != utils.Encrypt(password.(string)) {
			return nil, models.Error{Message: "password incorrect."}
		}

		user.Password = utils.Encrypt(newPassword.(string))
		updates = append(updates, "password")
	}

	if newPhone != nil {
		if password == nil || user.Password != utils.Encrypt(password.(string)) {
			return nil, models.Error{Message: "password incorrect."}
		}

		sessPhone := rootValue["phone"]
		sessMsgCode := rootValue["smsCode"]
		if sessPhone == nil || sessMsgCode == nil || sessPhone != newPhone || sessMsgCode != smsCode {
			return nil, models.Error{Message: "smsCode incorrect."}
		}

		user.Phone = newPhone.(string)
		updates = append(updates, "phone")
	}

	if len(updates) == 0 {
		return user, nil
	}

	if _, err := o.Update(&user, updates...); err != nil {
		return nil, models.Error{Message: "update user failed.", OriErr: err}
	}

	return user, nil
}

// LoadUser _
func LoadUser(params graphql.ResolveParams) (interface{}, error) {
	switch v := params.Source.(type) {
	case models.Device:
		return v.LoadUser()
	case *models.Device:
		return v.LoadUser()
	case *models.UserLogin:
		return v.LoadUser()
	case models.UserLogin:
		return v.LoadUser()
	case *models.Product:
		return v.LoadUser()
	case models.Product:
		return v.LoadUser()
	default:
		return nil, models.Error{Message: "load related user failed."}
	}
}
