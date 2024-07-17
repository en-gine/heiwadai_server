package entity

import (
	"testing"

	"server/infrastructure/encrypt"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Valid password", "ValidPass1", false},
		{"Invalid password", "short", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPassword(tt.password)
			if tt.wantErr {
				assert.Nil(t, got)
				assert.NotNil(t, err)
			} else {
				assert.NotNil(t, got)
				assert.Nil(t, err)
				assert.Equal(t, tt.password, got.string())
			}
		})
	}
}

func TestPassword_DecriptedString(t *testing.T) {
	// 暗号化されたパスワードを作成
	encryptedPass, _ := encrypt.Encrypt("TestPass1")
	password := Password(encryptedPass)

	decrypted, err := password.DecriptedString()
	assert.Nil(t, err)
	assert.Equal(t, "TestPass1", decrypted)
}

func TestPassword_EncriptedString(t *testing.T) {
	password := Password("TestPass1")

	encrypted, err := password.EncriptedString()
	assert.Nil(t, err)
	assert.NotNil(t, encrypted)
	assert.NotEqual(t, password.string(), encrypted.string())

	// 暗号化されたパスワードを復号化して元のパスワードと比較
	decrypted, err := encrypted.DecriptedString()
	assert.Nil(t, err)
	assert.Equal(t, password.string(), decrypted)
}

func TestGenerateRandomPassword(t *testing.T) {
	password, err := GenerateRandomPassword()
	assert.Nil(t, err)
	assert.NotNil(t, password)

	// パスワードが暗号化されていることを確認
	decrypted, domainErr := password.DecriptedString()
	assert.Nil(t, domainErr)

	// パスワードの長さと構成を確認
	assert.GreaterOrEqual(t, len(decrypted), 8)
	assert.Regexp(t, "[A-Z]", decrypted)
	assert.Regexp(t, "[a-z]", decrypted)
	assert.Regexp(t, "[0-9]", decrypted)
}
