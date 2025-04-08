package email


type Email struct {
	email   string
	isValid bool
	err     error
}

func NewEmail(email string) Email {
	return Email{
		email:   email,
		isValid: false,
		err:     nil,
	}
}

func NewEmailWithValidation(email string) (Email, error) {
	mayemail := NewEmail(email)
	if !mayemail.IsValid() {
		return Email{
			email:   email,
			isValid: false,
			err:     ErrInvalidEmailFormat,
		}, ErrEmailInvalid
	}
	return mayemail, nil
}

func (e *Email) GetEmail() string {
	return e.email
}

func (e *Email) IsValid() bool {
	if len(e.email) == 0 {
		e.err = ErrEmailEmpty
		return false
	}
	for _, char := range e.email {
		if char == '@' {
			return true
		}
	}
	return false
}

func (e *Email) String() string {
	return e.email
}
