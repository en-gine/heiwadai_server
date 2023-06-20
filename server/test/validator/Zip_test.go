package validator

import (
	"server/validator"
	"testing"
)

func TestZipValidator_IsValid(t *testing.T) {
	validator := &validator.ZipValidator{}

	// Valid inputs
	if IsValid, err := validator.IsValid("123-4567"); !IsValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got IsValid=%v, err=%v", IsValid, err)
	}

	// Invalid inputs
	if IsValid, err := validator.IsValid("12345"); IsValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got IsValid=%v, err=%v", IsValid, err)
	}
	if IsValid, err := validator.IsValid("12345a"); IsValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got IsValid=%v, err=%v", IsValid, err)
	}
}
