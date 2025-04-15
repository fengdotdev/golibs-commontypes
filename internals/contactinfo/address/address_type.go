package address

type Address struct {
	content map[string]string
	valid   bool
	err     error
}
