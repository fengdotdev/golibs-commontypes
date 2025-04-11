package interfaces_test

import (
	"errors"

	"github.com/fengdotdev/golibs-commontypes/internals/interfaces"
)

// This file contains a mock implementation of the BaseInterface for testing purposes.
// It is not intended to be used in production code.
// The mock implementation is used to test the behavior of the BaseInterface

type BaseInterface = interfaces.BaseInterface[string]

func fooValid(b BaseInterface) bool {
	return b.IsValid()
}

var _ BaseInterface = &BaseMock{}

type BaseMock struct {
	content string
	valid   bool
	err     error
}

func newValidMock() *BaseMock {
	return &BaseMock{
		content: "valid",
		valid:   true,
		err:     nil,
	}
}

func newInvalidMock() *BaseMock {
	return &BaseMock{
		content: "invalid",
		valid:   false,
		err:     errors.New("some error"),
	}
}

func (b *BaseMock) String() string {
	return b.content
}

func (b *BaseMock) ValueOr(or string) string {
	if b.valid && b.err == nil {
		return b.content
	}
	return or
}

func (b *BaseMock) ValueOrPanic() string {
	if b.valid && b.err == nil {
		return b.content
	}
	panic(b.err)
}

func (b *BaseMock) ValueOrErr() (string, error) {
	if b.valid && b.err == nil {
		return b.content, nil
	}
	return "", errors.New("some error")
}

func (b *BaseMock) IsValid() bool {
	return b.valid
}

func (b *BaseMock) Compare(other interfaces.BaseInterface[string]) bool {
	if other == nil {
		return false
	}
	return b.content == other.String()
}

func (b *BaseMock) Error() error {
	if b.valid {
		return nil
	}
	return b.err
}

func (b *BaseMock) Contains(s string) bool {
	panic("implement me")
}

func (b *BaseMock) EqualTo(other string) bool {
	panic("implement me")
}

func (b *BaseMock) Validate() {
	panic("implement me")
}

func (b *BaseMock) GetType() string {
	return "BaseMock"
}

func (b *BaseMock) Clone() interfaces.BaseInterface[string] {
	return &BaseMock{
		content: b.content,
		valid:   b.valid,
		err:     b.err,
	}
}
