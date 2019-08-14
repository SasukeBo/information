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
	var oriMsg = ""
	if e.OriErr != nil {
		oriMsg = e.OriErr.Error()
	}

	return fmt.Sprintf(
		`{"type": "%s Error", "field": "%s", "message": "%s", "originMessage": "%s"}`,
		e.Type,
		e.Field,
		e.Message,
		oriMsg,
	)
}
