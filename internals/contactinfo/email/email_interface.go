package email

type EmailInterface interface {
	EmailInners
}

type EmailInners interface {
	Validate()
	IsEmpty() bool
	ContainsArrow() bool
}
