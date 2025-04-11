package email_test

import (
	"strings"
	"testing"

	"github.com/fengdotdev/golibs-commontypes/internals/contactinfo/email"
	"github.com/fengdotdev/golibs-testing/assert"
)

type Client struct {
	Name  string
	Email email.Email
}

func mockValidator(s string) bool {
	if strings.Contains(s, "inc.com") {
		return false
	}
	return true
}

func TestEmail_asUse(t *testing.T) {

	fooCLient := Client{
		Name:  "Foo",
		Email: email.NewEmail("foo@inc.com"),
	}

	barCLient := Client{
		Name:  "Bar",
		Email: email.NewEmail("bar@inc.com"),
	}

	equal := fooCLient.Email.Compare(&barCLient.Email)

	assert.Equal(t, equal, false)

	assert.True(t, fooCLient.Email.IsValid())

	assert.True(t, barCLient.Email.IsValid())

	troll := Client{
		Name:  "Troll",
		Email: email.NewEmail("troll"),
	}

	assert.Equal(t, troll.Email.IsValid(), false)

	// use external validator


	email.SetExternalValidator(mockValidator)


	assert.False(t, mockValidator(barCLient.Email.String()))
	assert.False(t, barCLient.Email.IsValid())

	
}
