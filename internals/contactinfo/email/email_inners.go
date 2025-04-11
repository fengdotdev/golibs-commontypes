package email

import "strings"

var _ EmailInners = (*Email)(nil)

func (e *Email) IsEmpty() bool {
	return e.email == ""
}

func (e *Email) Contains(s string) bool {
	return strings.Contains(e.email, s)
}

func (e *Email) ContainsArrow() bool {
	return e.Contains("@")
}

func (e *Email) HasValidDomain() bool {
	return e.Contains(".")
}

func (e *Email) HasCorrectlen() bool {
	return len(e.email) <= 254 && len(e.email) >= 5 // x@x.x
}
