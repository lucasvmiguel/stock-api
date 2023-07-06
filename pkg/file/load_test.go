package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Foo string `json:"foo"`
}

func TestLoadSuccess(t *testing.T) {
	result, err := LoadJSON[TestStruct]("./test.json")
	if err != nil {
		t.Error("error should be nil")
	}

	assert.Equal(t, result.Foo, "bar")
}

func TestValidateError(t *testing.T) {
	_, err := LoadJSON[TestStruct]("./invalid.json")
	if err == nil {
		t.Error("error should not be nil")
	}
}
