package parser

import (
	"fmt"
	"testing"
)

func TestStringToUint(t *testing.T) {
	value := uint(10)
	num, err := StringToUint(fmt.Sprintf("%d", value))
	if err != nil {
		t.Error("error should be nil")
	}

	if num != value {
		t.Error("error should be nil")
	}
}

func TestStringToUintNegative(t *testing.T) {
	_, err := StringToUint("-10")
	if err == nil {
		t.Error("should return error")
	}
}

func TestStringToUintInvalid(t *testing.T) {
	_, err := StringToUint("abc")
	if err == nil {
		t.Error("should return error")
	}
}
