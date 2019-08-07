package resolvers

import (
	"github.com/SasukeBo/information/utils"
	"regexp"
)

var phoneRegexp = `^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[35678]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|66\d{2})\d{6}$`

// ValidateStringEmpty validate str is or not empty, return resolvers.ArgumentError type error
func ValidateStringEmpty(str, field string) error {
	if str == "" {
		return utils.ArgumentError{
			Field:   "roleName",
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
		return utils.ArgumentError{
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
		return utils.ArgumentError{
			Field:   "password",
			Message: "too short",
		}
	}
	return nil
}
