package email

var _ BaseInterface = (*Email)(nil)

func (e *Email) ValueOrPanic() string {
	if !e.IsValid() {
		panic("invalid email")
	}
	return e.email
}

func (e *Email) ValueOr(or string) string {
	if !e.IsValid() {
		return or
	}
	return e.email
}

func (e *Email) ValueOrErr() (string, error) {
	if !e.IsValid() {
		return "", e.err
	}
	return e.email, nil
}

func (e *Email) String() string {
	return e.email
}


//  check if the email is valid
// if the email never was validated, it will be validated now 
// if the email was validated, it will be returned as is 
func (e *Email) IsValid() bool {
	if e.err == ErrEmailNoValidated {
		e.Validate()
	}

	return e.isValid
}

func (e *Email) Compare(other BaseInterface) bool {
	otherEmail, ok := other.(*Email)
	if !ok {
		return false
	}
	return e.email == otherEmail.email
}

func (e *Email) Error() error {
	return e.err
}

func (e *Email) EqualTo(other string) bool {
	return e.email == other
}

func (e *Email) Clone() BaseInterface {
	return &Email{
		email:   e.email,
		isValid: e.isValid,
		err:     e.err,
	}
}

func (e *Email) GetType() string {
	return "Email"
}
