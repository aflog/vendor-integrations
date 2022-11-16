package vendors

import "fmt"

type ErrorType string

const (
	ErrorTypeConnection   ErrorType = "Connection Problem"
	ErrTypeResponseFormat           = "Invalid Response Format"
	ErrTypeUnimplemented            = "Unimplemented method"
)

type Error struct {
	ErrType ErrorType
	Msg     string
}

func (e *Error) Error() string {
	return fmt.Sprintf("Error Type: %s, Message: %s", e.ErrType, e.Msg)
}
