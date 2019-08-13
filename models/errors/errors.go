package errors

import (
	"fmt"
)

// LogicError Error with logic type
type LogicError struct {
	Type    string
	Field   string
	Message string
	OriErr  error
}

func (e LogicError) Error() string {
	str := fmt.Sprintf("%s Error: %s %s", e.Type, e.Field, e.Message)
	if e.OriErr != nil {
		str = fmt.Sprintf("%s,\n %s", str, e.OriErr.Error())
	}

	return str
}
