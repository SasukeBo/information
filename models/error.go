package models

// Error _
type Error struct {
	Message string
	OriErr  error
}

// Error _
func (e Error) Error() string {
	return e.Message
}
