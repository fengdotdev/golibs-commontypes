package email_test

import (
	"testing"

	"github.com/fengdotdev/golibs-commontypes/internals/contactinfo/email"
	"github.com/fengdotdev/golibs-testing/assert"
)

func TestEmail_IsEmpty(t *testing.T) {
	e := email.NewEmail("")
	assert.Equal(t, e.IsEmpty(), true)

	e = email.NewEmail("test@example.com")
	assert.Equal(t, e.IsEmpty(), false)
}

func TestEmail_Contains(t *testing.T) {
	e := email.NewEmail("test@example.com")
	assert.Equal(t, e.Contains("test"), true)
	assert.Equal(t, e.Contains("example"), true)
	assert.Equal(t, e.Contains("missing"), false)
}

func TestEmail_ContainsArrow(t *testing.T) {
	e := email.NewEmail("test@example.com")
	assert.Equal(t, e.ContainsArrow(), true)

	e = email.NewEmail("testexample.com")
	assert.Equal(t, e.ContainsArrow(), false)
}

func TestEmail_HasValidDomain(t *testing.T) {
	e := email.NewEmail("test@example.com")
	assert.Equal(t, e.HasValidDomain(), true)

	e = email.NewEmail("test@example")
	assert.Equal(t, e.HasValidDomain(), false)
}

func TestEmail_HasCorrectlen(t *testing.T) {
	e := email.NewEmail("a@b.c")
	assert.Equal(t, e.HasCorrectlen(), true)

	e = email.NewEmail("a@b.c" + string(make([]byte, 250))) // Exceeds 254 characters
	assert.Equal(t, e.HasCorrectlen(), false)

	e = email.NewEmail("a@b") // Less than 5 characters
	assert.Equal(t, e.HasCorrectlen(), false)
}
