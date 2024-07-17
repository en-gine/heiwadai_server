package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMail(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		isValid  bool
		expected string
	}{
		{
			name:     "Valid email",
			input:    "test@example.com",
			isValid:  true,
			expected: "test@example.com",
		},
		{
			name:    "Invalid email - no @",
			input:   "testexample.com",
			isValid: false,
		},
		{
			name:    "Invalid email - no domain",
			input:   "test@.com",
			isValid: false,
		},
		{
			name:    "Invalid email - no username",
			input:   "@example.com",
			isValid: false,
		},
		{
			name:     "Valid email with subdomain",
			input:    "test@sub.example.com",
			isValid:  true,
			expected: "test@sub.example.com",
		},
		{
			name:    "Invalid email - special characters",
			input:   "test!@example.com",
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mail, err := NewMail(tc.input)

			if tc.isValid {
				assert.Nil(t, err)
				assert.NotNil(t, mail)
				assert.Equal(t, tc.expected, mail.String())
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, mail)
				assert.Contains(t, err.Error(), "メールアドレスの形式が正しくありません")
			}
		})
	}
}

func TestMailString(t *testing.T) {
	email := "test@example.com"
	mail, _ := NewMail(email)

	assert.Equal(t, email, mail.String())
}
