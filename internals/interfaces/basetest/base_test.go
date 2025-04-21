package interfaces_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestBase_Zero(t *testing.T) {

	zero := BaseMock{}

	assert.Equal(t, zero.ValueOr("default"), "default")

	v, err := zero.ValueOrErr()
	assert.Equal(t, v, "")
	assert.Equal(t, err.Error(), "some error")

	assert.AssertPanic(t, func() {
		_ = zero.ValueOrPanic()
	})

	assert.Nil(t, zero.Error())

	// using the default validator

	assert.True(t, zero.IsDefaultValidator())
	assert.False(t, zero.IsExternalValidator())
}

func TestBase_Valid(t *testing.T) {
	valid := newValidMock()

	assert.Equal(t, valid.ValueOr("default"), "valid")

	v, err := valid.ValueOrErr()
	assert.Equal(t, v, "valid")
	assert.Nil(t, err)

	assert.AssertNotPanic(t, func() {
		_ = valid.ValueOrPanic()
	})

	assert.AssertFail(t, func(t *testing.T) {
		assert.AssertPanic(t, func() {
			_ = valid.ValueOrPanic()
		})
	})

	assert.True(t, valid.IsDefaultValidator())
	assert.False(t, valid.IsExternalValidator())

	if valid.Validator() != nil {
		t.Errorf("Validator should not be nil")
	}

}
