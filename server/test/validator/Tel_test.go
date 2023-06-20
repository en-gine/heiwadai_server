package validator

import (
	"server/validator"
	"testing"
)

func TestTelValidator_IsValid(t *testing.T) {
	validator := &validator.TelValidator{}

	// Valid inputs
	if isValid, err := validator.IsValid("080-1234-5678"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("03-1234-5678"); !isValid || err != nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}

	// Invalid inputs
	if isValid, err := validator.IsValid("12345678901"); isValid || err == nil {
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("12-3456-7890"); isValid || err == nil { // false, 先頭に0が必要
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("012-345-67"); isValid || err == nil { // false, 最後の数字グループは3～4桁必要
		t.Errorf("Expected true and no error for valid input, but got isValid=%v, err=%v", isValid, err)
	}
	if isValid, err := validator.IsValid("abc-defg-hijk"); isValid || err == nil {
		t.Errorf("Expected false and an error for invalid input, but got isValid=%v, err=%v", isValid, err)
	}
}
