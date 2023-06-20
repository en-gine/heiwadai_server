package validator

import (
	"server/validator"
	"testing"
)

func TestKanaValidator_IsValid(t *testing.T) {
	validator := &validator.KanaValidator{}

	// Valid inputs
	if isValid, err := validator.IsValid("アイウエオ"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("カキクケコ"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}

	// Invalid inputs
	if isValid, err := validator.IsValid("亜伊宇江於"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("アイウエオ1"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}

	if isValid, err := validator.IsValid("ｱｲｳｴ"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}
}
