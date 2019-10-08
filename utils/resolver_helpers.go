package utils

import (
	"fmt"
	"github.com/SasukeBo/information/models"
	"github.com/graphql-go/graphql"
	"regexp"
)

var phoneRegexp = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`
var emailRegexp = `[^\\.\\s@:](?:[^\\s@:]*[^\\s@:\\.])?@[^\\.\\s@]+(?:\\.[^\\.\\s@]+)*`

// EmptyResult list query empty result
var EmptyResult = []interface{}{}

// ValidateStringEmpty validate str is or not empty, return resolvers.ArgumentError type error
func ValidateStringEmpty(str, field string) error {
	if str == "" {
		return models.Error{Message: fmt.Sprintf("%s can't be blank.", field)}
	}
	return nil
}

// ValidatePhone _
func ValidatePhone(phone string) error {
	re := regexp.MustCompile(phoneRegexp)
	if !re.Match([]byte(phone)) {
		return models.Error{Message: "invalid phone number."}
	}
	return nil
}

// ValidateEmail _
func ValidateEmail(email string) error {
	reg := regexp.MustCompile(emailRegexp)
	if !reg.Match([]byte(email)) {
		return models.Error{Message: "invalid email."}
	}

	return nil
}

// ValidatePassword validate password is or not legal
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return models.Error{Message: "password too short."}
	}
	return nil
}

// ValidateAccess _
// privSign 权限 sign
// privType 权限类型
func ValidateAccess(params *graphql.ResolveParams, privSign string, privType int) error {
	var role *models.Role
	var err error
	user := params.Info.RootValue.(map[string]interface{})["currentUser"].(models.User)

	if role, err = user.LoadRole(); err != nil {
		return err
	}

	if err := role.Validate(privSign, privType); err != nil {
		return err
	}

	return nil
}
