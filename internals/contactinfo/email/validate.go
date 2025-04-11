package email

var externalValidatorFN func(email string) bool = nil


// SetExternalValidator sets the external validator function for email validation.
// This function will be used to validate the email address if provided.
// If the function returns false, the email will be considered invalid.
// If the function is nil, the default validation logic will be used.
// Careful using this function, as it will override the default validation logic.
func SetExternalValidator(fn func(string) bool) {
	externalValidatorFN = fn
}


// SetExternalValidatorNil sets the external validator function to nil.
// This will revert to the default validation logic.
func SetExternalValidatorNil() {
	externalValidatorFN = nil
}

func (e *Email) Validate() {
	// external Validation Have priority
	if externalValidatorFN != nil {
		if externalValidatorFN(e.email) {
			e.isValid = true
			e.err = nil
		}

		// email is valid
		e.isValid = false
		e.err = ErrEmailInvalid
		return
	}

	if e.IsEmpty() {
		e.err = ErrEmailEmpty
		return
	}

	if !e.ContainsArrow() {
		e.err = ErrInvalidEmailFormat
		return
	}

	if !e.HasCorrectlen() { // Check for correct length
		e.err = ErrInvalidEmailLength
	}

	if !e.HasValidDomain() { // Check for a valid domain
		e.err = ErrInvalidEmailDomain
		return
	}

	// email is valid
	e.isValid = true
	e.err = nil
}
