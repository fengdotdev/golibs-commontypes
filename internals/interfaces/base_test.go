package interfaces_test

import (
	"errors"
	"testing"

	"github.com/fengdotdev/golibs-commontypes/internals/interfaces"
	"github.com/fengdotdev/golibs-testing/assert"
)

type BaseInterface = interfaces.BaseInterface[string]

func FooValid(b BaseInterface) bool {
	return b.IsValid()
}

var _ BaseInterface = &BaseMock{}

type BaseMock struct {
	content string
	valid   bool
	err     error
}

func NewValidMock(content string) *BaseMock {
	return &BaseMock{
		content: content,
		valid:   true,
		err:     nil,
	}
}

func NewInvalidMock(content string) *BaseMock {
	return &BaseMock{
		content: content,
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

func TestBase_Zero(t *testing.T) {

	zero := BaseMock{}

	assert.Equal(t, zero.ValueOr("default"), "default")

	v, err := zero.ValueOrErr()
	assert.Equal(t, v, "")
	assert.Equal(t, err.Error(), "some error")

	value := ""

	assertPanic(t, func() {
		value = "prelude panic"
		value = zero.ValueOrPanic()
	})

	assert.Equal(t, value, "prelude panic")

}

func TestBase_Valid(t *testing.T) {
	valid := NewValidMock("valid")

	assert.Equal(t, valid.ValueOr("default"), "valid")

	v, err := valid.ValueOrErr()
	assert.Equal(t, v, "valid")
	assert.Nil(t, err)

	value := ""

	assertPanic(t, func() {
		value = valid.ValueOrPanic()
	})

	assert.Equal(t, value, "valid")
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r != nil {
			assert.TrueWithMessage(t, true, "the code panic and recovered")
		}
	}()
	f()
	assert.TrueWithMessage(t, false, "the code did not panic")
}

func assertNotPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			assert.TrueWithMessage(t, false, "the code did  panic")
		}
	}()
	f()
	assert.TrueWithMessage(t, true, "the code did not panic")
	assert.
}
