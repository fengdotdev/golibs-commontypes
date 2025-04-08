package email

import "errors"

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrEmailEmpty         = errors.New("email cannot be empty")
	ErrEmailInvalid       = errors.New("email is invalid")
)
