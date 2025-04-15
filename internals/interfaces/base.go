package interfaces

type BaseInterface[CONTENT any, SELF any, DTO any] interface {
	BaseInterfaceValue[CONTENT] // returns the values

	BaseInterfaceCompare[CONTENT, SELF, DTO] // compares the value with another

	BaseInterfaceClone[CONTENT, SELF, DTO] // returns a new instance of the same type

	BaseInterfaceType[CONTENT] // returns the type

	BaseInterfaceInternal[CONTENT] // internal methods

	BaseInterfaceValidator[CONTENT] // returns the validator functions

	BaseInterfaceDTO[DTO] // returns the DTO of the value
}

type BaseInterfaceValue[T any] interface {
	ValueOrPanic() T // returns the value or panics if the value is not valid
	ValueOr(or T) T  // returns the value or the given default value
	ValueOrErr() (T, error)
	Error() error   // returns the error if the value is not valid
	IsValid() bool  // checks if the value is valid
	String() string // returns the string representation of the value
}

type BaseInterfaceCompare[CONTENT any, SELF any, DTO any] interface {
	Compare(other BaseInterface[CONTENT, SELF, DTO]) bool // compares the value with another value of the same type

	Contains(s string) bool     // checks if the value contains the given string
	EqualTo(other CONTENT) bool // compares the value with another value of the same type
}

type BaseInterfaceValidator[T any] interface {
	// return the validator function // by default it returns nil
	Validator() func(T) (bool, error)

	ValidatorDefault() func(T) (bool, error) // returns the default validator function

	IsDefaultValidator() bool              // checks if the validator is the default one
	IsExternalValidator() bool             // checks if the validator is the external one
	SetValidator(fn func(T) (bool, error)) // sets the validator function
	SetValidatorNil()                      // sets the validator function to nil
}

type BaseInterfaceType[T any] interface {
	GetType() string // returns the type of the value as a string
}

type BaseInterfaceClone[CONTENT any, SELF any, DTO any] interface {
	Clone() BaseInterface[CONTENT, SELF, DTO] // returns a new instance of the same type
}

type BaseInterfaceInternal[T any] interface {
	// validate the value internally
	Validate()
}

type BaseInterfaceDTO[DTO any] interface {
	GetDTO() DTO // returns the DTO of the value
}
