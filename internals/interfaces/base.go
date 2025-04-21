package interfaces

type BaseInterface[SELF any, CONTENT any, DTO BaseInterfaceDTO[DTO] ] interface {
	BaseInterfaceValue[SELF] // returns the values

	BaseInterfaceCompare[SELF] // compares the value with another

	BaseInterfaceClone // returns a new instance of the same type

	BaseInterfaceType[CONTENT] // returns the type

	BaseInterfaceInternal[CONTENT] // internal methods

	BaseInterfaceValidator[CONTENT] // returns the validator functions

	BaseInterfaceDTO[DTO] // returns the DTO of the value
}

// done
type BaseInterfaceValue[CONTENT any] interface {
	ValueOrPanic() CONTENT      // returns the value or panics if the value is not valid
	ValueOr(or CONTENT) CONTENT // returns the value or the given default value
	ValueOrErr() (CONTENT, error)
	Error() error   // returns the error if the value is not valid
	IsValid() bool  // checks if the value is valid
	String() string // returns the string representation of the value
}

// may
type BaseInterfaceCompare[SELF any, CONTENT any] interface {
	Compare(otherType SELF) bool
	Contains(s string) bool
	EqualTo(otherContent CONTENT) bool
}

// done
type BaseInterfaceValidator[SELF any] interface {
	// return the validator function // by default it returns nil
	Validator() func(SELF) (bool, error)

	ValidatorDefault() func(SELF) (bool, error) // returns the default validator function

	IsDefaultValidator() bool                 // checks if the validator is the default one
	IsExternalValidator() bool                // checks if the validator is the external one
	SetValidator(fn func(SELF) (bool, error)) // sets the validator function
	SetValidatorNil()                         // sets the validator function to nil
}

// done
type BaseInterfaceType interface {
	GetType() string // returns the type of the value as a string
}

// done
type BaseInterfaceClone[SELF any] interface {
	Clone() SELF // returns a new instance of the same type
}

// done
type BaseInterfaceInternal interface {
	// validate the value internally
	Validate()
}

// done
type BaseInterfaceDTO[DTO any] interface {
	GetDTO() DTO // returns the DTO of the value
}
