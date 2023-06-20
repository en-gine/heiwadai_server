package validator

import (
	"server/validator"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	validator := validator.PasswordValidator{}

	testCases := []struct {
		password string
		want     bool
	}{
		{"Pa$$w0rd", true},
		{"password", false},    // 大文字と数字と記号が含まれていない
		{"PASSWORD123", false}, // 小文字と記号が含まれていない
		{"Pa$$", false},        // 8文字未満
		{"passw0rd", false},    // 大文字と記号が含まれていない
	}

	for _, tc := range testCases {
		isValid, _ := validator.IsValid(tc.password)
		if isValid != tc.want {
			t.Errorf("IsValidPassword(%q) = %v; want %v", tc.password, isValid, tc.want)
		}
	}
}
