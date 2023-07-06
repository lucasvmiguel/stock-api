package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	ID   int    `validate:"required"`
	Name string `validate:"required"`
}

func TestValidateSuccess(t *testing.T) {
	err := Validate(TestStruct{
		ID:   1,
		Name: "test",
	})
	if err != nil {
		t.Error("error should be nil")
	}
}

func TestValidateError(t *testing.T) {
	err := Validate(TestStruct{})

	if err == nil {
		t.Error("should return error")
	}

	if len(err) != 2 {
		t.Error("should return 2 errors")
	}

	assert.Equal(t, err[0], ValidationError{Field: "ID", Tag: "required", Value: ""})
	assert.Equal(t, err[1], ValidationError{Field: "Name", Tag: "required", Value: ""})
}
