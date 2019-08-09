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

// ORMError orm error 封装
type ORMError struct {
	OrmErr  error
	Message string
}

// PrivError privilege error 封装
type PrivError struct {
	PrivErr error
	Message string
}

func (e ArgumentError) Error() string {
	return fmt.Sprintf("%v %v", e.Field, e.Message)
}

func (e LogicError) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

func (e ORMError) Error() string {
	return fmt.Sprintf("orm error: %s, %v", e.Message, e.OrmErr.Error())
}

func (e PrivError) Error() string {
	return fmt.Sprintf("privilege error: %s, %v", e.Message, e.PrivErr.Error())
}
