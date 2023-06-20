package validator

import (
	"server/validator"
	"testing"
)

func TestEmailValidator_IsValid(t *testing.T) {
	validator := &validator.EmailValidator{}

	// Valid inputs
	if isValid, err := validator.IsValid("test@example.com"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("test.user-01@example.co.jp"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}

	// Invalid inputs
	if isValid, err := validator.IsValid("test@i"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("test.example.com"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}
}
