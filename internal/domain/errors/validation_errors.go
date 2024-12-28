package errors

import "fmt"

var (
	FieldRequired = func(field string) ValidationError {
		return New(field, "Is required")
	}
	OutRange = func(field string, min, max uint8) ValidationError {
		return New(field, fmt.Sprintf("Length should be between %d and %d ", min, max))
	}
	FormatInvalidad = func(field string) ValidationError {
		return New(field, "Format is not valid")
	}
	InvalidUUID = func (field string) ValidationError  {
		return New(field, "UUID format is not valid")
	}
)

type ValidationError struct {
	Field string
	Message string
}

func New(field, message string) ValidationError {
	return ValidationError{field, message}
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", v.Field, v.Message)
}