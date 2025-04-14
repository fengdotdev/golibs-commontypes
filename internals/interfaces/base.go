package interfaces

type BaseInterface[T any,OBJ any, DTO any] interface {
	BaseInterfaceValue[T] // returns the values

	BaseInterfaceCompare[T] // compares the value with another

	BaseInterfaceClone[OBJ] // returns a new instance of the same type

	BaseInterfaceType[T] // returns the type

	BaseInterfaceInternal[T] // internal methods

	BaseInterfaceValidator[T] // returns the validator functions

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

type BaseInterfaceCompare[T any,Obj any] interface {
	Compare(other BaseInterface[T]) bool // compares the value with another value of the same type

	Contains(s string) bool // checks if the value contains the given string
	EqualTo(other T) bool   // compares the value with another value of the same type
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

type BaseInterfaceClone[T any] interface {
	Clone() BaseInterface[T] // returns a new instance of the same type
}

type BaseInterfaceInternal[T any] interface {
	// validate the value internally
	Validate()
}

type BaseInterfaceDTO[DTO any] interface {
	GetDTO() DTO // returns the DTO of the value
}
