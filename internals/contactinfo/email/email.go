package email

type Email struct {
	email   string
	isValid bool
	err     error
}



func (e *Email) IsEmpty() bool {
	return e.email == ""
}

func (e *Email) Contains(s string) bool {
	panic("implement me")
}

func (e *Email) ContainsArrow() bool {
	return e.Contains("@")
}

func (e *Email) Validate() {

	if e.IsEmpty() {
		e.err = ErrEmailEmpty
		return
	}

	if !e.ContainsArrow() {
		e.err = ErrInvalidEmailFormat
		return
	}

	// email is valid
	e.isValid = true
	e.err = nil
}

