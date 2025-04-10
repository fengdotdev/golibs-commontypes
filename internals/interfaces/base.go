package interfaces

type BaseInterface[T any] interface {
	ValueOrPanic() T
	ValueOr(or T) T
	ValueOrErr() (T, error)
	String() string
	IsValid() bool
}
