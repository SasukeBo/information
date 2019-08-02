package utils

import (
	"fmt"
)

// ArgumentError 变量类型错误
type ArgumentError struct {
	Field   string
	Message string
}

// LogicError 逻辑类型错误
type LogicError struct {
	Message string
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("%v %v", e.Field, e.Message)
}

func (e LogicError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}
