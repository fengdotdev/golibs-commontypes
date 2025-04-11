package interfaces

type BaseInterface[T any] interface {
	ValueOrPanic() T // returns the value or panics if the value is not valid
	ValueOr(or T) T // returns the value or the given default value
	ValueOrErr() (T, error)
	String() string // returns the string representation of the value
	IsValid() bool // checks if the value is valid
	Compare(other BaseInterface[T]) bool // compares the value with another value of the same type
	Error() error // returns the error if the value is not valid
	Contains(s string) bool // checks if the value contains the given string
	EqualTo(other T) bool // compares the value with another value of the same type

	Clone() BaseInterface[T] // returns a new instance of the same type

	GetType() string // returns the type of the value as a string

	//internaly validates
	Validate()
}
