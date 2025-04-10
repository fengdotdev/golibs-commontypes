package commons

// Assumes that the type is created with some error and needs to be validated

type ValidInterface interface {
	IsValid() bool
	Error() error
}
