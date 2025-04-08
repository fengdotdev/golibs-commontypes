package email

func NewEmail(email string) Email {
	return Email{
		email:   email,
		isValid: false,
		err:     ErrEmailNoValidated,
	}
}

func NewEmailWithValidation(email string) (Email, error) {
	mayemail := NewEmail(email)

	mayemail.Validate()

	return mayemail, mayemail.err
}
