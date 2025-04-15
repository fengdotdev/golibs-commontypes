package interfaces_test

import (
	"errors"

	"github.com/fengdotdev/golibs-commontypes/internals/interfaces"
)

// This file contains a mock implementation of the BaseInterface for testing purposes.
// It is not intended to be used in production code.
// The mock implementation is used to test the behavior of the BaseInterface

type BaseInterface = interfaces.BaseInterface[string, BaseMock, BaseMockDTO]

func fooValid(b BaseInterface) bool {
	return b.IsValid()
}

var (
	validfn func(string) (bool, error) = nil
)

var _ BaseInterface = &BaseMock{}

type BaseMock struct {
	content string
	valid   bool
	err     error
}

type BaseMockDTO struct {
	Content string `json:"basemock"`
	Valid   bool   `json:"valid"`
	Err     error  `json:"err"`
}

type ValidBaseMockDTO struct {
	Content string `json:"basemock"`
}

//constructors

func newBaseMockFromDTO(dto BaseMockDTO) *BaseMock {
	return &BaseMock{
		content: dto.Content,
		valid:   dto.Valid,
		err:     dto.Err,
	}
}

func newBaseMockFromValidDTO(dto ValidBaseMockDTO) *BaseMock {
	return &BaseMock{
		content: dto.Content,
		valid:   true,
		err:     nil,
	}
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

// value
var _ interfaces.BaseInterfaceValue[string] = &BaseMock{}

func (b *BaseMock) ValueOrPanic() string {
	if b.valid && b.err == nil {
		return b.content
	}
	panic(b.err)
}

func (b *BaseMock) ValueOr(or string) string {
	if b.valid && b.err == nil {
		return b.content
	}
	return or
}

func (b *BaseMock) ValueOrErr() (string, error) {
	if b.valid && b.err == nil {
		return b.content, nil
	}
	return "", errors.New("some error")
}

func (b *BaseMock) Error() error {
	if b.valid {
		return nil
	}
	return b.err
}

func (b *BaseMock) IsValid() bool {
	return b.valid
}

func (b *BaseMock) String() string {
	return b.content
}

// compare
var _ interfaces.BaseInterfaceCompare[string] = &BaseMock{}

func (b *BaseMock) Compare(other interfaces.BaseInterface[string]) bool {
	if other == nil {
		return false
	}
	return b.content == other.String()
}

func (b *BaseMock) Contains(s string) bool {
	panic("implement me")
}

func (b *BaseMock) EqualTo(other string) bool {
	panic("implement me")
}

//clone

var _ interfaces.BaseInterfaceClone[string] = &BaseMock{}

func (b *BaseMock) Clone() interfaces.BaseInterface[string] {
	return &BaseMock{
		content: b.content,
		valid:   b.valid,
		err:     b.err,
	}
}

// type
var _ interfaces.BaseInterfaceType[string] = &BaseMock{}

func (b *BaseMock) GetType() string {
	return "BaseMock"
}

//internal

func (b *BaseMock) Validate() {
	if b.IsExternalValidator() {
		if validfn != nil {
			b.valid, b.err = validfn(b.content)
			return
		}
	}

	b.validate()

}

func (b *BaseMock) validate() {
	if b.valid && b.err == nil {
		return
	}

}

// validator
var _ interfaces.BaseInterfaceValidator[string] = &BaseMock{}

func (b *BaseMock) Validator() func(string) (bool, error) {
	if validfn != nil {
		return validfn
	}
	return nil
}

func (b *BaseMock) ValidatorExternal() func(string) (bool, error) {
	if validfn != nil {
		return validfn
	}
	return nil
}

func (b *BaseMock) ValidatorDefault() func(string) (bool, error) {

	return nil
}

func (b *BaseMock) IsDefaultValidator() bool {
	return false
}

func (b *BaseMock) IsExternalValidator() bool {
	return false
}

func (b *BaseMock) SetValidator(fn func(string) (bool, error)) {
	validfn = fn
}

func (b *BaseMock) SetValidatorNil() {
	validfn = nil
}
