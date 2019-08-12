package utils

import (
	"regexp"

	"github.com/graphql-go/graphql"

	"github.com/SasukeBo/information/errors"
	"github.com/SasukeBo/information/models"
)

var phoneRegexp = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`

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

// ValidatePhone validate phone number is or not legal
func ValidatePhone(phone string) error {
	if err := ValidateStringEmpty(phone, "phone"); err != nil {
		return err
	}

	re := regexp.MustCompile(phoneRegexp)
	if !re.Match([]byte(phone)) {
		return errors.LogicError{
			Type:    "Validate",
			Field:   "phone",
			Message: "is not a valid phone number",
		}
	}
	return nil
}

// ValidatePassword validate password is or not legal
func ValidatePassword(password string) error {
	if err := ValidateStringEmpty(password, "password"); err != nil {
		return err
	}
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
