package email

type EmailInterface interface {
	EmailGetters
	EmailCompare
	EmailValidators
}

type EmailGetters interface {
	GetEmail() string
	IsValid() bool
	String() string
	Error() error
}

type EmailCompare interface {
	Contains(s string) bool
	EqualTo(other string) bool
	EqualToEmail(other EmailInterface) bool
}

type EmailValidators interface {
	// internaly validates the email
	Validate()
	IsEmpty() bool
	ContainsArrow() bool
}
