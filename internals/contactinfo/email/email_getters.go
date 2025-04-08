package email

var _ EmailGetters = (*Email)(nil)

// returns the email as a string or panics
func (e *Email) GetEmail() string {
	if !e.isValid {
		panic(e.err)
	}
	return e.email
}

// IsValid returns true if the email is valid
func (e *Email) IsValid() bool {
	return e.isValid
}

// String returns the email as a string
// and is used for printing
func (e *Email) String() string {
	return e.email
}

// Error returns the error if any
// otherwise nil
func (e *Email) Error() error {
	return e.err
}
