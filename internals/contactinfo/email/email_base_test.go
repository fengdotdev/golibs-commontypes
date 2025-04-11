package email_test

import (
	"testing"

	"github.com/fengdotdev/golibs-commontypes/internals/contactinfo/email"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestEmail_Empty(t *testing.T) {
	e := email.NewEmail("")

	assert.Equal(t, e.ValueOr("default"), "default")

	v, err := e.ValueOrErr()
	assert.Equal(t, v, "")
	assert.Error(t, err)
	assert.AssertPanic(t, func() {
		_ = e.ValueOrPanic()
	})
}

func TestEmail_Valid(t *testing.T) {
	e := email.NewEmail("test@example.com")
	e.Validate()

	assert.Equal(t, e.ValueOr("default"), "test@example.com")

	v, err := e.ValueOrErr()
	assert.Equal(t, v, "test@example.com")
	assert.Nil(t, err)

	assert.AssertNotPanic(t, func() {
		_ = e.ValueOrPanic()
	})
}

func TestEmail_Invalid(t *testing.T) {
	e := email.NewEmail("invalid-email")
	e.Validate()

	assert.Equal(t, e.ValueOr("default"), "default")

	v, err := e.ValueOrErr()
	assert.Equal(t, v, "")
	assert.Equal(t, err, email.ErrInvalidEmailFormat)

	assert.AssertPanic(t, func() {
		_ = e.ValueOrPanic()
	})
}

func TestEmail_Compare(t *testing.T) {
	e1 := email.NewEmail("test@example.com")
	e2 := email.NewEmail("test@example.com")
	e3 := email.NewEmail("different@example.com")

	assert.Equal(t, e1.Compare(&e2), true)
	assert.Equal(t, e1.Compare(&e3), false)
}

func TestEmail_Clone(t *testing.T) {
	e := email.NewEmail("test@example.com")
	e.Validate()

	clone := e.Clone()

	assert.Equal(t, clone.String(), e.String())
	assert.Equal(t, clone.IsValid(), e.IsValid())
	assert.Equal(t, clone.Error(), e.Error())
	assert.NotEqual(t, &clone, &e) // Ensure it's a different instance
}
