package email

import "errors"

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrEmailEmpty         = errors.New("email cannot be empty")
	ErrEmailInvalid       = errors.New("email is invalid")
	ErrEmailNoValidated   = errors.New("email has not been validated yet")
	ErrEmailTooLong       = errors.New("email exceeds maximum length of 254 characters")
	ErrEmailIsTooShort    = errors.New("email is too short")
	ErrInvalidEmailLength = errors.New("invalid email length")
	ErrInvalidEmailDomain = errors.New("invalid email domain")
)
