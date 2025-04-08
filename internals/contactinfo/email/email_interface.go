package email

type EmailInterface interface {
	GetEmail() string
	IsValid() bool
	String() string
	EqualTo(other string) bool
	EqualToEmail(other EmailInterface) bool
}
