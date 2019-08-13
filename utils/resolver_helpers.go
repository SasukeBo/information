package utils

import (
	"regexp"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/models/errors"
	"github.com/SasukeBo/information/models"
)

var phoneRegexp = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`
var emailRegexp = `[^\\.\\s@:](?:[^\\s@:]*[^\\s@:\\.])?@[^\\.\\s@]+(?:\\.[^\\.\\s@]+)*`

// EmptyResult list query empty result
var EmptyResult = []interface{}{}

// ValidateStringEmpty validate str is or not empty, return resolvers.ArgumentError type error
func ValidateStringEmpty(str, field string) error {
	if str == "" {
		return errors.LogicError{
			Type:    "Validate",
			Field:   field,
			Message: "can not be blank",
		}
	}
	return nil
}

// ValidatePhone _
func ValidatePhone(phone string) error {
	re := regexp.MustCompile(phoneRegexp)
	if !re.Match([]byte(phone)) {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "phone",
			Message: "invalid.",
		}
	}
	return nil
}

// ValidateEmail _
func ValidateEmail(email string) error {
	reg := regexp.MustCompile(emailRegexp)
	if !reg.Match([]byte(email)) {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "email",
			Message: "invalid.",
		}
	}

	return nil
}

// ValidatePassword validate password is or not legal
func ValidatePassword(password string) error {
	if len(password) < 6 {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "password",
			Message: "too short",
		}
	}
	return nil
}

// ValidateAccess _
func ValidateAccess(params *graphql.ResolveParams, privSign string) error {
	currentUserUUID := params.Info.RootValue.(map[string]interface{})["currentUserUUID"].(string)
	user := models.User{UUID: currentUserUUID}
	if err := user.GetBy("uuid"); err != nil {
		return err
	}

	var role *models.Role
	var err error

	if role, err = user.LoadRole(); err != nil {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "role",
			Message: "load error",
			OriErr:  err,
		}
	}

	if err := role.Validate(privSign); err != nil {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "privilege",
			Message: "access error",
			OriErr:  err,
		}
	}

	return nil
}
